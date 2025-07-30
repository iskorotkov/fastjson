package compliance

import "testing"

func TestUnmarshal_Whitespace(t *testing.T) {
	runUnmarshalTest(t, "leading", "  42", int(42))
	runUnmarshalTest(t, "trailing", "42  ", int(42))
	runUnmarshalTest(t, "both", "  42  ", int(42))
	runUnmarshalTest(t, "newlines", "\n42\n", int(42))
	runUnmarshalTest(t, "tabs", "\t42\t", int(42))
	runUnmarshalTest(t, "mixed", " \n\t 42 \t\n ", int(42))

	runUnmarshalTest(t, "object/formatted", `{
		"name": "test",
		"value": 42
	}`, SimpleStruct{Name: "test", Value: 42})

	runUnmarshalTest(t, "array/formatted", `[
		1,
		2,
		3
	]`, []int{1, 2, 3})
}
