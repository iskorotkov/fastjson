//go:build !go1.25 || !goexperiment.jsonv2

package benchmarks

import "testing"

func additionalMarshalBenchmarks(b *testing.B) {}
