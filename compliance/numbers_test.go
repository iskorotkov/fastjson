package compliance

import (
	"math"
	"testing"
)

func TestMarshal_Ints(t *testing.T) {
	runMarshalTest(t, "int/zero", int(0), "0")
	runMarshalTest(t, "int/positive", int(42), "42")
	runMarshalTest(t, "int/negative", int(-42), "-42")
	runMarshalTest(t, "int/max", int(math.MaxInt), "9223372036854775807")
	runMarshalTest(t, "int/min", int(math.MinInt), "-9223372036854775808")

	runMarshalTest(t, "int8/zero", int8(0), "0")
	runMarshalTest(t, "int8/max", int8(math.MaxInt8), "127")
	runMarshalTest(t, "int8/min", int8(math.MinInt8), "-128")

	runMarshalTest(t, "int16/zero", int16(0), "0")
	runMarshalTest(t, "int16/max", int16(math.MaxInt16), "32767")
	runMarshalTest(t, "int16/min", int16(math.MinInt16), "-32768")

	runMarshalTest(t, "int32/zero", int32(0), "0")
	runMarshalTest(t, "int32/max", int32(math.MaxInt32), "2147483647")
	runMarshalTest(t, "int32/min", int32(math.MinInt32), "-2147483648")

	runMarshalTest(t, "int64/zero", int64(0), "0")
	runMarshalTest(t, "int64/max", int64(math.MaxInt64), "9223372036854775807")
	runMarshalTest(t, "int64/min", int64(math.MinInt64), "-9223372036854775808")

	runMarshalTest(t, "named/zero", namedInt(0), "0")
	runMarshalTest(t, "named/positive", namedInt(42), "42")
	runMarshalTest(t, "pointer/positive", addr(42), "42")
	runMarshalTest(t, "pointer/nil", (*int)(nil), "null")
}

func TestMarshal_Uints(t *testing.T) {
	runMarshalTest(t, "uint/zero", uint(0), "0")
	runMarshalTest(t, "uint/positive", uint(42), "42")
	runMarshalTest(t, "uint/max", uint(math.MaxUint), "18446744073709551615")

	runMarshalTest(t, "uint8/zero", uint8(0), "0")
	runMarshalTest(t, "uint8/max", uint8(math.MaxUint8), "255")

	runMarshalTest(t, "uint16/zero", uint16(0), "0")
	runMarshalTest(t, "uint16/max", uint16(math.MaxUint16), "65535")

	runMarshalTest(t, "uint32/zero", uint32(0), "0")
	runMarshalTest(t, "uint32/max", uint32(math.MaxUint32), "4294967295")

	runMarshalTest(t, "uint64/zero", uint64(0), "0")
	runMarshalTest(t, "uint64/max", uint64(math.MaxUint64), "18446744073709551615")

	runMarshalTest(t, "named/zero", namedUint(0), "0")
	runMarshalTest(t, "named/positive", namedUint(42), "42")
	runMarshalTest(t, "pointer/positive", addr(uint(42)), "42")
	runMarshalTest(t, "pointer/nil", (*uint)(nil), "null")
}

func TestMarshal_Uintptr(t *testing.T) {
	runMarshalTest(t, "zero", uintptr(0), "0")
	runMarshalTest(t, "positive", uintptr(1234), "1234")
	runMarshalTest(t, "max", uintptr(math.MaxUint), "18446744073709551615")
	runMarshalTest(t, "named/zero", namedUintptr(0), "0")
	runMarshalTest(t, "named/positive", namedUintptr(42), "42")
	runMarshalTest(t, "pointer/nil", (*uintptr)(nil), "null")
	runMarshalTest(t, "pointer/value", addr(uintptr(42)), "42")
}

func TestMarshal_Floats(t *testing.T) {
	runMarshalTest(t, "float32/zero", float32(0), "0")
	runMarshalTest(t, "float32/positive", float32(3.14), "3.14")
	runMarshalTest(t, "float32/negative", float32(-3.14), "-3.14")
	runMarshalTest(t, "float32/small", float32(0.000001), "0.000001")

	runMarshalTest(t, "float64/zero", float64(0), "0")
	runMarshalTest(t, "float64/positive", float64(3.14159265358979), "3.14159265358979")
	runMarshalTest(t, "float64/negative", float64(-3.14159265358979), "-3.14159265358979")
	runMarshalTest(t, "float64/small", float64(0.000000001), "1e-9")
	runMarshalTest(t, "float64/large", float64(1e100), "1e+100")

	runMarshalTest(t, "named/zero", namedFloat64(0), "0")
	runMarshalTest(t, "named/positive", namedFloat64(3.14), "3.14")
	runMarshalTest(t, "pointer/positive", addr(3.14), "3.14")
	runMarshalTest(t, "pointer/nil", (*float64)(nil), "null")

	runMarshalErrorTest(t, "float64/nan", math.NaN())
	runMarshalErrorTest(t, "float64/+inf", math.Inf(1))
	runMarshalErrorTest(t, "float64/-inf", math.Inf(-1))
	runMarshalErrorTest(t, "float32/nan", float32(math.NaN()))
	runMarshalErrorTest(t, "float32/+inf", float32(math.Inf(1)))
	runMarshalErrorTest(t, "float32/-inf", float32(math.Inf(-1)))
}

func TestUnmarshal_Ints(t *testing.T) {
	runUnmarshalTest(t, "int/zero", "0", int(0))
	runUnmarshalTest(t, "int/positive", "42", int(42))
	runUnmarshalTest(t, "int/negative", "-42", int(-42))
	runUnmarshalTest(t, "int/max", "9223372036854775807", int(math.MaxInt))
	runUnmarshalTest(t, "int/min", "-9223372036854775808", int(math.MinInt))

	runUnmarshalTest(t, "int8/zero", "0", int8(0))
	runUnmarshalTest(t, "int8/max", "127", int8(math.MaxInt8))
	runUnmarshalTest(t, "int8/min", "-128", int8(math.MinInt8))

	runUnmarshalTest(t, "int16/zero", "0", int16(0))
	runUnmarshalTest(t, "int16/max", "32767", int16(math.MaxInt16))
	runUnmarshalTest(t, "int16/min", "-32768", int16(math.MinInt16))

	runUnmarshalTest(t, "int32/zero", "0", int32(0))
	runUnmarshalTest(t, "int32/max", "2147483647", int32(math.MaxInt32))
	runUnmarshalTest(t, "int32/min", "-2147483648", int32(math.MinInt32))

	runUnmarshalTest(t, "int64/zero", "0", int64(0))
	runUnmarshalTest(t, "int64/max", "9223372036854775807", int64(math.MaxInt64))
	runUnmarshalTest(t, "int64/min", "-9223372036854775808", int64(math.MinInt64))

	runUnmarshalTest(t, "named/zero", "0", namedInt(0))
	runUnmarshalTest(t, "named/positive", "42", namedInt(42))

	runUnmarshalErrorTest[int](t, "invalid/string", `"42"`)
	runUnmarshalErrorTest[int](t, "invalid/bool", "true")
	runUnmarshalErrorTest[int](t, "invalid/float", "3.14")
	runUnmarshalErrorTest[int8](t, "overflow/int8", "128")
	runUnmarshalErrorTest[int8](t, "underflow/int8", "-129")
}

func TestUnmarshal_Uints(t *testing.T) {
	runUnmarshalTest(t, "uint/zero", "0", uint(0))
	runUnmarshalTest(t, "uint/positive", "42", uint(42))
	runUnmarshalTest(t, "uint/max", "18446744073709551615", uint(math.MaxUint))

	runUnmarshalTest(t, "uint8/zero", "0", uint8(0))
	runUnmarshalTest(t, "uint8/max", "255", uint8(math.MaxUint8))

	runUnmarshalTest(t, "uint16/zero", "0", uint16(0))
	runUnmarshalTest(t, "uint16/max", "65535", uint16(math.MaxUint16))

	runUnmarshalTest(t, "uint32/zero", "0", uint32(0))
	runUnmarshalTest(t, "uint32/max", "4294967295", uint32(math.MaxUint32))

	runUnmarshalTest(t, "uint64/zero", "0", uint64(0))
	runUnmarshalTest(t, "uint64/max", "18446744073709551615", uint64(math.MaxUint64))

	runUnmarshalTest(t, "named/zero", "0", namedUint(0))
	runUnmarshalTest(t, "named/positive", "42", namedUint(42))

	runUnmarshalErrorTest[uint](t, "invalid/negative", "-1")
	runUnmarshalErrorTest[uint](t, "invalid/string", `"42"`)
	runUnmarshalErrorTest[uint8](t, "overflow/uint8", "256")
}

func TestUnmarshal_Uintptr(t *testing.T) {
	runUnmarshalTest(t, "zero", "0", uintptr(0))
	runUnmarshalTest(t, "positive", "1234", uintptr(1234))
	runUnmarshalTest(t, "max", "18446744073709551615", uintptr(math.MaxUint))
	runUnmarshalTest(t, "named/zero", "0", namedUintptr(0))
	runUnmarshalTest(t, "named/positive", "42", namedUintptr(42))
	runUnmarshalErrorTest[uintptr](t, "invalid/negative", "-1")
	runUnmarshalErrorTest[uintptr](t, "invalid/string", `"42"`)
	runUnmarshalErrorTest[uintptr](t, "invalid/float", "3.14")
}

func TestUnmarshal_Floats(t *testing.T) {
	runUnmarshalTest(t, "float32/zero", "0", float32(0))
	runUnmarshalTest(t, "float32/positive", "3.14", float32(3.14))
	runUnmarshalTest(t, "float32/negative", "-3.14", float32(-3.14))
	runUnmarshalTest(t, "float32/integer", "42", float32(42))
	runUnmarshalTest(t, "float32/scientific", "1e10", float32(1e10))

	runUnmarshalTest(t, "float64/zero", "0", float64(0))
	runUnmarshalTest(t, "float64/positive", "3.14159265358979", float64(3.14159265358979))
	runUnmarshalTest(t, "float64/negative", "-3.14159265358979", float64(-3.14159265358979))
	runUnmarshalTest(t, "float64/integer", "42", float64(42))
	runUnmarshalTest(t, "float64/scientific", "1e100", float64(1e100))
	runUnmarshalTest(t, "float64/scientific_neg", "1e-100", float64(1e-100))

	runUnmarshalTest(t, "named/zero", "0", namedFloat64(0))
	runUnmarshalTest(t, "named/positive", "3.14", namedFloat64(3.14))

	runUnmarshalErrorTest[float64](t, "invalid/string", `"3.14"`)
	runUnmarshalErrorTest[float64](t, "invalid/bool", "true")
}
