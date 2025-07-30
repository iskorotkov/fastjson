package benchmarks

import (
	"encoding/json"
	"os"
	"runtime"
	"runtime/debug"
	"testing"

	gojson "github.com/goccy/go-json"
	"github.com/iskorotkov/fastjson"
	jsoniter "github.com/json-iterator/go"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

func BenchmarkUnmarshal(b *testing.B) {
	additionalBenchmarks(b)

	b.Run("iskorotkov/fastjson", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(Data)))

		dec := fastjson.NewDecoder[UserManagementResponse]()
		for b.Loop() {
			var result UserManagementResponse
			if err := dec.Unmarshal(Data, &result); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})

	b.Run("iskorotkov/fastjson/naive", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(Data)))

		for b.Loop() {
			var result UserManagementResponse
			if err := fastjson.NewDecoder[UserManagementResponse]().Unmarshal(Data, &result); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})

	b.Run("encoding/json", func(b *testing.B) {
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

	b.Run("json-iterator/go", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(Data)))

		for b.Loop() {
			var resp UserManagementResponse
			if err := jsoniter.Unmarshal(Data, &resp); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})

	b.Run("json-iterator/go/fastest", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(Data)))

		for b.Loop() {
			var resp UserManagementResponse
			if err := jsoniter.ConfigFastest.Unmarshal(Data, &resp); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})

	b.Run("goccy/go-json", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(Data)))

		for b.Loop() {
			var resp UserManagementResponse
			if err := gojson.Unmarshal(Data, &resp); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})
}
