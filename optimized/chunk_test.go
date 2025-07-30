package optimized_test

import (
	"os"
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/iskorotkov/fastjson/optimized"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

func BenchmarkFor(b *testing.B) {
	b.Run("for", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			for _, b := range bytes {
				counter += int(b)
			}
		}
	})

	b.Run("for/unrolled", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			b := bytes
			for i := 0; i < len(bytes); i += 64 {
				counter += int(b[0] + b[1] + b[2] + b[3] + b[4] + b[5] + b[6] + b[7])
				counter += int(b[8] + b[9] + b[10] + b[11] + b[12] + b[13] + b[14] + b[15])
				counter += int(b[16] + b[17] + b[18] + b[19] + b[20] + b[21] + b[22] + b[23])
				counter += int(b[24] + b[25] + b[26] + b[27] + b[28] + b[29] + b[30] + b[31])
				counter += int(b[32] + b[33] + b[34] + b[35] + b[36] + b[37] + b[38] + b[39])
				counter += int(b[40] + b[41] + b[42] + b[43] + b[44] + b[45] + b[46] + b[47])
				counter += int(b[48] + b[49] + b[50] + b[51] + b[52] + b[53] + b[54] + b[55])
				counter += int(b[56] + b[57] + b[58] + b[59] + b[60] + b[61] + b[62] + b[63])
			}
		}
		_ = counter
	})

	b.Run("foreach", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			optimized.ForEach(bytes, func(b byte) {
				counter += int(b)
			})
		}
	})

	b.Run("for-each-chunk", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			optimized.ForEachChunk(bytes, func(b []byte) {
				for _, v := range b {
					counter += int(v)
				}
			})
		}
	})

	b.Run("for-each-chunk/unrolled", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			optimized.ForEachChunk(bytes, func(b []byte) {
				counter += int(b[0] + b[1] + b[2] + b[3] + b[4] + b[5] + b[6] + b[7])
				counter += int(b[8] + b[9] + b[10] + b[11] + b[12] + b[13] + b[14] + b[15])
				counter += int(b[16] + b[17] + b[18] + b[19] + b[20] + b[21] + b[22] + b[23])
				counter += int(b[24] + b[25] + b[26] + b[27] + b[28] + b[29] + b[30] + b[31])
				counter += int(b[32] + b[33] + b[34] + b[35] + b[36] + b[37] + b[38] + b[39])
				counter += int(b[40] + b[41] + b[42] + b[43] + b[44] + b[45] + b[46] + b[47])
				counter += int(b[48] + b[49] + b[50] + b[51] + b[52] + b[53] + b[54] + b[55])
				counter += int(b[56] + b[57] + b[58] + b[59] + b[60] + b[61] + b[62] + b[63])
			})
		}
	})
}

func BenchmarkFind(b *testing.B) {
	b.Run("for", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			for i, e := range bytes {
				if e == 42 {
					counter = i
				}
			}
		}
		_ = counter
	})

	b.Run("for/unrolled", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			s := bytes
			v := byte(42)
			for i := 0; i < len(bytes); i += 64 {
				found := s[i] == v || s[i+1] == v || s[i+2] == v || s[i+3] == v ||
					s[i+4] == v || s[i+5] == v || s[i+6] == v || s[i+7] == v ||
					s[i+8] == v || s[i+9] == v || s[i+10] == v || s[i+11] == v ||
					s[i+12] == v || s[i+13] == v || s[i+14] == v || s[i+15] == v ||
					s[i+16] == v || s[i+17] == v || s[i+18] == v || s[i+19] == v ||
					s[i+20] == v || s[i+21] == v || s[i+22] == v || s[i+23] == v ||
					s[i+24] == v || s[i+25] == v || s[i+26] == v || s[i+27] == v ||
					s[i+28] == v || s[i+29] == v || s[i+30] == v || s[i+31] == v ||
					s[i+32] == v || s[i+33] == v || s[i+34] == v || s[i+35] == v ||
					s[i+36] == v || s[i+37] == v || s[i+38] == v || s[i+39] == v ||
					s[i+40] == v || s[i+41] == v || s[i+42] == v || s[i+43] == v ||
					s[i+44] == v || s[i+45] == v || s[i+46] == v || s[i+47] == v ||
					s[i+48] == v || s[i+49] == v || s[i+50] == v || s[i+51] == v ||
					s[i+52] == v || s[i+53] == v || s[i+54] == v || s[i+55] == v ||
					s[i+56] == v || s[i+57] == v || s[i+58] == v || s[i+59] == v ||
					s[i+60] == v || s[i+61] == v || s[i+62] == v || s[i+63] == v
				if found {
					counter = i
					break
				}
			}
		}
		_ = counter
	})

	b.Run("find", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			counter = optimized.Find(bytes, 42)
		}
		_ = counter
	})

	b.Run("find-chunk", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			counter += optimized.FindChunk(bytes, 42)
		}
		_ = counter
	})

	b.Run("find-chunk/unrolled", func(b *testing.B) {
		bytes := make([]byte, 1_000_000)

		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(bytes)))

		var counter int
		for b.Loop() {
			counter += optimized.FindChunkUnrolled(bytes, 42)
		}
		_ = counter
	})
}
