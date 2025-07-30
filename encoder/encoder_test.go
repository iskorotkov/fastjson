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

	"github.com/iskorotkov/fastjson/encoder"
	"github.com/iskorotkov/fastjson/tiler"
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
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !c.value.IsValid() {
				enc := encoder.New(reflect.TypeOf(nil))

				tiler := tiler.New()
				enc(c.value, &tiler)
				if c.value.IsValid() {
					t.Fatalf("expected nil destination, got %v", c.value.Interface())
				}
				return
			}

			enc := encoder.New(c.value.Type())

			tiler := tiler.New()
			enc(c.value, &tiler)

			got := string(tiler.Clone())
			if got != c.expected && !slices.Contains(c.expectedAlt, got) {
				t.Fatalf("expected %v, got %v", c.expected, got)
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
		tiler := tiler.New()
		enc(reflect.ValueOf(&val).Elem(), &tiler)
		b.SetBytes(int64(len(tiler.Clone())))
	}
}
