package cast

import (
	"encoding/json"
	"reflect"
	"testing"
)

type pointerMarshaler struct{}

func (p *pointerMarshaler) MarshalJSON() ([]byte, error) {
	return []byte(`"pointer"`), nil
}

type valueMarshaler struct{}

func (v valueMarshaler) MarshalJSON() ([]byte, error) {
	return []byte(`"value"`), nil
}

func TestTo_DirectSuccess(t *testing.T) {
	var m valueMarshaler
	val := reflect.ValueOf(&m).Elem()
	result, ok := To[json.Marshaler](val)
	if !ok {
		t.Fatal("expected success")
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestTo_AddressSuccess(t *testing.T) {
	var m pointerMarshaler
	val := reflect.ValueOf(&m).Elem()
	result, ok := To[json.Marshaler](val)
	if !ok {
		t.Fatal("expected success via address")
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestTo_BothFail(t *testing.T) {
	s := "test"
	val := reflect.ValueOf(&s).Elem()
	_, ok := To[json.Marshaler](val)
	if ok {
		t.Fatal("expected failure")
	}
}

func TestTo_NonAddressable(t *testing.T) {
	var m pointerMarshaler
	val := reflect.ValueOf(m)
	_, ok := To[json.Marshaler](val)
	if ok {
		t.Fatal("expected failure for non-addressable value")
	}
}
