package decoder_test

import (
	"net"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"testing"
	"time"

	"github.com/iskorotkov/fastjson/decoder"
	"github.com/iskorotkov/fastjson/tokenizer"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

func TestDecoder(t *testing.T) {
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
		tokens      tokenizer.Tokenizer
		destination reflect.Value
		expected    any
	}{
		{
			name:        "null",
			tokens:      tokenizer.NewFromString("null"),
			destination: reflect.ValueOf(nil),
			expected:    nil,
		},
		{
			name:        "bool",
			tokens:      tokenizer.NewFromString("true"),
			destination: reflect.ValueOf(new(bool)).Elem(),
			expected:    true,
		},
		{
			name:        "int",
			tokens:      tokenizer.NewFromString("42"),
			destination: reflect.ValueOf(new(int)).Elem(),
			expected:    42,
		},
		{
			name:        "float",
			tokens:      tokenizer.NewFromString("3.14"),
			destination: reflect.ValueOf(new(float64)).Elem(),
			expected:    3.14,
		},
		{
			name:        "string",
			tokens:      tokenizer.NewFromString(`"hello"`),
			destination: reflect.ValueOf(new(string)).Elem(),
			expected:    "hello",
		},
		{
			name:        "slice",
			tokens:      tokenizer.NewFromString(`[1, 2, 3]`),
			destination: reflect.ValueOf(new([]int)).Elem(),
			expected:    []int{1, 2, 3},
		},
		{
			name:        "null slice",
			tokens:      tokenizer.NewFromString(`null`),
			destination: reflect.ValueOf(new([]int)).Elem(),
			expected:    []int(nil),
		},
		{
			name:        "array",
			tokens:      tokenizer.NewFromString(`[1, 2, 3]`),
			destination: reflect.ValueOf(new([3]int)).Elem(),
			expected:    [3]int{1, 2, 3},
		},
		{
			name:        "map",
			tokens:      tokenizer.NewFromString(`{"key1": 1, "key2": 2}`),
			destination: reflect.ValueOf(new(map[string]int)).Elem(),
			expected:    map[string]int{"key1": 1, "key2": 2},
		},
		{
			name:        "null map",
			tokens:      tokenizer.NewFromString(`null`),
			destination: reflect.ValueOf(new(map[string]int)).Elem(),
			expected:    map[string]int(nil),
		},
		{
			name:        "struct",
			tokens:      tokenizer.NewFromString(`{"name":"John","age":30}`),
			destination: reflect.ValueOf(new(objectType)).Elem(),
			expected:    objectType{Name: "John", Age: 30},
		},
		{
			name:        "nested structs",
			tokens:      tokenizer.NewFromString(`{"name":"John","age":30,"friends":[{"name":"Doe"}]}`),
			destination: reflect.ValueOf(new(complexObjectType)).Elem(),
			expected: complexObjectType{
				Name: "John",
				Age:  30,
				Friends: []struct {
					Name string `json:"name"`
				}{
					{Name: "Doe"},
				},
			},
		},
		{
			name:        "pointers",
			tokens:      tokenizer.NewFromString(`{"name":"John","email":null,"age":30,"friends":null}`),
			destination: reflect.ValueOf(new(*objectTypeWithPointer)).Elem(),
			expected: &objectTypeWithPointer{
				Name: func() *string {
					name := "John"
					return &name
				}(),
				Age: func() *int {
					age := 30
					return &age
				}(),
				Friends: nil,
			},
		},
		{
			name:        "duration as int",
			tokens:      tokenizer.NewFromString(`42`),
			destination: reflect.ValueOf(new(time.Duration)).Elem(),
			expected:    time.Duration(42),
		},
		{
			name:        "duration as string",
			tokens:      tokenizer.NewFromString(`"1h2m3s"`),
			destination: reflect.ValueOf(new(time.Duration)).Elem(),
			expected:    time.Hour + 2*time.Minute + 3*time.Second,
		},
		{
			name:        "time",
			tokens:      tokenizer.NewFromString(`"2023-10-01T12:00:00Z"`),
			destination: reflect.ValueOf(new(time.Time)).Elem(),
			expected:    time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			name:        "ip",
			tokens:      tokenizer.NewFromString(`"192.168.0.1"`),
			destination: reflect.ValueOf(new(net.IP)).Elem(),
			expected:    net.IPv4(192, 168, 0, 1),
		},
		{
			name:        "url",
			tokens:      tokenizer.NewFromString(`"https://example.com"`),
			destination: reflect.ValueOf(new(url.URL)).Elem(),
			expected:    url.URL{Scheme: "https", Host: "example.com"},
		},
		{
			name:        "struct with unknown string field",
			tokens:      tokenizer.NewFromString(`{"name":"John","unknown":"value","age":30}`),
			destination: reflect.ValueOf(new(objectType)).Elem(),
			expected:    objectType{Name: "John", Age: 30},
		},
		{
			name:        "struct with unknown object field",
			tokens:      tokenizer.NewFromString(`{"name":"John","unknown":{"nested":"obj","deep":{"value":123}},"age":30}`),
			destination: reflect.ValueOf(new(objectType)).Elem(),
			expected:    objectType{Name: "John", Age: 30},
		},
		{
			name:        "struct with unknown array field",
			tokens:      tokenizer.NewFromString(`{"name":"John","unknown":[1,2,3],"age":30}`),
			destination: reflect.ValueOf(new(objectType)).Elem(),
			expected:    objectType{Name: "John", Age: 30},
		},
		{
			name:        "struct with unknown nested array field",
			tokens:      tokenizer.NewFromString(`{"name":"John","unknown":[[1,2],[3,4]],"age":30}`),
			destination: reflect.ValueOf(new(objectType)).Elem(),
			expected:    objectType{Name: "John", Age: 30},
		},
		{
			name:        "struct with multiple unknown fields",
			tokens:      tokenizer.NewFromString(`{"a":"x","name":"John","b":123,"age":30,"c":[1,2],"d":{"e":"f"}}`),
			destination: reflect.ValueOf(new(objectType)).Elem(),
			expected:    objectType{Name: "John", Age: 30},
		},
		{
			name:        "struct with only unknown fields",
			tokens:      tokenizer.NewFromString(`{"unknown1":"value1","unknown2":123}`),
			destination: reflect.ValueOf(new(objectType)).Elem(),
			expected:    objectType{},
		},
		{
			name:        "nested struct with unknown fields",
			tokens:      tokenizer.NewFromString(`{"name":"John","unknown":"skip","age":30,"friends":[{"name":"Doe","extra":true}]}`),
			destination: reflect.ValueOf(new(complexObjectType)).Elem(),
			expected: complexObjectType{
				Name: "John",
				Age:  30,
				Friends: []struct {
					Name string `json:"name"`
				}{
					{Name: "Doe"},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !c.destination.IsValid() {
				dec := decoder.New(nil)
				dec(c.destination, &c.tokens)
				if c.destination.IsValid() {
					t.Fatalf("expected nil destination, got %v", c.destination.Interface())
				}
				return
			}

			dec := decoder.New(c.destination.Type())
			dec(c.destination, &c.tokens)
			if !reflect.DeepEqual(c.destination.Interface(), c.expected) {
				t.Fatalf("expected %v, got %v", c.expected, c.destination.Interface())
			}
		})
	}
}

func TestDecoderSkipField(t *testing.T) {
	type skipFieldStruct struct {
		Visible string `json:"visible"`
		Skipped string `json:"-"`
		Another string `json:"another"`
	}

	cases := []struct {
		name     string
		json     string
		expected skipFieldStruct
	}{
		{
			name:     "skip field absent in JSON",
			json:     `{"visible":"shown","another":"also shown"}`,
			expected: skipFieldStruct{Visible: "shown", Skipped: "", Another: "also shown"},
		},
		{
			name:     "all fields present except skipped",
			json:     `{"visible":"v1","another":"v2"}`,
			expected: skipFieldStruct{Visible: "v1", Skipped: "", Another: "v2"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var result skipFieldStruct
			tokens := tokenizer.NewFromString(c.json)
			dec := decoder.New(reflect.TypeFor[skipFieldStruct]())
			dec(reflect.ValueOf(&result).Elem(), &tokens)

			if !reflect.DeepEqual(result, c.expected) {
				t.Errorf("expected %+v, got %+v", c.expected, result)
			}
		})
	}
}

func TestStringIndependenceFromInputBuffer(t *testing.T) {
	type dto struct {
		Name string `json:"name"`
	}

	input := []byte(`{"name":"hello"}`)
	tokens := tokenizer.NewFromBytes(input)
	dec := decoder.New(reflect.TypeFor[dto]())

	var result dto
	dec(reflect.ValueOf(&result).Elem(), &tokens)

	if result.Name != "hello" {
		t.Fatalf("expected 'hello', got %q", result.Name)
	}

	copy(input[9:14], "XXXXX")

	if result.Name != "hello" {
		t.Fatalf("string changed after modifying input buffer: got %q, want 'hello'", result.Name)
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
		_ = decoder.New(typ)
	}
}

func BenchmarkDecode(b *testing.B) {
	value := `{"name":"John","age":30,"friends":[{"name":"Doe"}]}`

	type dto struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Friends []struct {
			Name string `json:"name"`
		} `json:"friends"`
	}

	dec := decoder.New(reflect.TypeFor[dto]())

	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(value)))

	for b.Loop() {
		tokens := tokenizer.NewFromString(value)
		dec(reflect.ValueOf(&dto{}).Elem(), &tokens)
	}
}
