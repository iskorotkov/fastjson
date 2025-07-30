package pooler_test

import (
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/iskorotkov/fastjson/pooler"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

func BenchmarkPooler(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		pooler := pooler.NewPooler()
		for range 100 {
			slice := reflect.ValueOf(new([]int)).Elem()
			slicePtr := &slice
			item := reflect.ValueOf(42)

			for range 100 {
				slicePtr = pooler.Append(slicePtr, &item)
			}
		}
	}
}
