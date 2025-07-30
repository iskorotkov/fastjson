package rustgo

import (
	"context"
	_ "embed"
	"unsafe"

	"github.com/iskorotkov/fastjson/tokenizer"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed wasm/tokenize_opt.wasm
var goWasmBytes []byte

//go:embed wasm/tokenize_rs_opt.wasm
var rustWasmBytes []byte

func NewGo(ctx context.Context) (func([]byte) []tokenizer.Token, func()) {
	r := wazero.NewRuntime(ctx)

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		panic(err)
	}

	mod, err := r.InstantiateWithConfig(ctx, goWasmBytes, wazero.NewModuleConfig().WithStartFunctions("_initialize"))
	if err != nil {
		panic(err)
	}

	tokenize := mod.ExportedFunction("tokenize")
	malloc := mod.ExportedFunction("malloc")
	free := mod.ExportedFunction("free")

	return func(b []byte) []tokenizer.Token {
			var mallocStack [1]uint64
			mallocStack[0] = uint64(len(b))
			if err := malloc.CallWithStack(ctx, mallocStack[:]); err != nil {
				panic(err)
			}

			defer func(stack [1]uint64) {
				if err := free.CallWithStack(ctx, stack[:]); err != nil {
					panic(err)
				}
			}(mallocStack)

			if !mod.Memory().Write(uint32(mallocStack[0]), b) {
				panic("failed to write memory")
			}

			var tokenizeStack [2]uint64
			tokenizeStack[1] = uint64(len(b))
			tokenizeStack[0] = mallocStack[0]
			if err := tokenize.CallWithStack(ctx, tokenizeStack[:]); err != nil {
				panic(err)
			}

			ptr := uint32(tokenizeStack[0] >> 32)
			size := uint32(tokenizeStack[0])

			defer func(stack [2]uint64) {
				var freeStack [1]uint64
				freeStack[0] = stack[0] >> 32
				if err := free.CallWithStack(ctx, freeStack[:]); err != nil {
					panic(err)
				}
			}(tokenizeStack)

			res, ok := mod.Memory().Read(ptr, size)
			if !ok {
				panic("failed to read memory")
			}

			return bytesToTokens(res)
		}, func() {
			if err := r.Close(ctx); err != nil {
				panic(err)
			}
		}
}

func NewRust(ctx context.Context) (func([]byte) []tokenizer.Token, func()) {
	r := wazero.NewRuntime(ctx)

	mod, err := r.Instantiate(ctx, rustWasmBytes)
	if err != nil {
		panic(err)
	}

	tokenize := mod.ExportedFunction("tokenize")
	malloc := mod.ExportedFunction("allocate")
	free := mod.ExportedFunction("deallocate")

	return func(b []byte) []tokenizer.Token {
			var mallocStack [1]uint64
			mallocStack[0] = uint64(len(b))
			if err := malloc.CallWithStack(ctx, mallocStack[:]); err != nil {
				panic(err)
			}

			defer func(stack [1]uint64, size uint64) {
				var freeStack [2]uint64
				freeStack[0] = stack[0]
				freeStack[1] = size
				if err := free.CallWithStack(ctx, stack[:]); err != nil {
					panic(err)
				}
			}(mallocStack, uint64(len(b)))

			if !mod.Memory().Write(uint32(mallocStack[0]), b) {
				panic("failed to write memory")
			}

			var tokenizeStack [2]uint64
			tokenizeStack[1] = uint64(len(b))
			tokenizeStack[0] = mallocStack[0]
			if err := tokenize.CallWithStack(ctx, tokenizeStack[:]); err != nil {
				panic(err)
			}

			ptr := uint32(tokenizeStack[0] >> 32)
			size := uint32(tokenizeStack[0])

			defer func(stack [2]uint64) {
				if err := free.CallWithStack(ctx, stack[:]); err != nil {
					panic(err)
				}
			}(tokenizeStack)

			res, ok := mod.Memory().Read(ptr, size)
			if !ok {
				panic("failed to read memory")
			}

			return bytesToTokens(res)
		}, func() {
			if err := r.Close(ctx); err != nil {
				panic(err)
			}
		}
}

func bytesToTokens(b []byte) []tokenizer.Token {
	slice := unsafe.SliceData(b)
	size := unsafe.Sizeof(tokenizer.Token{})
	return unsafe.Slice((*tokenizer.Token)(unsafe.Pointer(slice)), len(b)/int(size))
}
