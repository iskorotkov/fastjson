package decoder

import (
	"sync"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

var unescapePool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 256)
		return &b
	},
}

var escapeTable = [256]byte{
	'"':  '"',
	'\\': '\\',
	'/':  '/',
	'b':  '\b',
	'f':  '\f',
	'n':  '\n',
	'r':  '\r',
	't':  '\t',
}

var hexTable [256]int8

func init() {
	for i := range hexTable {
		hexTable[i] = -1
	}
	for i := '0'; i <= '9'; i++ {
		hexTable[i] = int8(i - '0')
	}
	for i := 'a'; i <= 'f'; i++ {
		hexTable[i] = int8(i - 'a' + 10)
	}
	for i := 'A'; i <= 'F'; i++ {
		hexTable[i] = int8(i - 'A' + 10)
	}
}

func parseHex4(b []byte) (rune, bool) {
	h0, h1, h2, h3 := hexTable[b[0]], hexTable[b[1]], hexTable[b[2]], hexTable[b[3]]
	if h0|h1|h2|h3 < 0 {
		return 0, false
	}
	return rune(h0)<<12 | rune(h1)<<8 | rune(h2)<<4 | rune(h3), true
}

func unescapeJSON(s []byte) (string, error) {
	bufp := unescapePool.Get().(*[]byte)
	buf := (*bufp)[:0]
	if cap(buf) < len(s) {
		buf = make([]byte, 0, len(s))
	}

	for len(s) > 0 {
		if s[0] != '\\' {
			i := 1
			for i < len(s) && s[i] != '\\' {
				i++
			}
			buf = append(buf, s[:i]...)
			s = s[i:]
			continue
		}

		if len(s) < 2 {
			*bufp = buf
			unescapePool.Put(bufp)
			return "", &EscapeError{Message: "incomplete escape sequence"}
		}

		if replacement := escapeTable[s[1]]; replacement != 0 {
			buf = append(buf, replacement)
			s = s[2:]
			continue
		}

		if s[1] == 'u' {
			if len(s) < 6 {
				*bufp = buf
				unescapePool.Put(bufp)
				return "", &EscapeError{Message: "incomplete unicode escape"}
			}
			rune1, ok := parseHex4(s[2:6])
			if !ok {
				*bufp = buf
				unescapePool.Put(bufp)
				return "", &EscapeError{Message: "invalid unicode escape"}
			}
			s = s[6:]

			if utf16.IsSurrogate(rune1) && len(s) >= 6 && s[0] == '\\' && s[1] == 'u' {
				if rune2, ok := parseHex4(s[2:6]); ok {
					if decoded := utf16.DecodeRune(rune1, rune2); decoded != unicode.ReplacementChar {
						buf = utf8.AppendRune(buf, decoded)
						s = s[6:]
						continue
					}
				}
			}
			buf = utf8.AppendRune(buf, rune1)
			continue
		}

		*bufp = buf
		unescapePool.Put(bufp)
		return "", &EscapeError{Message: "invalid escape character: \\" + string(s[1])}
	}

	result := string(buf)
	*bufp = buf
	unescapePool.Put(bufp)
	return result, nil
}
