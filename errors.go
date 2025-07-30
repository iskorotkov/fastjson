package fastjson

import (
	"reflect"
	"strings"
)

type NotAddressableError struct {
	Value reflect.Value
}

func (e *NotAddressableError) Error() string {
	var sb strings.Builder
	sb.WriteString("value ")
	sb.WriteString(e.Value.Type().String())
	sb.WriteString(" is not addressable")
	return sb.String()
}
