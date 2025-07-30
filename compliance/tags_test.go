package compliance

import (
	"testing"
)

func TestMarshal_OmitEmpty(t *testing.T) {
	runMarshalTest(t, "all_empty", StructOmitEmpty{}, `{}`)
	runMarshalTest(t, "string_set", StructOmitEmpty{Name: "test"}, `{"name":"test"}`)
	runMarshalTest(t, "int_zero", StructOmitEmpty{Value: 0}, `{}`)
	runMarshalTest(t, "int_nonzero", StructOmitEmpty{Value: 1}, `{"value":1}`)
	runMarshalTest(t, "slice_nil", StructOmitEmpty{Slice: nil}, `{}`)
	runMarshalTest(t, "slice_empty", StructOmitEmpty{Slice: []string{}}, `{}`)
	runMarshalTest(t, "slice_values", StructOmitEmpty{Slice: []string{"a"}}, `{"slice":["a"]}`)
	runMarshalTest(t, "map_nil", StructOmitEmpty{Map: nil}, `{}`)
	runMarshalTest(t, "map_empty", StructOmitEmpty{Map: map[string]int{}}, `{}`)
	runMarshalTest(t, "pointer_nil", StructOmitEmpty{Pointer: nil}, `{}`)
	runMarshalTest(t, "pointer_set", StructOmitEmpty{Pointer: addr("")}, `{"pointer":""}`)
}

func TestMarshal_OmitZero(t *testing.T) {
	runMarshalTest(t, "all_zero", StructOmitZero{}, `{}`)
	runMarshalTest(t, "slice_nil", StructOmitZero{Slice: nil}, `{}`)
	runMarshalTest(t, "slice_empty", StructOmitZero{Slice: []string{}}, `{"slice":[]}`)
}

func TestMarshal_Stringified(t *testing.T) {
	runMarshalTest(t, "int", StructStringified{IntAsString: 42}, `{"int_as_string":"42"}`)
	runMarshalTest(t, "float", StructStringified{FloatAsString: 3.14}, `{"float_as_string":"3.14"}`)
	runMarshalTest(t, "bool/true", StructStringified{BoolAsString: true}, `{"bool_as_string":"true"}`)
	runMarshalTest(t, "bool/false", StructStringified{BoolAsString: false}, `{"bool_as_string":"false"}`)
}

func TestUnmarshal_Stringified(t *testing.T) {
	runUnmarshalTest(t, "int", `{"int_as_string":"42"}`, StructStringified{IntAsString: 42})
	runUnmarshalTest(t, "float", `{"float_as_string":"3.14"}`, StructStringified{FloatAsString: 3.14})
	runUnmarshalTest(t, "bool", `{"bool_as_string":"true"}`, StructStringified{BoolAsString: true})
}

func TestMarshal_Inline(t *testing.T) {
	runMarshalTest(t, "embedded", StructInline{Name: "test", StructInlineInner: StructInlineInner{Value: 1, Extra: "x"}},
		`{"name":"test","value":1,"extra":"x"}`)

	runMarshalTest(t, "inline_map", StructInlineMap{Name: "test", Extra: map[string]any{"foo": "bar"}},
		`{"name":"test","foo":"bar"}`)
}

func TestUnmarshal_Inline(t *testing.T) {
	runUnmarshalTest(t, "embedded", `{"name":"test","value":1,"extra":"x"}`,
		StructInline{Name: "test", StructInlineInner: StructInlineInner{Value: 1, Extra: "x"}})
}

func TestUnmarshal_Unknown(t *testing.T) {
	runUnmarshalTest(t, "captures_unknown", `{"name":"test","foo":"bar","num":42}`,
		StructUnknown{Name: "test", Unknown: map[string]any{"foo": "bar", "num": float64(42)}})
}

func TestUnmarshal_CaseInsensitive(t *testing.T) {
	runUnmarshalTest(t, "lowercase", `{"name":"test","value":1}`, StructCaseInsensitive{Name: "test", Value: 1})
	runUnmarshalTest(t, "uppercase", `{"NAME":"test","VALUE":1}`, StructCaseInsensitive{Name: "test", Value: 1})
	runUnmarshalTest(t, "mixed", `{"NaMe":"test","vAlUe":1}`, StructCaseInsensitive{Name: "test", Value: 1})
}

func TestMarshal_Embedded(t *testing.T) {
	runMarshalTest(t, "simple", EmbeddedDerived{EmbeddedBase: EmbeddedBase{ID: 1, Name: "test"}, Extra: "x"},
		`{"id":1,"name":"test","extra":"x"}`)

	runMarshalTest(t, "conflict", EmbeddedConflict{EmbeddedBase: EmbeddedBase{ID: 1, Name: "base"}, Name: "derived"},
		`{"id":1,"name":"derived"}`)
}
