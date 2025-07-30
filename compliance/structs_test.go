package compliance

import "testing"

func TestMarshal_Structs(t *testing.T) {
	runMarshalTest(t, "empty", struct{}{}, "{}")

	runMarshalTest(t, "simple", SimpleStruct{Name: "test", Value: 42}, `{"name":"test","value":42}`)
	runMarshalTest(t, "simple/zero", SimpleStruct{}, `{"name":"","value":0}`)

	runMarshalTest(t, "nested", NestedStruct{
		Inner: SimpleStruct{Name: "inner", Value: 1},
		Extra: "extra",
	}, `{"inner":{"name":"inner","value":1},"extra":"extra"}`)

	runMarshalTest(t, "with_pointers/nil", StructWithPointers{}, `{"name":null,"value":null}`)
	runMarshalTest(t, "with_pointers/set", StructWithPointers{
		Name:  addr("test"),
		Value: addr(42),
	}, `{"name":"test","value":42}`)

	runMarshalTest(t, "with_slice/nil", StructWithSlice{}, `{"items":null}`)
	runMarshalTest(t, "with_slice/empty", StructWithSlice{Items: []string{}}, `{"items":[]}`)
	runMarshalTest(t, "with_slice/values", StructWithSlice{Items: []string{"a", "b"}}, `{"items":["a","b"]}`)

	runMarshalTest(t, "with_map/nil", StructWithMap{}, `{"data":null}`)
	runMarshalTest(t, "with_map/empty", StructWithMap{Data: map[string]int{}}, `{"data":{}}`)
	runMarshalTest(t, "with_map/values", StructWithMap{Data: map[string]int{"x": 1}}, `{"data":{"x":1}}`)

	runMarshalTest(t, "all_types", StructWithAllTypes{
		Bool:    true,
		Int:     42,
		Int64:   9223372036854775807,
		Uint:    42,
		Float64: 3.14,
		String:  "hello",
	}, `{"bool":true,"int":42,"int64":9223372036854775807,"uint":42,"float64":3.14,"string":"hello"}`)

	runMarshalTest(t, "pointer/nil", (*SimpleStruct)(nil), "null")
	runMarshalTest(t, "pointer/value", addr(SimpleStruct{Name: "test", Value: 42}), `{"name":"test","value":42}`)
}

func TestMarshal_Pointers(t *testing.T) {
	runMarshalTest(t, "nil/bool", (*bool)(nil), "null")
	runMarshalTest(t, "nil/string", (*string)(nil), "null")
	runMarshalTest(t, "nil/int", (*int)(nil), "null")
	runMarshalTest(t, "nil/slice", (*[]int)(nil), "null")
	runMarshalTest(t, "nil/map", (*map[string]int)(nil), "null")
	runMarshalTest(t, "nil/struct", (*SimpleStruct)(nil), "null")

	runMarshalTest(t, "value/bool", addr(true), "true")
	runMarshalTest(t, "value/string", addr("hello"), `"hello"`)
	runMarshalTest(t, "value/int", addr(42), "42")
	runMarshalTest(t, "value/slice", addr([]int{1, 2}), "[1,2]")
	runMarshalTest(t, "value/struct", addr(SimpleStruct{Name: "test", Value: 1}), `{"name":"test","value":1}`)

	var pp **int = nil
	runMarshalTest(t, "double/nil", pp, "null")
	runMarshalTest(t, "double/nil_inner", addr((*int)(nil)), "null")
	runMarshalTest(t, "double/value", addr(addr(42)), "42")
}

func TestUnmarshal_Structs(t *testing.T) {
	runUnmarshalTest(t, "empty", "{}", struct{}{})

	runUnmarshalTest(t, "simple", `{"name":"test","value":42}`, SimpleStruct{Name: "test", Value: 42})
	runUnmarshalTest(t, "simple/partial", `{"name":"test"}`, SimpleStruct{Name: "test", Value: 0})
	runUnmarshalTest(t, "simple/extra_whitespace", `{ "name" : "test" , "value" : 42 }`, SimpleStruct{Name: "test", Value: 42})

	runUnmarshalTest(t, "nested", `{"inner":{"name":"inner","value":1},"extra":"extra"}`, NestedStruct{
		Inner: SimpleStruct{Name: "inner", Value: 1},
		Extra: "extra",
	})

	runUnmarshalTest(t, "with_pointers/null", `{"name":null,"value":null}`, StructWithPointers{})
	runUnmarshalTest(t, "with_pointers/set", `{"name":"test","value":42}`, StructWithPointers{
		Name:  addr("test"),
		Value: addr(42),
	})

	runUnmarshalTest(t, "with_slice/null", `{"items":null}`, StructWithSlice{})
	runUnmarshalTest(t, "with_slice/empty", `{"items":[]}`, StructWithSlice{Items: []string{}})
	runUnmarshalTest(t, "with_slice/values", `{"items":["a","b"]}`, StructWithSlice{Items: []string{"a", "b"}})

	runUnmarshalTest(t, "with_map/null", `{"data":null}`, StructWithMap{})
	runUnmarshalTest(t, "with_map/empty", `{"data":{}}`, StructWithMap{Data: map[string]int{}})
	runUnmarshalTest(t, "with_map/values", `{"data":{"x":1}}`, StructWithMap{Data: map[string]int{"x": 1}})

	runUnmarshalTest(t, "all_types", `{"bool":true,"int":42,"int64":9223372036854775807,"uint":42,"float64":3.14,"string":"hello"}`, StructWithAllTypes{
		Bool:    true,
		Int:     42,
		Int64:   9223372036854775807,
		Uint:    42,
		Float64: 3.14,
		String:  "hello",
	})

	runUnmarshalErrorTest[SimpleStruct](t, "invalid/array", "[]")
	runUnmarshalErrorTest[SimpleStruct](t, "invalid/string", `"{}"`)
}

func TestUnmarshal_Pointers(t *testing.T) {
	runUnmarshalTest(t, "null/bool", "null", (*bool)(nil))
	runUnmarshalTest(t, "null/string", "null", (*string)(nil))
	runUnmarshalTest(t, "null/int", "null", (*int)(nil))
	runUnmarshalTest(t, "null/slice", "null", (*[]int)(nil))
	runUnmarshalTest(t, "null/map", "null", (*map[string]int)(nil))
	runUnmarshalTest(t, "null/struct", "null", (*SimpleStruct)(nil))

	runUnmarshalTest(t, "value/bool", "true", addr(true))
	runUnmarshalTest(t, "value/string", `"hello"`, addr("hello"))
	runUnmarshalTest(t, "value/int", "42", addr(42))
	runUnmarshalTest(t, "value/slice", "[1,2]", addr([]int{1, 2}))
	runUnmarshalTest(t, "value/struct", `{"name":"test","value":1}`, addr(SimpleStruct{Name: "test", Value: 1}))

	runUnmarshalTest(t, "double/null", "null", (**int)(nil))
	runUnmarshalTest(t, "double/value", "42", addr(addr(42)))
}
