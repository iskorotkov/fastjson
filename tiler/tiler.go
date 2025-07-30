package tiler

import (
	"strconv"
	"time"

	"github.com/iskorotkov/fastjson/xstrconv"
)

func New() Tiler {
	return Tiler{
		buf: make([]byte, 0, 8192),
	}
}

type Tiler struct {
	buf []byte
}

func (t *Tiler) PutComma() {
	t.buf = append(t.buf, ',')
}

func (t *Tiler) PutColon() {
	t.buf = append(t.buf, ':')
}

func (t *Tiler) PutObjectStart() {
	t.buf = append(t.buf, '{')
}

func (t *Tiler) PutObjectEnd() {
	t.buf = append(t.buf, '}')
}

func (t *Tiler) PutArrayStart() {
	t.buf = append(t.buf, '[')
}

func (t *Tiler) PutArrayEnd() {
	t.buf = append(t.buf, ']')
}

func (t *Tiler) PutString(s string) {
	t.PutBytes(xstrconv.StringToBytes(s))
}

func (t *Tiler) PutQuotedString(s string) {
	t.PutQuotedBytes(xstrconv.StringToBytes(s))
}

func (t *Tiler) PutBytes(b []byte) {
	t.buf = append(t.buf, b...)
}

func (t *Tiler) PutQuotedBytes(b []byte) {
	t.buf = AppendQuote(t.buf, b)
}

func (t *Tiler) PutInt(i int64) {
	t.buf = strconv.AppendInt(t.buf, i, 10)
}

func (t *Tiler) PutUint(u uint64) {
	t.buf = strconv.AppendUint(t.buf, u, 10)
}

func (t *Tiler) PutFloat(f float64) {
	t.buf = strconv.AppendFloat(t.buf, f, 'g', -1, 64)
}

func (t *Tiler) PutBool(b bool) {
	if b {
		t.buf = append(t.buf, 't', 'r', 'u', 'e')
	} else {
		t.buf = append(t.buf, 'f', 'a', 'l', 's', 'e')
	}
}

func (t *Tiler) PutDuration(d time.Duration) {
	t.PutQuotedString(d.String())
}

func (t *Tiler) PutNull() {
	t.buf = append(t.buf, 'n', 'u', 'l', 'l')
}

func (t *Tiler) Clone() []byte {
	res := make([]byte, len(t.buf))
	copy(res, t.buf)
	return res
}

func (t *Tiler) Reset() {
	t.buf = t.buf[:0]
}

func AppendQuote(buf, literal []byte) []byte {
	buf = append(buf, '"')

	from := 0
	for i, c := range literal {
		if c == '"' || c == '\\' {
			buf = append(buf, literal[from:i]...)
			buf = append(buf, '\\', c)
			from = i + 1
		}
	}

	buf = append(buf, literal[from:]...)
	return append(buf, '"')
}
