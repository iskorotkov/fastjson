package protobuf

import (
	"os"
	"runtime"
	"runtime/debug"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

func BenchmarkUnmarshal(b *testing.B) {
	b.Run("protobuf/proto", func(b *testing.B) {
		data := GetSampleData()
		bytes, err := proto.Marshal(data)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		for b.Loop() {
			var resp UserManagementResponse
			if err := proto.Unmarshal(bytes, &resp); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})
}
