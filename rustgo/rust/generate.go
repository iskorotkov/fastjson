package rust

//go:generate rustc -o ../wasm/tokenize_rs.wasm --target wasm32-wasip1 --crate-type cdylib -C debuginfo=0 -C opt-level=3 -C codegen-units=1 -C lto=true lib.rs
//go:generate wasm-opt -o ../wasm/tokenize_rs_opt.wasm -O4 --traps-never-happen --fast-math --enable-simd --enable-bulk-memory --enable-tail-call ../wasm/tokenize_rs.wasm
