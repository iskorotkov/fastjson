//go:build !go1.25

package xreflect

import "reflect"

func TypeAssert[T any](value reflect.Value) (T, bool) {
	v, ok := value.Interface().(T)
	return v, ok
}
