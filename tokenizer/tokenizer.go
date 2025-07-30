package tokenizer

import (
	"io"

	"github.com/iskorotkov/fastjson/xstrconv"
)

var numberLiteralBytes = [256]byte{
	'0': 1, '1': 1, '2': 1, '3': 1, '4': 1,
	'5': 1, '6': 1, '7': 1, '8': 1, '9': 1,
	'-': 1, '+': 1, '.': 1, 'e': 1, 'E': 1,
}

var skipableBytes = [256]byte{
	' ': 1, '\n': 1, '\r': 1, '\t': 1, ',': 1, ':': 1,
}

var jumplist = [256]func(buf []byte) (int, Token){
	'{': func(buf []byte) (int, Token) {
		return 1, Token{Type: TokenTypeObjectStart}
	},
	'}': func(buf []byte) (int, Token) {
		return 1, Token{Type: TokenTypeObjectEnd}
	},
	'[': func(buf []byte) (int, Token) {
		return 1, Token{Type: TokenTypeArrayStart}
	},
	']': func(buf []byte) (int, Token) {
		return 1, Token{Type: TokenTypeArrayEnd}
	},
	'n': func(buf []byte) (int, Token) {
		if !isNull(buf) {
			panic(&InvalidTokenError{
				Expected: TokenTypeNull,
				Buf:      buf,
			})
		}
		return 4, Token{Type: TokenTypeNull}
	},
	't': func(buf []byte) (int, Token) {
		if !isTrue(buf) {
			panic(&InvalidTokenError{
				Expected: TokenTypeTrue,
				Buf:      buf,
			})
		}
		return 4, Token{Type: TokenTypeTrue}
	},
	'f': func(buf []byte) (int, Token) {
		if !isFalse(buf) {
			panic(&InvalidTokenError{
				Expected: TokenTypeFalse,
				Buf:      buf,
			})
		}
		return 5, Token{Type: TokenTypeFalse}
	},
	'"': stringToken,
	'0': numberToken,
	'1': numberToken,
	'2': numberToken,
	'3': numberToken,
	'4': numberToken,
	'5': numberToken,
	'6': numberToken,
	'7': numberToken,
	'8': numberToken,
	'9': numberToken,
	'-': numberToken,
	'+': numberToken,
}

func NewFromBytes(b []byte) Tokenizer {
	return Tokenizer{
		buf: b,
	}
}

func NewFromString(s string) Tokenizer {
	return NewFromBytes(xstrconv.StringToBytes(s))
}

func NewFromReader(r io.Reader) Tokenizer {
	buf := read(r)
	return NewFromBytes(buf)
}

type Tokenizer struct {
	buf         []byte
	hasPeeked   bool
	peekedToken Token
}

func (t *Tokenizer) All() []Token {
	tokens := make([]Token, 0, len(t.buf)/16)
	for {
		tok := t.Next()
		if tok.Type == TokenTypeEOF {
			return tokens
		}
		tokens = append(tokens, tok)
	}
}

func (t *Tokenizer) Peek() Token {
	if t.hasPeeked {
		return t.peekedToken
	}

	t.buf = skipBytes(t.buf)
	if len(t.buf) == 0 {
		return Token{Type: TokenTypeEOF}
	}

	f := jumplist[t.buf[0]]
	if f == nil {
		panic(&InvalidTokenError{
			Expected: TokenTypeObjectStart,
			Buf:      t.buf,
		})
	}

	skip, token := f(t.buf)
	t.buf = t.buf[skip:]
	t.hasPeeked = true
	t.peekedToken = token

	return token
}

func (t *Tokenizer) Next() Token {
	if t.hasPeeked {
		t.hasPeeked = false
		return t.peekedToken
	}

	t.buf = skipBytes(t.buf)
	if len(t.buf) == 0 {
		return Token{Type: TokenTypeEOF}
	}

	f := jumplist[t.buf[0]]
	if f == nil {
		panic(&InvalidTokenError{
			Expected: TokenTypeObjectStart,
			Buf:      t.buf,
		})
	}

	skip, token := f(t.buf)
	t.buf = t.buf[skip:]

	return token
}

func stringToken(buf []byte) (int, Token) {
	length := stringLiteral(buf)
	return length, Token{Type: TokenTypeQuotedLiteral, Literal: buf[:length]}
}

func numberToken(buf []byte) (int, Token) {
	length := numberLiteral(buf)
	return length, Token{Type: TokenTypeLiteral, Literal: buf[:length]}
}

func skipBytes(buf []byte) []byte {
	for len(buf) > 0 && skipableBytes[buf[0]] == 1 {
		buf = buf[1:]
	}
	return buf
}

func read(r io.Reader) []byte {
	sized, ok := r.(interface{ Size() int64 })
	if !ok {
		b, err := io.ReadAll(r)
		if err != nil {
			panic(&ReadError{Err: err})
		}
		return b
	}
	buf := make([]byte, sized.Size())
	n, err := r.Read(buf)
	if err != nil {
		panic(&ReadError{Err: err})
	}
	return buf[:n]
}

func numberLiteral(src []byte) int {
	for i, char := range src {
		if numberLiteralBytes[char] == 0 {
			return i
		}
	}
	return len(src)
}

func stringLiteral(src []byte) int {
	if len(src) == 0 {
		return 0
	}
	for i, b := range src[1:] {
		if b != '"' {
			continue
		}
		if oddEscapes(src, i) {
			continue
		}
		return i + 2
	}
	return len(src)
}

func oddEscapes(src []byte, i int) bool {
	_ = src[i]
	var odd bool
	for j := i; j >= 0; j-- {
		if src[j] == '\\' {
			odd = !odd
		} else {
			break
		}
	}
	return odd
}

func isNull(b []byte) bool {
	return len(b) >= 4 &&
		b[0] == 'n' &&
		b[1] == 'u' &&
		b[2] == 'l' &&
		b[3] == 'l'
}

func isTrue(b []byte) bool {
	return len(b) >= 4 &&
		b[0] == 't' &&
		b[1] == 'r' &&
		b[2] == 'u' &&
		b[3] == 'e'
}

func isFalse(b []byte) bool {
	return len(b) >= 5 &&
		b[0] == 'f' &&
		b[1] == 'a' &&
		b[2] == 'l' &&
		b[3] == 's' &&
		b[4] == 'e'
}
