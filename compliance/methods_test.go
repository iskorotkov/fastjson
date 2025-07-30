package compliance

import "testing"

func TestMarshal_Methods(t *testing.T) {
	runMarshalTest(t, "json_marshaler/simple", stringMarshalerStruct{Value: "test"}, `"test"`)
	runMarshalTest(t, "json_marshaler/empty", stringMarshalerStruct{}, `""`)
	runMarshalTest(t, "json_marshaler/pointer", addr(stringMarshalerStruct{Value: "test"}), `"test"`)
	runMarshalTest(t, "json_marshaler/nil_pointer", (*stringMarshalerStruct)(nil), "null")

	runMarshalTest(t, "text_marshaler/simple", textMarshalerStruct{Value: "test"}, `"test"`)
	runMarshalTest(t, "text_marshaler/empty", textMarshalerStruct{}, `""`)
	runMarshalTest(t, "text_marshaler/pointer", addr(textMarshalerStruct{Value: "test"}), `"test"`)
	runMarshalTest(t, "text_marshaler/nil_pointer", (*textMarshalerStruct)(nil), "null")
}

func TestMarshal_Interfaces(t *testing.T) {
	runMarshalTest(t, "nil", (any)(nil), "null")
	runMarshalTest(t, "bool/true", any(true), "true")
	runMarshalTest(t, "bool/false", any(false), "false")
	runMarshalTest(t, "string", any("hello"), `"hello"`)
	runMarshalTest(t, "int", any(42), "42")
	runMarshalTest(t, "float", any(3.14), "3.14")
	runMarshalTest(t, "slice", any([]int{1, 2}), "[1,2]")
	runMarshalTest(t, "map", any(map[string]int{"a": 1}), `{"a":1}`)
	runMarshalTest(t, "struct", any(SimpleStruct{Name: "test", Value: 42}), `{"name":"test","value":42}`)
}

func TestMarshal_Recursive(t *testing.T) {
	runMarshalTest(t, "map/empty", recursiveMap{}, "{}")
	runMarshalTest(t, "map/nested", recursiveMap{"a": recursiveMap{"b": nil}}, `{"a":{"b":null}}`)
	runMarshalTest(t, "slice/empty", recursiveSlice{}, "[]")
	runMarshalTest(t, "slice/nested", recursiveSlice{recursiveSlice{nil}}, "[[null]]")
	runMarshalTest(t, "pointer/nil", recursivePointer{}, `{"p":null}`)
	runMarshalTest(t, "pointer/nested", recursivePointer{P: &recursivePointer{}}, `{"p":{"p":null}}`)
}

func TestUnmarshal_Methods(t *testing.T) {
	runUnmarshalTest(t, "json_unmarshaler/simple", `"test"`, stringMarshalerStruct{Value: "test"})
	runUnmarshalTest(t, "json_unmarshaler/empty", `""`, stringMarshalerStruct{Value: ""})

	runUnmarshalTest(t, "text_unmarshaler/simple", `"test"`, textMarshalerStruct{Value: "test"})
	runUnmarshalTest(t, "text_unmarshaler/empty", `""`, textMarshalerStruct{Value: ""})
}

func TestUnmarshal_Interfaces(t *testing.T) {
	runUnmarshalTest(t, "null", "null", (any)(nil))
	runUnmarshalTest(t, "bool/true", "true", any(true))
	runUnmarshalTest(t, "bool/false", "false", any(false))
	runUnmarshalTest(t, "string", `"hello"`, any("hello"))
	runUnmarshalTest(t, "number/int", "42", any(float64(42)))
	runUnmarshalTest(t, "number/float", "3.14", any(3.14))
	runUnmarshalTest(t, "array", "[1,2,3]", any([]any{float64(1), float64(2), float64(3)}))
	runUnmarshalTest(t, "object", `{"a":1}`, any(map[string]any{"a": float64(1)}))
}

func TestUnmarshal_Recursive(t *testing.T) {
	runUnmarshalTest(t, "map/empty", "{}", recursiveMap{})
	runUnmarshalTest(t, "map/nested", `{"a":{"b":null}}`, recursiveMap{"a": recursiveMap{"b": nil}})
	runUnmarshalTest(t, "slice/empty", "[]", recursiveSlice{})
	runUnmarshalTest(t, "slice/nested", "[[null]]", recursiveSlice{recursiveSlice{nil}})
	runUnmarshalTest(t, "pointer/nil", `{"p":null}`, recursivePointer{})
	runUnmarshalTest(t, "pointer/nested", `{"p":{"p":null}}`, recursivePointer{P: &recursivePointer{}})
}
