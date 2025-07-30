//go:build goexperiment.jsonv2

package benchmarks

import (
	"encoding/json/v2"
	"testing"
)

func benchmarkMarshalEncodingJSONV2(b *testing.B) {
	b.Helper()
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		res, err := json.Marshal(Response)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		b.SetBytes(int64(len(res)))
	}
}

func benchmarkUnmarshalEncodingJSONV2(b *testing.B) {
	b.Helper()
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(Data)))

	for b.Loop() {
		var resp UserManagementResponse
		if err := json.Unmarshal(Data, &resp); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
