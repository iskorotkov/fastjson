//go:build !goexperiment.jsonv2

package benchmarks

import "testing"

func benchmarkMarshalEncodingJSONV2(b *testing.B) {
	b.Helper()
	b.Skip()
}

func benchmarkUnmarshalEncodingJSONV2(b *testing.B) {
	b.Helper()
	b.Skip()
}
