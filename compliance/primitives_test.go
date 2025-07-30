package compliance

import "testing"

func TestMarshal_Nil(t *testing.T) {
	runMarshalTest(t, "nil/pointer", (*int)(nil), "null")
	runMarshalTest(t, "nil/slice", ([]int)(nil), "null")
	runMarshalTest(t, "nil/map", (map[string]int)(nil), "null")
}

func TestMarshal_Bools(t *testing.T) {
	runMarshalTest(t, "true", true, "true")
	runMarshalTest(t, "false", false, "false")
	runMarshalTest(t, "named/true", namedBool(true), "true")
	runMarshalTest(t, "named/false", namedBool(false), "false")
	runMarshalTest(t, "pointer/true", addr(true), "true")
	runMarshalTest(t, "pointer/false", addr(false), "false")
	runMarshalTest(t, "pointer/nil", (*bool)(nil), "null")
}

func TestMarshal_Strings(t *testing.T) {
	runMarshalTest(t, "empty", "", `""`)
	runMarshalTest(t, "simple", "hello", `"hello"`)
	runMarshalTest(t, "with_spaces", "hello world", `"hello world"`)
	runMarshalTest(t, "unicode", "hello\u0000world", `"hello\u0000world"`)
	runMarshalTest(t, "emoji", "\U0001F600", `"😀"`)
	runMarshalTest(t, "escape/quote", `"quoted"`, `"\"quoted\""`)
	runMarshalTest(t, "escape/backslash", `back\slash`, `"back\\slash"`)
	runMarshalTest(t, "escape/newline", "line1\nline2", `"line1\nline2"`)
	runMarshalTest(t, "escape/tab", "col1\tcol2", `"col1\tcol2"`)
	runMarshalTest(t, "escape/carriage_return", "line1\rline2", `"line1\rline2"`)
	runMarshalTest(t, "named/simple", namedString("hello"), `"hello"`)
	runMarshalTest(t, "pointer/simple", addr("hello"), `"hello"`)
	runMarshalTest(t, "pointer/nil", (*string)(nil), "null")
}

func TestMarshal_Bytes(t *testing.T) {
	runMarshalTest(t, "nil", ([]byte)(nil), `""`)
	runMarshalTest(t, "empty", []byte{}, `""`)
	runMarshalTest(t, "single", []byte{1}, `"AQ=="`)
	runMarshalTest(t, "multiple", []byte{1, 2, 3}, `"AQID"`)
	runMarshalTest(t, "hello", []byte("hello"), `"aGVsbG8="`)
	runMarshalTest(t, "named/nil", namedBytes(nil), `""`)
	runMarshalTest(t, "named/empty", namedBytes{}, `""`)
	runMarshalTest(t, "named/hello", namedBytes("hello"), `"aGVsbG8="`)
	runMarshalTest(t, "pointer/nil", (*[]byte)(nil), "null")
	runMarshalTest(t, "pointer/empty", addr([]byte{}), `""`)
}

func TestUnmarshal_Nil(t *testing.T) {
	runUnmarshalErrorTest[int](t, "int", "null")
	runUnmarshalErrorTest[string](t, "string", "null")
	runUnmarshalErrorTest[bool](t, "bool", "null")
	runUnmarshalErrorTest[float64](t, "float64", "null")
}

func TestUnmarshal_Bools(t *testing.T) {
	runUnmarshalTest(t, "true", "true", true)
	runUnmarshalTest(t, "false", "false", false)
	runUnmarshalTest(t, "named/true", "true", namedBool(true))
	runUnmarshalTest(t, "named/false", "false", namedBool(false))

	runUnmarshalErrorTest[bool](t, "invalid/string", `"true"`)
	runUnmarshalErrorTest[bool](t, "invalid/number", "1")
	runUnmarshalErrorTest[bool](t, "invalid/null", "null")
}

func TestUnmarshal_Strings(t *testing.T) {
	runUnmarshalTest(t, "empty", `""`, "")
	runUnmarshalTest(t, "simple", `"hello"`, "hello")
	runUnmarshalTest(t, "with_spaces", `"hello world"`, "hello world")
	runUnmarshalTest(t, "unicode", `"hello\u0000world"`, "hello\u0000world")
	runUnmarshalTest(t, "emoji", `"😀"`, "\U0001F600")
	runUnmarshalTest(t, "escape/quote", `"\"quoted\""`, `"quoted"`)
	runUnmarshalTest(t, "escape/backslash", `"back\\slash"`, `back\slash`)
	runUnmarshalTest(t, "escape/newline", `"line1\nline2"`, "line1\nline2")
	runUnmarshalTest(t, "escape/tab", `"col1\tcol2"`, "col1\tcol2")
	runUnmarshalTest(t, "escape/carriage_return", `"line1\rline2"`, "line1\rline2")
	runUnmarshalTest(t, "escape/unicode", `"\u0048\u0065\u006c\u006c\u006f"`, "Hello")
	runUnmarshalTest(t, "named/simple", `"hello"`, namedString("hello"))

	runUnmarshalErrorTest[string](t, "invalid/number", "42")
	runUnmarshalErrorTest[string](t, "invalid/bool", "true")
	runUnmarshalErrorTest[string](t, "invalid/null", "null")
	runUnmarshalErrorTest[string](t, "invalid/unquoted", "hello")
}

func TestUnmarshal_Bytes(t *testing.T) {
	runUnmarshalTest(t, "empty", `""`, []byte{})
	runUnmarshalTest(t, "single", `"AQ=="`, []byte{1})
	runUnmarshalTest(t, "multiple", `"AQID"`, []byte{1, 2, 3})
	runUnmarshalTest(t, "hello", `"aGVsbG8="`, []byte("hello"))
	runUnmarshalTest(t, "null", "null", ([]byte)(nil))
	runUnmarshalTest(t, "named/empty", `""`, namedBytes{})
	runUnmarshalTest(t, "named/hello", `"aGVsbG8="`, namedBytes("hello"))

	runUnmarshalErrorTest[[]byte](t, "invalid/number", "123")
	runUnmarshalErrorTest[[]byte](t, "invalid/bool", "true")
	runUnmarshalErrorTest[[]byte](t, "invalid/array", "[1,2,3]")
}
