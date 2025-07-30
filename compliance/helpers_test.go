package compliance

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	fastjson "github.com/iskorotkov/fastjson"
)

func addr[T any](v T) *T {
	return &v
}

func mustParseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		panic(err)
	}
	return t
}

func runMarshalTest[T any](t *testing.T, name string, input T, want string) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		got, err := fastjson.NewEncoder[T]().Marshal(input)
		if err != nil {
			t.Fatalf("Marshal error: %v", err)
		}
		if string(got) != want {
			t.Errorf("Marshal(%v):\n  got:  %s\n  want: %s", input, got, want)
		}
	})
}

func runMarshalErrorTest[T any](t *testing.T, name string, input T) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		if _, err := fastjson.NewEncoder[T]().Marshal(input); err == nil {
			t.Errorf("Marshal(%v): expected error, got nil", input)
		}
	})
}

func runUnmarshalTest[T any](t *testing.T, name string, input string, want T) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		var got T
		if err := fastjson.NewDecoder[T]().Unmarshal([]byte(input), &got); err != nil {
			t.Fatalf("Unmarshal error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unmarshal(%s):\n  got:  %#v\n  want: %#v", input, got, want)
		}
	})
}

func runUnmarshalErrorTest[T any](t *testing.T, name string, input string) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		var got T
		if err := fastjson.NewDecoder[T]().Unmarshal([]byte(input), &got); err == nil {
			t.Errorf("Unmarshal(%s): expected error, got nil (value: %#v)", input, got)
		}
	})
}

type namedBool bool
type namedString string
type namedInt int
type namedUint uint
type namedFloat64 float64
type namedUintptr uintptr
type namedBytes []byte
type recursiveMap map[string]recursiveMap
type recursiveSlice []recursiveSlice
type recursivePointer struct {
	P *recursivePointer `json:"p"`
}

type stringMarshalerStruct struct {
	Value string
}

func (s stringMarshalerStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Value)
}

func (s *stringMarshalerStruct) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.Value)
}

type textMarshalerStruct struct {
	Value string
}

func (t textMarshalerStruct) MarshalText() ([]byte, error) {
	return []byte(t.Value), nil
}

func (t *textMarshalerStruct) UnmarshalText(data []byte) error {
	t.Value = string(data)
	return nil
}

type SimpleStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type NestedStruct struct {
	Inner SimpleStruct `json:"inner"`
	Extra string       `json:"extra"`
}

type StructWithPointers struct {
	Name  *string `json:"name"`
	Value *int    `json:"value"`
}

type StructWithSlice struct {
	Items []string `json:"items"`
}

type StructWithMap struct {
	Data map[string]int `json:"data"`
}

type StructWithAllTypes struct {
	Bool    bool    `json:"bool"`
	Int     int     `json:"int"`
	Int64   int64   `json:"int64"`
	Uint    uint    `json:"uint"`
	Float64 float64 `json:"float64"`
	String  string  `json:"string"`
}

type cyclicA struct {
	B *cyclicB `json:"b"`
}

type cyclicB struct {
	A *cyclicA `json:"a"`
}

type StructOmitEmpty struct {
	Name    string         `json:"name,omitempty"`
	Value   int            `json:"value,omitempty"`
	Slice   []string       `json:"slice,omitempty"`
	Map     map[string]int `json:"map,omitempty"`
	Pointer *string        `json:"pointer,omitempty"`
}

type StructOmitZero struct {
	Name    string         `json:"name,omitzero"`
	Value   int            `json:"value,omitzero"`
	Slice   []string       `json:"slice,omitzero"`
	Map     map[string]int `json:"map,omitzero"`
	Pointer *string        `json:"pointer,omitzero"`
}

type StructStringified struct {
	IntAsString   int     `json:"int_as_string,string"`
	FloatAsString float64 `json:"float_as_string,string"`
	BoolAsString  bool    `json:"bool_as_string,string"`
}

type StructInline struct {
	Name string `json:"name"`
	StructInlineInner
}

type StructInlineInner struct {
	Value int    `json:"value"`
	Extra string `json:"extra"`
}

type StructInlinePointer struct {
	Name string `json:"name"`
	*StructInlineInner
}

type StructInlineMap struct {
	Name  string         `json:"name"`
	Extra map[string]any `json:",inline"`
}

type StructUnknown struct {
	Name    string         `json:"name"`
	Unknown map[string]any `json:",unknown"`
}

type StructCaseInsensitive struct {
	Name  string `json:"name,nocase"`
	Value int    `json:"VALUE,nocase"`
}

type StructCaseStrict struct {
	Name  string `json:"name,case:strict"`  //nolint:staticcheck
	Value int    `json:"Value,case:strict"` //nolint:staticcheck
}

type StructFormatBytes struct {
	Base64    []byte `json:"base64"`
	Base64URL []byte `json:"base64url,format:base64url"`
	Base32    []byte `json:"base32,format:base32"`
	Base32Hex []byte `json:"base32hex,format:base32hex"`
	Base16    []byte `json:"base16,format:base16"`
	Array     []byte `json:"array,format:array"`
}

type StructFormatTime struct {
	RFC3339   time.Time `json:"rfc3339"`
	RFC822    time.Time `json:"rfc822,format:RFC822"`
	RFC850    time.Time `json:"rfc850,format:RFC850"`
	RFC1123   time.Time `json:"rfc1123,format:RFC1123"`
	ANSIC     time.Time `json:"ansic,format:ANSIC"`
	Unix      time.Time `json:"unix,format:Unix"`
	UnixMilli time.Time `json:"unix_milli,format:UnixMilli"`
	UnixMicro time.Time `json:"unix_micro,format:UnixMicro"`
	UnixNano  time.Time `json:"unix_nano,format:UnixNano"`
	DateOnly  time.Time `json:"date_only,format:DateOnly"`
	TimeOnly  time.Time `json:"time_only,format:TimeOnly"`
	DateTime  time.Time `json:"date_time,format:DateTime"`
}

type StructFormatDuration struct {
	Default      time.Duration `json:"default"`
	Nano         time.Duration `json:"nano,format:nano"`
	Micro        time.Duration `json:"micro,format:micro"`
	Milliseconds time.Duration `json:"milli,format:milli"` //nolint:staticcheck
	Seconds      time.Duration `json:"sec,format:sec"`     //nolint:staticcheck
	Units        time.Duration `json:"units,format:units"`
	ISO8601      time.Duration `json:"iso8601,format:iso8601"`
}

type StructFormatFloats struct {
	Normal    float64 `json:"normal"`
	NonFinite float64 `json:"nonfinite,format:nonfinite"`
}

type namedInt64 int64
type namedUint64 uint64

type deepPointerStruct struct {
	Value ***int `json:"value"`
}

type EmbeddedBase struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EmbeddedDerived struct {
	EmbeddedBase
	Extra string `json:"extra"`
}

type EmbeddedConflict struct {
	EmbeddedBase
	Name string `json:"name"`
}

type EmbeddedMultiple struct {
	EmbeddedBase
	StructInlineInner
}

type methodPrecedenceJSONv2 struct {
	Value string
}

func (m methodPrecedenceJSONv2) MarshalJSON() ([]byte, error) {
	return json.Marshal("jsonv2:" + m.Value)
}

func (m methodPrecedenceJSONv2) MarshalText() ([]byte, error) {
	return []byte("text:" + m.Value), nil
}

type methodPrecedenceJSONv1 struct {
	Value string
}

func (m methodPrecedenceJSONv1) MarshalJSON() ([]byte, error) {
	return json.Marshal("jsonv1:" + m.Value)
}

func (m methodPrecedenceJSONv1) MarshalText() ([]byte, error) {
	return []byte("text:" + m.Value), nil
}

type methodPrecedenceText struct {
	Value string
}

func (m methodPrecedenceText) MarshalText() ([]byte, error) {
	return []byte("text:" + m.Value), nil
}
