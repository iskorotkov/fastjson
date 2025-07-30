//go:build go1.25 && goexperiment.jsonv2

package benchmarks

import (
	"encoding/json/v2"
	"testing"
)

func additionalMarshalBenchmarks(b *testing.B) {
	b.Run("encoding/json/v2", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for b.Loop() {
			res, err := json.Marshal(Response)
			if err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
			b.SetBytes(int64(len(res)))
		}
	})
}
