package simd

import "encoding/binary"

const (
	QuoteMask     = 0x2222222222222222 // '"' repeated 8 times.
	BackslashMask = 0x5C5C5C5C5C5C5C5C // '\\' repeated 8 times.
	ControlMask   = 0xE0E0E0E0E0E0E0E0 // Mask for control char detection.
)

func HasZeroByte(v uint64) bool {
	return ((v - 0x0101010101010101) & ^v & 0x8080808080808080) != 0
}

func LoadWord(s []byte) uint64 {
	return binary.LittleEndian.Uint64(s)
}

func HasQuote(word uint64) bool {
	return HasZeroByte(word ^ QuoteMask)
}

func HasBackslash(word uint64) bool {
	return HasZeroByte(word ^ BackslashMask)
}

func HasControlChar(word uint64) bool {
	return HasZeroByte(word & ControlMask)
}

func HasEscapeChar(word uint64) bool {
	return HasQuote(word) || HasBackslash(word) || HasControlChar(word)
}

func HasStringTerminator(word uint64) bool {
	return HasQuote(word) || HasBackslash(word)
}
