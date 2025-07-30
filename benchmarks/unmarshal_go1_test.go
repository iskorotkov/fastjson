//go:build !go1.25 || !goexperiment.jsonv2

package benchmarks

import "testing"

func additionalBenchmarks(b *testing.B) {}
