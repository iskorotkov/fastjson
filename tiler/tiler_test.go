package tiler_test

import (
	"bytes"
	"strconv"
	"testing"
	"time"

	"github.com/iskorotkov/fastjson/tiler"
)

const tokens = 10000

func BenchmarkTiler(b *testing.B) {
	b.Run("PutString", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		tiler := tiler.New()
		for b.Loop() {
			tiler.Reset()
			for range tokens {
				tiler.PutString("Hello, World!")
			}
		}
	})

	b.Run("PutQuotedString", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		tiler := tiler.New()
		for b.Loop() {
			tiler.Reset()
			for range tokens {
				tiler.PutQuotedString("Hello, \"World\"!")
			}
		}
	})

	b.Run("PutInt", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		tiler := tiler.New()
		for b.Loop() {
			tiler.Reset()
			for range tokens {
				tiler.PutInt(1234567890)
			}
		}
	})

	b.Run("PutFloat", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		tiler := tiler.New()
		for b.Loop() {
			tiler.Reset()
			for range tokens {
				tiler.PutFloat(3.14159265359)
			}
		}
	})

	b.Run("PutBool", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		tiler := tiler.New()
		for b.Loop() {
			tiler.Reset()
			for range tokens {
				tiler.PutBool(true)
			}
		}
	})

	b.Run("PutDuration", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		tiler := tiler.New()
		d := 5 * time.Second
		for b.Loop() {
			tiler.Reset()
			for range tokens {
				tiler.PutDuration(d)
			}
		}
	})

	b.Run("Mixed", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		tiler := tiler.New()
		for b.Loop() {
			tiler.Reset()

			for range tokens {
				tiler.PutString("{")
				tiler.PutQuotedString("name")
				tiler.PutString(":")
				tiler.PutQuotedString("John Doe")
				tiler.PutString(",")
				tiler.PutQuotedString("age")
				tiler.PutString(":")
				tiler.PutInt(30)
				tiler.PutString(",")
				tiler.PutQuotedString("active")
				tiler.PutString(":")
				tiler.PutBool(true)
				tiler.PutString("}")
			}
		}
	})
}

func BenchmarkBytesBuffer(b *testing.B) {
	b.Run("WriteString", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := new(bytes.Buffer)
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.WriteString("Hello, World!")
			}
		}
	})

	b.Run("WriteQuotedString", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := new(bytes.Buffer)
		for b.Loop() {
			buf.Reset()

			for range tokens {
				buf.WriteByte('"')
				for _, r := range "Hello, \"World\"!" {
					switch r {
					case '"':
						buf.WriteString(`\"`)
					case '\\':
						buf.WriteString(`\\`)
					default:
						buf.WriteRune(r)
					}
				}
				buf.WriteByte('"')
			}
		}
	})

	b.Run("WriteInt", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := new(bytes.Buffer)
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.WriteString(strconv.FormatInt(1234567890, 10))
			}
		}
	})

	b.Run("WriteFloat", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := new(bytes.Buffer)
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.WriteString(strconv.FormatFloat(3.14159265359, 'g', -1, 64))
			}
		}
	})

	b.Run("WriteBool", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := new(bytes.Buffer)
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.WriteString(strconv.FormatBool(true))
			}
		}
	})

	b.Run("WriteDuration", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := new(bytes.Buffer)
		d := 5 * time.Second
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.WriteString(d.String())
			}
		}
	})

	b.Run("Mixed", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := new(bytes.Buffer)
		for b.Loop() {
			buf.Reset()

			for range tokens {
				buf.WriteString("{")
				buf.WriteByte('"')
				buf.WriteString("name")
				buf.WriteByte('"')
				buf.WriteString(":")
				buf.WriteByte('"')
				buf.WriteString("John Doe")
				buf.WriteByte('"')
				buf.WriteString(",")
				buf.WriteByte('"')
				buf.WriteString("age")
				buf.WriteByte('"')
				buf.WriteString(":")
				buf.WriteString(strconv.FormatInt(30, 10))
				buf.WriteString(",")
				buf.WriteByte('"')
				buf.WriteString("active")
				buf.WriteByte('"')
				buf.WriteString(":")
				buf.WriteString(strconv.FormatBool(true))
				buf.WriteString("}")
			}
		}
	})
}
