package buffer_test

import (
	"strings"
	"testing"

	"github.com/iskorotkov/fastjson/buffer"
)

func BenchmarkAppendQuote(b *testing.B) {
	b.Run("ShortClean", func(b *testing.B) {
		buf := buffer.New()
		s := "name"
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			buf.Reset()
			buf.PutQuotedString(s)
		}
	})

	b.Run("LongClean", func(b *testing.B) {
		buf := buffer.New()
		s := strings.Repeat("abcdefghij", 10)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			buf.Reset()
			buf.PutQuotedString(s)
		}
	})

	b.Run("VeryLongClean", func(b *testing.B) {
		buf := buffer.New()
		s := strings.Repeat("abcdefghij", 100)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			buf.Reset()
			buf.PutQuotedString(s)
		}
	})

	b.Run("EarlyEscape", func(b *testing.B) {
		buf := buffer.New()
		s := "\nsome data here"
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			buf.Reset()
			buf.PutQuotedString(s)
		}
	})

	b.Run("LateEscape", func(b *testing.B) {
		buf := buffer.New()
		s := strings.Repeat("abcdefghij", 10) + "\n"
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			buf.Reset()
			buf.PutQuotedString(s)
		}
	})

	b.Run("MultipleEscapes", func(b *testing.B) {
		buf := buffer.New()
		s := "hello\tworld\nfoo\"bar\\baz"
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			buf.Reset()
			buf.PutQuotedString(s)
		}
	})

	b.Run("AllEscapes", func(b *testing.B) {
		buf := buffer.New()
		s := "\"\\\t\n\r\b\f"
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			buf.Reset()
			buf.PutQuotedString(s)
		}
	})

	b.Run("Empty", func(b *testing.B) {
		buf := buffer.New()
		s := ""
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			buf.Reset()
			buf.PutQuotedString(s)
		}
	})
}

func TestAppendQuoteCorrectness(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", `""`},
		{"simple", "hello", `"hello"`},
		{"quote", `hello"world`, `"hello\"world"`},
		{"backslash", `hello\world`, `"hello\\world"`},
		{"newline", "hello\nworld", `"hello\nworld"`},
		{"tab", "hello\tworld", `"hello\tworld"`},
		{"carriage return", "hello\rworld", `"hello\rworld"`},
		{"backspace", "hello\bworld", `"hello\bworld"`},
		{"formfeed", "hello\fworld", `"hello\fworld"`},
		{"control char", "hello\x00world", `"hello\u0000world"`},
		{"multiple escapes", "\"\\\t\n", `"\"\\\t\n"`},
		{"long clean", strings.Repeat("a", 100), `"` + strings.Repeat("a", 100) + `"`},
		{"escape at end", "hello\n", `"hello\n"`},
		{"escape at start", "\nhello", `"\nhello"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := buffer.New()
			buf.PutQuotedString(tt.input)
			got := string(buf.Clone())
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
