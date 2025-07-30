package compliance

import (
	"testing"
)

func TestMarshal_MapIntKeys(t *testing.T) {
	runMarshalTest(t, "int/single", map[int]string{1: "one"}, `{"1":"one"}`)
	runMarshalTest(t, "int/multiple", map[int]string{1: "one", 2: "two"}, `{"1":"one","2":"two"}`)
	runMarshalTest(t, "int/negative", map[int]string{-1: "neg"}, `{"-1":"neg"}`)
	runMarshalTest(t, "int/zero", map[int]string{0: "zero"}, `{"0":"zero"}`)
	runMarshalTest(t, "int/empty", map[int]string{}, `{}`)
	runMarshalTest(t, "int/nil", (map[int]string)(nil), "null")
}

func TestUnmarshal_MapIntKeys(t *testing.T) {
	runUnmarshalTest(t, "int/single", `{"1":"one"}`, map[int]string{1: "one"})
	runUnmarshalTest(t, "int/negative", `{"-1":"neg"}`, map[int]string{-1: "neg"})
	runUnmarshalErrorTest[map[int]string](t, "int/invalid_key", `{"abc":"val"}`)
}

func TestMarshal_MapInt64Keys(t *testing.T) {
	runMarshalTest(t, "int64/large", map[int64]string{9223372036854775807: "max"}, `{"9223372036854775807":"max"}`)
	runMarshalTest(t, "int64/min", map[int64]string{-9223372036854775808: "min"}, `{"-9223372036854775808":"min"}`)
}

func TestMarshal_MapUintKeys(t *testing.T) {
	runMarshalTest(t, "uint/single", map[uint]string{1: "one"}, `{"1":"one"}`)
	runMarshalTest(t, "uint64/max", map[uint64]string{18446744073709551615: "max"}, `{"18446744073709551615":"max"}`)
}

func TestMarshal_MapBoolKeys(t *testing.T) {
	runMarshalTest(t, "bool/true", map[bool]string{true: "yes"}, `{"true":"yes"}`)
	runMarshalTest(t, "bool/false", map[bool]string{false: "no"}, `{"false":"no"}`)
	runMarshalTest(t, "bool/both", map[bool]string{true: "yes", false: "no"}, `{"false":"no","true":"yes"}`)
}

func TestUnmarshal_MapBoolKeys(t *testing.T) {
	runUnmarshalTest(t, "bool/true", `{"true":"yes"}`, map[bool]string{true: "yes"})
	runUnmarshalTest(t, "bool/false", `{"false":"no"}`, map[bool]string{false: "no"})
	runUnmarshalErrorTest[map[bool]string](t, "bool/invalid", `{"maybe":"?"}`)
}

func TestMarshal_MapFloatKeys(t *testing.T) {
	runMarshalTest(t, "float64/simple", map[float64]string{1.5: "one.five"}, `{"1.5":"one.five"}`)
	runMarshalTest(t, "float64/integer", map[float64]string{1.0: "one"}, `{"1":"one"}`)
	runMarshalTest(t, "float64/scientific", map[float64]string{1e10: "big"}, `{"10000000000":"big"}`)
}

func TestMarshal_MapNamedKeys(t *testing.T) {
	runMarshalTest(t, "namedInt64", map[namedInt64]string{1: "one"}, `{"1":"one"}`)
	runMarshalTest(t, "namedUint64", map[namedUint64]string{1: "one"}, `{"1":"one"}`)
	runMarshalTest(t, "namedString", map[namedString]int{"key": 1}, `{"key":1}`)
}

func TestMarshal_MapTextMarshalerKeys(t *testing.T) {
	runMarshalTest(t, "textMarshaler", map[textMarshalerStruct]int{{Value: "key"}: 1}, `{"key":1}`)
}
