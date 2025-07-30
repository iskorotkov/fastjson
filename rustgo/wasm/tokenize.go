package main

// #include <stdlib.h>
import "C"

import (
	"unsafe"

	"github.com/iskorotkov/fastjson/tokenizer"
)

//go:generate tinygo build -o tokenize.wasm -buildmode=c-shared -target=wasip1 -scheduler=none -no-debug -gc=leaking -opt=2 tokenize.go
//go:generate wasm-opt -o tokenize_opt.wasm -O4 --traps-never-happen --fast-math --enable-simd --enable-bulk-memory --enable-tail-call tokenize.wasm

var _ = tokenize

func main() {}

//go:wasmexport tokenize
func tokenize(ptr, size uint32) uint64 {
	b := importSlice[byte](ptr, size)
	t := tokenizer.NewTokenizerFromBytes(b)

	var tokens []tokenizer.Token
	for {
		tok := t.Next()
		if tok.Type == tokenizer.TokenTypeEOF {
			return exportSlice(tokens)
		}
		tokens = append(tokens, tok)
	}
}

func importSlice[T any](ptr, size uint32) []T {
	var def T
	return unsafe.Slice((*T)(unsafe.Pointer(uintptr(ptr))), size/uint32(unsafe.Sizeof(def)))
}

func exportSlice[T any](tokens []T) uint64 {
	var def T
	size := C.ulong(len(tokens) * int(unsafe.Sizeof(def)))
	ptr := unsafe.Pointer(C.malloc(size))
	copy(unsafe.Slice((*T)(ptr), len(tokens)), tokens)
	return uint64(uintptr(ptr))<<32 | uint64(size)
}
