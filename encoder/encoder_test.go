package encoder_test

import (
	"net"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"slices"
	"testing"
	"time"

	"github.com/iskorotkov/fastjson/buffer"
	"github.com/iskorotkov/fastjson/encoder"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

func TestEncoder(t *testing.T) {
	type objectType struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type complexObjectType struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Friends []struct {
			Name string `json:"name"`
		} `json:"friends"`
	}

	type objectTypeWithPointer struct {
		Name     *string  `json:"name"`
		LastName *string  `json:"last_name"`
		Email    *string  `json:"email"`
		Age      *int     `json:"age"`
		Friends  []string `json:"friends"`
	}

	cases := []struct {
		name        string
		value       reflect.Value
		expected    string
		expectedAlt []string
	}{
		{
			name:     "null",
			value:    reflect.ValueOf(nil),
			expected: "null",
		},
		{
			name:     "bool",
			value:    reflect.ValueOf(true),
			expected: "true",
		},
		{
			name:     "int",
			value:    reflect.ValueOf(42),
			expected: "42",
		},
		{
			name:     "float",
			value:    reflect.ValueOf(3.14),
			expected: "3.14",
		},
		{
			name:     "string",
			value:    reflect.ValueOf("hello"),
			expected: `"hello"`,
		},
		{
			name:     "slice",
			value:    reflect.ValueOf([]int{1, 2, 3}),
			expected: "[1,2,3]",
		},
		{
			name:     "null slice",
			value:    reflect.ValueOf(([]int)(nil)),
			expected: "null",
		},
		{
			name:     "array",
			value:    reflect.ValueOf([3]int{1, 2, 3}),
			expected: "[1,2,3]",
		},
		{
			name:        "map",
			value:       reflect.ValueOf(map[string]int{"key1": 1, "key2": 2}),
			expected:    `{"key1":1,"key2":2}`,
			expectedAlt: []string{`{"key2":2,"key1":1}`},
		},
		{
			name:     "null map",
			value:    reflect.ValueOf((map[string]int)(nil)),
			expected: "null",
		},
		{
			name:     "struct",
			value:    reflect.ValueOf(objectType{Name: "John", Age: 30}),
			expected: `{"name":"John","age":30}`,
		},
		{
			name: "nested structs",
			value: reflect.ValueOf(complexObjectType{
				Name: "John",
				Age:  30,
				Friends: []struct {
					Name string `json:"name"`
				}{
					{Name: "Doe"},
				},
			}),
			expected: `{"name":"John","age":30,"friends":[{"name":"Doe"}]}`,
		},
		{
			name: "pointers",
			value: reflect.ValueOf(&objectTypeWithPointer{
				Name: func() *string {
					name := "John"
					return &name
				}(),
				Age: func() *int {
					age := 30
					return &age
				}(),
				Friends: []string{},
			}),
			expected: `{"name":"John","last_name":null,"email":null,"age":30,"friends":[]}`,
		},
		{
			name:     "duration",
			value:    reflect.ValueOf(time.Hour + 2*time.Minute + 3*time.Second),
			expected: `"1h2m3s"`,
		},
		{
			name:     "time",
			value:    reflect.ValueOf(time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)),
			expected: `"2023-10-01T12:00:00Z"`,
		},
		{
			name:     "ip",
			value:    reflect.ValueOf(net.IPv4(192, 168, 0, 1)),
			expected: `"192.168.0.1"`,
		},
		{
			name:     "url",
			value:    reflect.ValueOf(&url.URL{Scheme: "https", Host: "example.com"}),
			expected: `"https://example.com"`,
		},
		{
			name: "struct with unexported field",
			value: reflect.ValueOf(struct {
				Name    string `json:"name"`
				Age     int    `json:"age"`
				private string
			}{Name: "John", Age: 30, private: "secret"}),
			expected: `{"name":"John","age":30}`,
		},
		{
			name: "struct with json:- field",
			value: reflect.ValueOf(struct {
				Name   string `json:"name"`
				Secret string `json:"-"`
			}{Name: "John", Secret: "hidden"}),
			expected: `{"name":"John"}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !c.value.IsValid() {
				enc := encoder.New(nil)

				buf := buffer.New()
				enc(c.value, &buf)
				if c.value.IsValid() {
					t.Fatalf("expected nil destination, got %v", c.value.Interface())
				}
				return
			}

			enc := encoder.New(c.value.Type())

			buf := buffer.New()
			enc(c.value, &buf)

			got := string(buf.Clone())
			if got != c.expected && !slices.Contains(c.expectedAlt, got) {
				t.Fatalf("expected %v, got %v", c.expected, got)
			}
		})
	}
}

func TestEncoderOmitEmpty(t *testing.T) {
	type omitEmptyStruct struct {
		String      string         `json:"string,omitempty"`
		Int         int            `json:"int,omitempty"`
		Float       float64        `json:"float,omitempty"`
		Bool        bool           `json:"bool,omitempty"`
		Slice       []int          `json:"slice,omitempty"`
		Pointer     *string        `json:"pointer,omitempty"`
		Map         map[string]int `json:"map,omitempty"`
		AlwaysShown string         `json:"always_shown"`
	}

	cases := []struct {
		name     string
		value    omitEmptyStruct
		expected string
	}{
		{
			name:     "all zero values",
			value:    omitEmptyStruct{},
			expected: `{"always_shown":""}`,
		},
		{
			name: "all non-zero values",
			value: omitEmptyStruct{
				String:      "hello",
				Int:         42,
				Float:       3.14,
				Bool:        true,
				Slice:       []int{1, 2, 3},
				Pointer:     func() *string { s := "ptr"; return &s }(),
				Map:         map[string]int{"a": 1},
				AlwaysShown: "shown",
			},
			expected: `{"string":"hello","int":42,"float":3.14,"bool":true,"slice":[1,2,3],"pointer":"ptr","map":{"a":1},"always_shown":"shown"}`,
		},
		{
			name: "mixed zero and non-zero",
			value: omitEmptyStruct{
				String:      "hello",
				Int:         0,
				Float:       0,
				Bool:        false,
				Slice:       nil,
				Pointer:     nil,
				Map:         nil,
				AlwaysShown: "visible",
			},
			expected: `{"string":"hello","always_shown":"visible"}`,
		},
		{
			name: "empty slice vs nil slice",
			value: omitEmptyStruct{
				Slice:       []int{},
				AlwaysShown: "test",
			},
			expected: `{"always_shown":"test"}`,
		},
		{
			name: "empty map vs nil map",
			value: omitEmptyStruct{
				Map:         map[string]int{},
				AlwaysShown: "test",
			},
			expected: `{"always_shown":"test"}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			enc := encoder.New(reflect.TypeFor[omitEmptyStruct]())
			buf := buffer.New()
			enc(reflect.ValueOf(c.value), &buf)

			got := string(buf.Clone())
			if got != c.expected {
				t.Errorf("expected %s, got %s", c.expected, got)
			}
		})
	}
}

func TestEncoderSkipField(t *testing.T) {
	type skipFieldStruct struct {
		Visible string `json:"visible"`
		Skipped string `json:"-"`
		Another string `json:"another"`
	}

	cases := []struct {
		name     string
		value    skipFieldStruct
		expected string
	}{
		{
			name: "skip field with value",
			value: skipFieldStruct{
				Visible: "shown",
				Skipped: "should not appear",
				Another: "also shown",
			},
			expected: `{"visible":"shown","another":"also shown"}`,
		},
		{
			name:     "skip field empty",
			value:    skipFieldStruct{},
			expected: `{"visible":"","another":""}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			enc := encoder.New(reflect.TypeFor[skipFieldStruct]())
			buf := buffer.New()
			enc(reflect.ValueOf(c.value), &buf)

			got := string(buf.Clone())
			if got != c.expected {
				t.Errorf("expected %s, got %s", c.expected, got)
			}
		})
	}
}

func TestEncoderNestedOmitEmpty(t *testing.T) {
	type inner struct {
		Value string `json:"value,omitempty"`
	}

	type outer struct {
		Inner inner  `json:"inner"`
		Name  string `json:"name,omitempty"`
	}

	cases := []struct {
		name     string
		value    outer
		expected string
	}{
		{
			name:     "nested struct with empty inner value",
			value:    outer{Inner: inner{Value: ""}, Name: "test"},
			expected: `{"inner":{},"name":"test"}`,
		},
		{
			name:     "nested struct with non-empty inner value",
			value:    outer{Inner: inner{Value: "hello"}, Name: ""},
			expected: `{"inner":{"value":"hello"}}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			enc := encoder.New(reflect.TypeFor[outer]())
			buf := buffer.New()
			enc(reflect.ValueOf(c.value), &buf)

			got := string(buf.Clone())
			if got != c.expected {
				t.Errorf("expected %s, got %s", c.expected, got)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	typ := reflect.TypeOf(struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Friends []struct {
			Name string `json:"name"`
		} `json:"friends"`
	}{})

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = encoder.New(typ)
	}
}

func BenchmarkEncode(b *testing.B) {
	val := struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Friends []struct {
			Name string `json:"name"`
		} `json:"friends"`
	}{
		Name: "John",
		Age:  30,
		Friends: []struct {
			Name string `json:"name"`
		}{
			{Name: "Doe"},
		},
	}

	enc := encoder.New(reflect.TypeOf(val))

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		buf := buffer.New()
		enc(reflect.ValueOf(&val).Elem(), &buf)
		b.SetBytes(int64(len(buf.Clone())))
	}
}
