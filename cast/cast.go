package cast

import "reflect"

func To[T any](value reflect.Value) (T, bool) {
	result, ok := reflect.TypeAssert[T](value)
	if ok {
		return result, true
	}
	if !value.CanAddr() {
		var zero T
		return zero, false
	}
	return reflect.TypeAssert[T](value.Addr())
}
