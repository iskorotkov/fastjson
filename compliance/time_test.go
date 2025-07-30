package compliance

import (
	"testing"
	"time"
)

func TestMarshal_Duration(t *testing.T) {
	runMarshalTest(t, "zero", time.Duration(0), `"0s"`)
	runMarshalTest(t, "nanosecond", time.Nanosecond, `"1ns"`)
	runMarshalTest(t, "microsecond", time.Microsecond, `"1µs"`)
	runMarshalTest(t, "millisecond", time.Millisecond, `"1ms"`)
	runMarshalTest(t, "second", time.Second, `"1s"`)
	runMarshalTest(t, "minute", time.Minute, `"1m0s"`)
	runMarshalTest(t, "hour", time.Hour, `"1h0m0s"`)
	runMarshalTest(t, "complex", 2*time.Hour+30*time.Minute+45*time.Second, `"2h30m45s"`)
	runMarshalTest(t, "negative", -time.Second, `"-1s"`)
	runMarshalTest(t, "pointer/nil", (*time.Duration)(nil), "null")
	runMarshalTest(t, "pointer/value", addr(time.Second), `"1s"`)
}

func TestMarshal_Time(t *testing.T) {
	runMarshalTest(t, "zero", time.Time{}, `"0001-01-01T00:00:00Z"`)
	runMarshalTest(t, "unix_epoch", time.Unix(0, 0).UTC(), `"1970-01-01T00:00:00Z"`)
	runMarshalTest(t, "specific", mustParseTime("2023-06-15T14:30:00Z"), `"2023-06-15T14:30:00Z"`)
	runMarshalTest(t, "with_nanos", mustParseTime("2023-06-15T14:30:00.123456789Z"), `"2023-06-15T14:30:00.123456789Z"`)
	runMarshalTest(t, "pointer/nil", (*time.Time)(nil), "null")
	runMarshalTest(t, "pointer/value", addr(mustParseTime("2023-06-15T14:30:00Z")), `"2023-06-15T14:30:00Z"`)
}

func TestUnmarshal_Duration(t *testing.T) {
	runUnmarshalTest(t, "zero", `"0s"`, time.Duration(0))
	runUnmarshalTest(t, "nanosecond", `"1ns"`, time.Nanosecond)
	runUnmarshalTest(t, "microsecond", `"1µs"`, time.Microsecond)
	runUnmarshalTest(t, "microsecond_alt", `"1us"`, time.Microsecond)
	runUnmarshalTest(t, "millisecond", `"1ms"`, time.Millisecond)
	runUnmarshalTest(t, "second", `"1s"`, time.Second)
	runUnmarshalTest(t, "minute", `"1m"`, time.Minute)
	runUnmarshalTest(t, "minute_full", `"1m0s"`, time.Minute)
	runUnmarshalTest(t, "hour", `"1h"`, time.Hour)
	runUnmarshalTest(t, "hour_full", `"1h0m0s"`, time.Hour)
	runUnmarshalTest(t, "complex", `"2h30m45s"`, 2*time.Hour+30*time.Minute+45*time.Second)
	runUnmarshalTest(t, "negative", `"-1s"`, -time.Second)

	runUnmarshalTest(t, "integer/nanoseconds", "1000000000", time.Second)
	runUnmarshalTest(t, "integer/zero", "0", time.Duration(0))

	runUnmarshalErrorTest[time.Duration](t, "invalid/bool", "true")
	runUnmarshalErrorTest[time.Duration](t, "invalid/malformed", `"invalid"`)
}

func TestUnmarshal_Time(t *testing.T) {
	runUnmarshalTest(t, "rfc3339", `"2023-06-15T14:30:00Z"`, mustParseTime("2023-06-15T14:30:00Z"))
	runUnmarshalTest(t, "rfc3339_nanos", `"2023-06-15T14:30:00.123456789Z"`, mustParseTime("2023-06-15T14:30:00.123456789Z"))
	runUnmarshalTest(t, "unix_epoch", `"1970-01-01T00:00:00Z"`, time.Unix(0, 0).UTC())

	runUnmarshalErrorTest[time.Time](t, "invalid/number", "1234567890")
	runUnmarshalErrorTest[time.Time](t, "invalid/bool", "true")
	runUnmarshalErrorTest[time.Time](t, "invalid/malformed", `"not-a-date"`)
}
