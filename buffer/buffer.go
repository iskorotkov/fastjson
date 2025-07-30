package buffer

import (
	"encoding/base64"
	"math"
	"slices"
	"strconv"
	"time"

	"github.com/iskorotkov/fastjson/simd"
	"github.com/iskorotkov/fastjson/view"
)

func New() Buffer {
	return Buffer{
		buf: make([]byte, 0, 8192),
	}
}

type Buffer struct {
	buf []byte
}

func (t *Buffer) PutComma() {
	t.buf = append(t.buf, ',')
}

func (t *Buffer) PutColon() {
	t.buf = append(t.buf, ':')
}

func (t *Buffer) PutObjectStart() {
	t.buf = append(t.buf, '{')
}

func (t *Buffer) PutObjectEnd() {
	t.buf = append(t.buf, '}')
}

func (t *Buffer) PutArrayStart() {
	t.buf = append(t.buf, '[')
}

func (t *Buffer) PutArrayEnd() {
	t.buf = append(t.buf, ']')
}

func (t *Buffer) PutString(s string) {
	t.PutBytes(view.StrAsBytes(s))
}

func (t *Buffer) PutQuotedString(s string) {
	t.PutQuotedBytes(view.StrAsBytes(s))
}

func (t *Buffer) PutBytes(b []byte) {
	t.buf = append(t.buf, b...)
}

func (t *Buffer) PutQuotedBytes(b []byte) {
	t.buf = appendQuote(t.buf, b)
}

func (t *Buffer) PutInt(i int64) {
	t.buf = strconv.AppendInt(t.buf, i, 10)
}

func (t *Buffer) PutUint(u uint64) {
	t.buf = strconv.AppendUint(t.buf, u, 10)
}

func (t *Buffer) PutFloat(f float64) {
	t.buf = appendFloat(t.buf, f, 64)
}

func (t *Buffer) PutFloat32(f float32) {
	t.buf = appendFloat(t.buf, float64(f), 32)
}

func (t *Buffer) PutBool(b bool) {
	if b {
		t.buf = append(t.buf, 't', 'r', 'u', 'e')
	} else {
		t.buf = append(t.buf, 'f', 'a', 'l', 's', 'e')
	}
}

func (t *Buffer) PutDuration(d time.Duration) {
	t.PutQuotedString(d.String())
}

func (t *Buffer) PutNull() {
	t.buf = append(t.buf, 'n', 'u', 'l', 'l')
}

func (t *Buffer) PutBase64(data []byte) {
	encodedLen := base64.StdEncoding.EncodedLen(len(data))
	start := len(t.buf)
	t.buf = slices.Grow(t.buf, encodedLen)
	t.buf = t.buf[:start+encodedLen]
	base64.StdEncoding.Encode(t.buf[start:], data)
}

func (t *Buffer) Clone() []byte {
	res := make([]byte, len(t.buf))
	copy(res, t.buf)
	return res
}

func (t *Buffer) Reset() {
	t.buf = t.buf[:0]
}

var controlCharEscapes = [32][]byte{
	0:    []byte(`\u0000`),
	1:    []byte(`\u0001`),
	2:    []byte(`\u0002`),
	3:    []byte(`\u0003`),
	4:    []byte(`\u0004`),
	5:    []byte(`\u0005`),
	6:    []byte(`\u0006`),
	7:    []byte(`\u0007`),
	'\b': []byte(`\b`),
	'\t': []byte(`\t`),
	'\n': []byte(`\n`),
	11:   []byte(`\u000b`),
	'\f': []byte(`\f`),
	'\r': []byte(`\r`),
	14:   []byte(`\u000e`),
	15:   []byte(`\u000f`),
	16:   []byte(`\u0010`),
	17:   []byte(`\u0011`),
	18:   []byte(`\u0012`),
	19:   []byte(`\u0013`),
	20:   []byte(`\u0014`),
	21:   []byte(`\u0015`),
	22:   []byte(`\u0016`),
	23:   []byte(`\u0017`),
	24:   []byte(`\u0018`),
	25:   []byte(`\u0019`),
	26:   []byte(`\u001a`),
	27:   []byte(`\u001b`),
	28:   []byte(`\u001c`),
	29:   []byte(`\u001d`),
	30:   []byte(`\u001e`),
	31:   []byte(`\u001f`),
}

func appendFloat(buf []byte, f float64, bitSize int) []byte {
	abs := math.Abs(f)
	format := byte('f')
	if abs != 0 {
		if bitSize == 64 && (abs < 1e-6 || abs >= 1e21) ||
			bitSize == 32 && (float32(abs) < 1e-6 || float32(abs) >= 1e21) {
			format = 'e'
		}
	}

	buf = strconv.AppendFloat(buf, f, format, -1, bitSize)
	if format == 'e' {
		n := len(buf)
		if n >= 4 && buf[n-4] == 'e' && buf[n-3] == '-' && buf[n-2] == '0' {
			buf[n-2] = buf[n-1]
			buf = buf[:n-1]
		}
	}

	return buf
}

func appendQuote(buf, literal []byte) []byte {
	buf = append(buf, '"')

	i := 0
	for i+8 <= len(literal) {
		word := simd.LoadWord(literal[i:])
		if simd.HasEscapeChar(word) {
			buf = append(buf, literal[:i]...)
			return appendQuoteSlowFrom(buf, literal, i)
		}
		i += 8
	}

	for ; i < len(literal); i++ {
		c := literal[i]
		if c == '"' || c == '\\' || c < 32 {
			buf = append(buf, literal[:i]...)
			return appendQuoteSlowFrom(buf, literal, i)
		}
	}

	buf = append(buf, literal...)
	return append(buf, '"')
}

func appendQuoteSlowFrom(buf, literal []byte, start int) []byte {
	from := start
	i := start

	for i < len(literal) {
		for i+8 <= len(literal) {
			word := simd.LoadWord(literal[i:])
			if simd.HasEscapeChar(word) {
				break
			}
			i += 8
		}

		if i >= len(literal) {
			break
		}

		c := literal[i]
		if c == '"' || c == '\\' {
			buf = append(buf, literal[from:i]...)
			buf = append(buf, '\\', c)
			i++
			from = i
		} else if c < 32 {
			buf = append(buf, literal[from:i]...)
			buf = append(buf, controlCharEscapes[c]...)
			i++
			from = i
		} else {
			i++
		}
	}

	buf = append(buf, literal[from:]...)
	return append(buf, '"')
}
