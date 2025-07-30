package optimized

const ChunkSize = 64

func ForEach[S ~[]T, T any](s S, f func(T)) {
	for _, e := range s {
		f(e)
	}
}

func ForEachChunk[S ~[]T, T any](s S, f func(S)) {
	for i := 0; i < len(s); i += ChunkSize {
		if i+ChunkSize > len(s) {
			var s2 [ChunkSize]T
			copy(s2[:], s[i:])
			f(s2[:])
		} else {
			f(s[i : i+ChunkSize])
		}
	}
}

func Find[S ~[]T, T comparable](s S, v T) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}

func FindChunk[S ~[]T, T comparable](s S, v T) int {
	for i := 0; i < len(s); i += ChunkSize {
		if i+ChunkSize > len(s) {
			var s2 [ChunkSize]T
			copy(s2[:], s[i:])
			if idx := Find(s2[:], v); idx != -1 {
				return i + idx
			}
		} else {
			if idx := Find(s[i:i+ChunkSize], v); idx != -1 {
				return i + idx
			}
		}
	}
	return -1
}

func FindChunkUnrolled[S ~[]T, T comparable](s S, v T) int {
	for i := 0; i < len(s); i += ChunkSize {
		if i+ChunkSize > len(s) {
			var s2 [ChunkSize]T
			copy(s2[:], s[i:])
			if idx := Find(s2[:], v); idx != -1 {
				return i + idx
			}
		} else {
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
				return i
			}
		}
	}
	return -1
}
