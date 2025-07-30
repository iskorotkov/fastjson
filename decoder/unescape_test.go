package decoder

import "testing"

func BenchmarkUnescapeJSON(b *testing.B) {
	cases := []struct {
		name string
		data []byte
	}{
		{"simple_escape", []byte(`hello \"world\"`)},
		{"multiple_escapes", []byte(`line1\nline2\tline3\r\n`)},
		{"unicode", []byte(`hello \u0048\u0065\u006c\u006c\u006f`)},
		{"surrogate_pair", []byte(`emoji: \uD83D\uDE00`)},
		{"mixed", []byte(`path: C:\\Users\\test\nname: \"John\"`)},
		{"long_with_few_escapes", []byte(`This is a very long string with only a few escape sequences like \"this\" and maybe a newline\n at the end`)},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(c.data)))
			for b.Loop() {
				_, _ = unescapeJSON(c.data)
			}
		})
	}
}
