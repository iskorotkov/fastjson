//go:build go1.25 && goexperiment.jsonv2

package benchmarks

import (
	"encoding/json/v2"
	"testing"
)

func additionalUnmarshalBenchmarks(b *testing.B) {
	b.Run("encoding/json/v2", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(Data)))

		for b.Loop() {
			var resp UserManagementResponse
			if err := json.Unmarshal(Data, &resp); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})
}
