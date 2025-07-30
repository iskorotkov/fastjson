package encoder

import (
	"reflect"
	"strings"
)

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return "unsupported type " + e.Type.String()
}

type MarshalerError struct {
	Err   error
	Value reflect.Value
}

func (e *MarshalerError) Error() string {
	var sb strings.Builder
	sb.WriteString("can't marshal ")
	sb.WriteString(e.Value.Type().String())
	sb.WriteString(": ")
	sb.WriteString(e.Err.Error())
	return sb.String()
}

type WriteError struct {
	Err error
}

func (e *WriteError) Error() string {
	return "write error: " + e.Err.Error()
}
