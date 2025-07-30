package buffer_test

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"testing"
	"time"

	"github.com/iskorotkov/fastjson/buffer"
)

const tokens = 10000

func BenchmarkBuffer(b *testing.B) {
	b.Run("PutString", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := buffer.New()
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.PutString("Hello, World!")
			}
		}
	})

	b.Run("PutQuotedString", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := buffer.New()
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.PutQuotedString("Hello, \"World\"!")
			}
		}
	})

	b.Run("PutInt", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := buffer.New()
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.PutInt(1234567890)
			}
		}
	})

	b.Run("PutFloat", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := buffer.New()
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.PutFloat(3.14159265359)
			}
		}
	})

	b.Run("PutBool", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := buffer.New()
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.PutBool(true)
			}
		}
	})

	b.Run("PutDuration", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := buffer.New()
		d := 5 * time.Second
		for b.Loop() {
			buf.Reset()
			for range tokens {
				buf.PutDuration(d)
			}
		}
	})

	b.Run("Mixed", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		buf := buffer.New()
		for b.Loop() {
			buf.Reset()

			for range tokens {
				buf.PutString("{")
				buf.PutQuotedString("name")
				buf.PutString(":")
				buf.PutQuotedString("John Doe")
				buf.PutString(",")
				buf.PutQuotedString("age")
				buf.PutString(":")
				buf.PutInt(30)
				buf.PutString(",")
				buf.PutQuotedString("active")
				buf.PutString(":")
				buf.PutBool(true)
				buf.PutString("}")
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

func BenchmarkPutBase64(b *testing.B) {
	sizes := []struct {
		name string
		size int
	}{
		{"16B", 16},
		{"64B", 64},
		{"256B", 256},
		{"1KB", 1024},
		{"4KB", 4096},
		{"16KB", 16384},
	}

	for _, s := range sizes {
		data := make([]byte, s.size)
		for i := range data {
			data[i] = byte(i % 256)
		}

		b.Run(s.name, func(b *testing.B) {
			buf := buffer.New()
			b.ReportAllocs()
			b.ResetTimer()

			for b.Loop() {
				buf.Reset()
				buf.PutBase64(data)
			}
		})
	}
}

func TestPutBase64(t *testing.T) {
	testCases := []struct {
		name  string
		input []byte
	}{
		{"empty", []byte{}},
		{"small", []byte("hello")},
		{"medium", make([]byte, 100)},
		{"large", make([]byte, 10000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf := buffer.New()
			buf.PutBase64(tc.input)

			expected := base64.StdEncoding.EncodeToString(tc.input)
			actual := string(buf.Clone())

			if actual != expected {
				t.Errorf("expected %q, got %q", expected, actual)
			}
		})
	}
}
