package compliance

import (
	"encoding/json"
	"math"
	"testing"
)

func TestMarshal_DeepPointers(t *testing.T) {
	val := 42
	p1 := &val
	p2 := &p1
	p3 := &p2

	runMarshalTest(t, "triple_pointer", p3, "42")
	runMarshalTest(t, "triple_nil", (***int)(nil), "null")

	var p1nil *int = nil
	p2tonil := &p1nil
	runMarshalTest(t, "pointer_to_nil", p2tonil, "null")

	deepVal := 123
	dp1 := &deepVal
	dp2 := &dp1
	dp3 := &dp2
	runMarshalTest(t, "deep_struct", deepPointerStruct{Value: dp3}, `{"value":123}`)
}

func TestMarshal_LargeNumbers(t *testing.T) {
	runMarshalTest(t, "int64/max", int64(math.MaxInt64), "9223372036854775807")
	runMarshalTest(t, "int64/min", int64(math.MinInt64), "-9223372036854775808")
	runMarshalTest(t, "uint64/max", uint64(math.MaxUint64), "18446744073709551615")
	runMarshalTest(t, "float64/max", math.MaxFloat64, "1.7976931348623157e+308")
	runMarshalTest(t, "float64/smallest", math.SmallestNonzeroFloat64, "5e-324")
}

func TestMarshal_SpecialFloats(t *testing.T) {
	runMarshalErrorTest(t, "nan", math.NaN())
	runMarshalErrorTest(t, "inf", math.Inf(1))
	runMarshalErrorTest(t, "neg_inf", math.Inf(-1))
	runMarshalTest(t, "neg_zero", math.Copysign(0, -1), "-0")
}

func TestMarshal_ZeroValues(t *testing.T) {
	runMarshalTest(t, "struct/zero", SimpleStruct{}, `{"name":"","value":0}`)
	runMarshalTest(t, "slice/nil", ([]int)(nil), "null")
	runMarshalTest(t, "slice/empty", []int{}, "[]")
	runMarshalTest(t, "map/nil", (map[string]int)(nil), "null")
	runMarshalTest(t, "map/empty", map[string]int{}, "{}")
	runMarshalTest(t, "string/empty", "", `""`)
	runMarshalTest(t, "interface/nil", (any)(nil), "null")
}

func TestMarshal_Unicode(t *testing.T) {
	runMarshalTest(t, "null_char", "\u0000", `"\u0000"`)
	runMarshalTest(t, "bom", "\ufeff", `"\ufeff"`)
	runMarshalTest(t, "replacement", "\ufffd", `"�"`)
	runMarshalTest(t, "emoji", "\U0001F44B\U0001F30D", `"👋🌍"`)
	runMarshalTest(t, "surrogate_pair", "\U0001F600", `"😀"`)
	runMarshalTest(t, "rtl", "مرحبا", `"مرحبا"`)
	runMarshalTest(t, "chinese", "你好", `"你好"`)
	runMarshalTest(t, "japanese", "こんにちは", `"こんにちは"`)
}

func TestMarshal_Escapes(t *testing.T) {
	runMarshalTest(t, "quote", `"`, `"\""`)
	runMarshalTest(t, "backslash", `\`, `"\\"`)
	runMarshalTest(t, "newline", "\n", `"\n"`)
	runMarshalTest(t, "tab", "\t", `"\t"`)
	runMarshalTest(t, "carriage_return", "\r", `"\r"`)
	runMarshalTest(t, "form_feed", "\f", `"\f"`)
	runMarshalTest(t, "backspace", "\b", `"\b"`)
	runMarshalTest(t, "all_control", "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f",
		`"\u0000\u0001\u0002\u0003\u0004\u0005\u0006\u0007\b\t\n\u000b\f\r\u000e\u000f"`)
}

func TestMarshal_RawJSON(t *testing.T) {
	type withRaw struct {
		Name string          `json:"name"`
		Raw  json.RawMessage `json:"raw"`
	}

	runMarshalTest(t, "raw/object", withRaw{Name: "test", Raw: json.RawMessage(`{"nested":true}`)},
		`{"name":"test","raw":{"nested":true}}`)
	runMarshalTest(t, "raw/array", withRaw{Name: "test", Raw: json.RawMessage(`[1,2,3]`)},
		`{"name":"test","raw":[1,2,3]}`)
	runMarshalTest(t, "raw/null", withRaw{Name: "test", Raw: nil},
		`{"name":"test","raw":null}`)
}

func TestMarshal_MethodPrecedence(t *testing.T) {
	runMarshalTest(t, "jsonv2_wins", methodPrecedenceJSONv2{Value: "test"}, `"jsonv2:test"`)
	runMarshalTest(t, "jsonv1_over_text", methodPrecedenceJSONv1{Value: "test"}, `"jsonv1:test"`)
	runMarshalTest(t, "text_only", methodPrecedenceText{Value: "test"}, `"text:test"`)
}

func TestMarshal_NilMarshaler(t *testing.T) {
	var s *stringMarshalerStruct = nil
	runMarshalTest(t, "nil_json_marshaler", s, "null")

	var tm *textMarshalerStruct = nil
	runMarshalTest(t, "nil_text_marshaler", tm, "null")
}
