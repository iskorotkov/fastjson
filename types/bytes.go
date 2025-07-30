package types

import "unsafe"

func BytesToString(b []byte) string {
	data := unsafe.SliceData(b)
	return unsafe.String(data, len(b))
}

func StringToBytes(s string) []byte {
	data := unsafe.StringData(s)
	return unsafe.Slice(data, len(s))
}
