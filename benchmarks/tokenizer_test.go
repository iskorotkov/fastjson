package benchmarks

import (
	"testing"

	"github.com/iskorotkov/fastjson/tokenizer"
)

func BenchmarkTokenizer(b *testing.B) {
	b.Run("next", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(Data)))

		for b.Loop() {
			tokens := tokenizer.NewFromBytes(Data)
			for {
				token := tokens.Next()
				if token.Type == tokenizer.TokenTypeEOF {
					break
				}
			}
		}
	})

	b.Run("all", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(Data)))

		for b.Loop() {
			tok := tokenizer.NewFromBytes(Data)
			tokens := tok.All()
			_ = tokens
		}
	})
}
