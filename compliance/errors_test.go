package compliance

import (
	"testing"
)

func TestMarshal_InvalidUTF8(t *testing.T) {
	invalid := string([]byte{0xff, 0xfe, 0xfd})
	runMarshalErrorTest(t, "invalid_utf8", invalid)

	type structWithInvalidUTF8 struct {
		Name string `json:"name"`
	}
	runMarshalErrorTest(t, "struct_invalid_utf8", structWithInvalidUTF8{Name: invalid})
}

func TestUnmarshal_InvalidUTF8(t *testing.T) {
	runUnmarshalErrorTest[string](t, "invalid_utf8", `"\xff\xfe"`)
}

func TestUnmarshal_IntOverflow(t *testing.T) {
	runUnmarshalErrorTest[int8](t, "int8/overflow", "128")
	runUnmarshalErrorTest[int8](t, "int8/underflow", "-129")
	runUnmarshalErrorTest[int16](t, "int16/overflow", "32768")
	runUnmarshalErrorTest[int16](t, "int16/underflow", "-32769")
	runUnmarshalErrorTest[int32](t, "int32/overflow", "2147483648")
	runUnmarshalErrorTest[uint8](t, "uint8/overflow", "256")
	runUnmarshalErrorTest[uint16](t, "uint16/overflow", "65536")
	runUnmarshalErrorTest[uint32](t, "uint32/overflow", "4294967296")
}

func TestUnmarshal_TypeMismatch(t *testing.T) {
	runUnmarshalErrorTest[int](t, "string_to_int", `"hello"`)
	runUnmarshalErrorTest[string](t, "int_to_string", `42`)
	runUnmarshalErrorTest[bool](t, "string_to_bool", `"true"`)
	runUnmarshalErrorTest[[]int](t, "object_to_array", `{}`)
	runUnmarshalErrorTest[map[string]int](t, "array_to_object", `[]`)
}

func TestUnmarshal_SyntaxErrors(t *testing.T) {
	runUnmarshalErrorTest[any](t, "unclosed_brace", `{`)
	runUnmarshalErrorTest[any](t, "unclosed_bracket", `[`)
	runUnmarshalErrorTest[any](t, "unclosed_string", `"hello`)
	runUnmarshalErrorTest[any](t, "trailing_comma_object", `{"a":1,}`)
	runUnmarshalErrorTest[any](t, "trailing_comma_array", `[1,2,]`)
	runUnmarshalErrorTest[any](t, "double_comma", `[1,,2]`)
	runUnmarshalErrorTest[any](t, "missing_colon", `{"a" 1}`)
	runUnmarshalErrorTest[any](t, "missing_value", `{"a":}`)
}

func TestUnmarshal_DuplicateKeys(t *testing.T) {
	type simple struct {
		Name string `json:"name"`
	}
	runUnmarshalErrorTest[simple](t, "duplicate_keys", `{"name":"first","name":"second"}`)
}

func TestUnmarshal_EOF(t *testing.T) {
	runUnmarshalErrorTest[int](t, "empty", "")
	runUnmarshalErrorTest[int](t, "whitespace_only", "   ")
	runUnmarshalErrorTest[string](t, "incomplete", `"hello`)
}

func TestUnmarshal_ArrayLength(t *testing.T) {
	runUnmarshalErrorTest[[3]int](t, "too_short", `[1,2]`)
	runUnmarshalErrorTest[[3]int](t, "too_long", `[1,2,3,4]`)
}
