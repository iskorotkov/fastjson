package msgpack

import (
	"encoding/json"
	"os"
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/iskorotkov/fastjson/benchmarks"
	"github.com/vmihailenco/msgpack/v5"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

var msgpackData []byte

func init() {
	var jsonResp UserManagementResponse
	if err := json.Unmarshal(benchmarks.Data, &jsonResp); err != nil {
		panic(err)
	}

	var err error
	msgpackData, err = msgpack.Marshal(jsonResp)
	if err != nil {
		panic(err)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	b.Run("vmihailenco/msgpack/v5", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(msgpackData)))

		for b.Loop() {
			var resp UserManagementResponse
			if err := msgpack.Unmarshal(msgpackData, &resp); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})
}
