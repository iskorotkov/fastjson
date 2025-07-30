package easyjson

import (
	"os"
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/iskorotkov/fastjson/benchmarks"
	"github.com/mailru/easyjson"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

func BenchmarkUnmarshal(b *testing.B) {
	b.Run("mailru/easyjson", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(benchmarks.Data)))

		for b.Loop() {
			var resp UserManagementResponse
			if err := easyjson.Unmarshal(benchmarks.Data, &resp); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})
}
