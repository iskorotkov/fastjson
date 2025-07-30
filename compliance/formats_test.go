package compliance

import (
	"math"
	"testing"
	"time"
)

func TestMarshal_ByteFormats(t *testing.T) {
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}

	runMarshalTest(t, "base64/default", data, `"SGVsbG8="`)
	runMarshalTest(t, "base64url", StructFormatBytes{Base64URL: data}, `{"base64url":"SGVsbG8"}`)
	runMarshalTest(t, "base32", StructFormatBytes{Base32: data}, `{"base32":"JBSWY3DP"}`)
	runMarshalTest(t, "base16/hex", StructFormatBytes{Base16: data}, `{"base16":"48656C6C6F"}`)
	runMarshalTest(t, "array", StructFormatBytes{Array: data}, `{"array":[72,101,108,108,111]}`)
}

func TestUnmarshal_ByteFormats(t *testing.T) {
	runUnmarshalTest(t, "base64", `"SGVsbG8="`, []byte("Hello"))
	runUnmarshalTest(t, "base64url", `"SGVsbG8"`, []byte("Hello"))
	runUnmarshalTest(t, "base16", `"48656C6C6F"`, []byte("Hello"))
	runUnmarshalTest(t, "array", `[72,101,108,108,111]`, []byte("Hello"))
}

func TestMarshal_TimeFormats(t *testing.T) {
	tm := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)

	runMarshalTest(t, "rfc3339", tm, `"2023-06-15T14:30:00Z"`)
	runMarshalTest(t, "rfc822", StructFormatTime{RFC822: tm}, `{"rfc822":"15 Jun 23 14:30 UTC"}`)
	runMarshalTest(t, "unix", StructFormatTime{Unix: tm}, `{"unix":1686839400}`)
	runMarshalTest(t, "unix_milli", StructFormatTime{UnixMilli: tm}, `{"unix_milli":1686839400000}`)
	runMarshalTest(t, "unix_nano", StructFormatTime{UnixNano: tm}, `{"unix_nano":1686839400000000000}`)
}

func TestUnmarshal_TimeFormats(t *testing.T) {
	tm := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)

	runUnmarshalTest(t, "rfc3339", `"2023-06-15T14:30:00Z"`, tm)
	runUnmarshalTest(t, "unix", `1686839400`, tm)
	runUnmarshalTest(t, "unix_milli", `1686839400000`, tm)
}

func TestMarshal_DurationFormats(t *testing.T) {
	d := 2*time.Hour + 30*time.Minute + 45*time.Second

	runMarshalTest(t, "default/string", d, `"2h30m45s"`)
	runMarshalTest(t, "nano", StructFormatDuration{Nano: d}, `{"nano":9045000000000}`)
	runMarshalTest(t, "milli", StructFormatDuration{Milliseconds: d}, `{"milli":9045000}`)
	runMarshalTest(t, "iso8601", StructFormatDuration{ISO8601: d}, `{"iso8601":"PT2H30M45S"}`)
}

func TestUnmarshal_DurationFormats(t *testing.T) {
	d := 2*time.Hour + 30*time.Minute + 45*time.Second

	runUnmarshalTest(t, "string", `"2h30m45s"`, d)
	runUnmarshalTest(t, "nano", `9045000000000`, d)
	runUnmarshalTest(t, "iso8601", `"PT2H30M45S"`, d)
}

func TestMarshal_FloatFormats(t *testing.T) {
	runMarshalTest(t, "nan", StructFormatFloats{NonFinite: math.NaN()}, `{"nonfinite":"NaN"}`)
	runMarshalTest(t, "inf", StructFormatFloats{NonFinite: math.Inf(1)}, `{"nonfinite":"Infinity"}`)
	runMarshalTest(t, "neg_inf", StructFormatFloats{NonFinite: math.Inf(-1)}, `{"nonfinite":"-Infinity"}`)
}

func TestUnmarshal_FloatFormats(t *testing.T) {
	runUnmarshalTest(t, "nan", `{"nonfinite":"NaN"}`, StructFormatFloats{NonFinite: math.NaN()})
	runUnmarshalTest(t, "inf", `{"nonfinite":"Infinity"}`, StructFormatFloats{NonFinite: math.Inf(1)})
}
