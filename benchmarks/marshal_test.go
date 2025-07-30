package benchmarks

import (
	"encoding/json"
	"testing"

	gojson "github.com/goccy/go-json"
	"github.com/iskorotkov/fastjson"
	jsoniter "github.com/json-iterator/go"
)

func BenchmarkMarshal(b *testing.B) {
	additionalMarshalBenchmarks(b)

	b.Run("iskorotkov/fastjson", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		enc := fastjson.NewEncoder[UserManagementResponse]()
		for b.Loop() {
			res, err := enc.Marshal(Response)
			if err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
			b.SetBytes(int64(len(res)))
		}
	})

	b.Run("iskorotkov/fastjson/naive", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for b.Loop() {
			enc := fastjson.NewEncoder[UserManagementResponse]()
			res, err := enc.Marshal(Response)
			if err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
			b.SetBytes(int64(len(res)))
		}
	})

	b.Run("encoding/json", func(b *testing.B) {
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

	b.Run("json-iterator/go", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for b.Loop() {
			res, err := jsoniter.Marshal(Response)
			if err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
			b.SetBytes(int64(len(res)))
		}
	})

	b.Run("json-iterator/go/fastest", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for b.Loop() {
			res, err := jsoniter.ConfigFastest.Marshal(Response)
			if err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
			b.SetBytes(int64(len(res)))
		}
	})

	b.Run("goccy/go-json", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for b.Loop() {
			res, err := gojson.Marshal(Response)
			if err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
			b.SetBytes(int64(len(res)))
		}
	})
}
