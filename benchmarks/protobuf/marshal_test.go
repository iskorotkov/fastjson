package protobuf

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

var Response = GetSampleData()

func BenchmarkMarshal(b *testing.B) {
	b.Run("protobuf/proto", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for b.Loop() {
			res, err := proto.Marshal(Response)
			if err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
			b.SetBytes(int64(len(res)))
		}
	})
}
