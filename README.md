# fastjson

Main focus:

1. low memory usage when unmarshaling
2. fast unmarshaling
3. low memory usage when marshaling
4. fast marshaling
5. everything else is secondary

Pros:

- lowest memory usage when unmarshaling
- fastest unmarshaling among JSON libraries

Cons:

- slower than Avro and Protobuf
- not JSON spec compliant
- not fully compatible with `encoding/json` or `encoding/json/v2`

## Future development

### Core features

Legend:

- HP - high priority (v0.1.0 requirement)
- MP - medium priority (v1.0.0 requirement)
- LP - low priority (v1.x.0 requirement).

Unmarshal:

  - HP optionally skip unknown fields
  - HP ignore unexported fields
  - HP support smaller number types
  - HP add tests for smaller number types
  - LP consider replacing stats with reusable type buffer and allocating the result in one go
  - LP add efficient read from io.Reader & benchmarks
  - LP support encoding/json/v2 API
  - MP validate json to prevent injection attacks
  - LP use fuzzing in benchmarks
  - LP rewrite encoding/json/v2 benchmark to be more performant due to new API
  - LP compare to https://flatbuffers.dev/
  - LP compare to https://github.com/go-faster/jx

Marshal:

  - HP support json tag options
  - HP support custom types: time.Time, time.Duration, net.IP, net.IPNet, url.URL
  - HP support custom marshalers
  - HP marshal maps with any keys as string maps
  - LP use more efficient interface alternatives
  - LP add marshaling to byte array
  - HP support smaller number types
  - HP add tests for smaller number types
  - HP add tests
  - LP add efficient write to io.Writer & benchmarks
  - LP support formatting time.Duration and some other types as both int and string
  - LP support encoding/json/v2 API
  - LP use fuzzing in benchmarks
  - LP rewrite encoding/json/v2 benchmark to be more performant due to new API
  - LP compare to https://flatbuffers.dev/
  - LP compare to https://github.com/go-faster/jx

Docs

- write benchmarks below
- fill README.md
- fill LICENSE
- point out issues with existing implementations - see:
    1. https://github.com/goccy/go-json/issues/549
    2. https://github.com/goccy/go-json/issues/535
    3. https://github.com/goccy/go-json/issues/534

### Additional features

CGO:

  - use rust for tokenization

Generate:

  - generate unmarshaling code
  - generate marshaling code

Monitoring

  - monitor newly created and updated JSON libs

Compliance:

  - add tests from https://cs.opensource.google/go/go/+/refs/tags/go1.25.1:src/encoding/json/v2_decode_test.go
  - add tests from https://cs.opensource.google/go/go/+/refs/tags/go1.25.1:src/encoding/json/v2_encode_test.go

Obscure optimization techniques:

  - return tokenizer from functions instead of passing it as a pointer
  - SIMD 256-bit and branchless might be faster https://www.alphaxiv.org/abs/1902.08318
  - rewrite tokenizer to pull model - add method `expect_next_token([',', ']'])`
  - parse and serialize JSON in parallel - needs research whether it will be useful in generic web apps
  - look at Rust serde and C# System.Text.Json to get inspiration
  - use arenas for marshal/unmarshal allocs optimization

## Benchmarks

See Github Actions for latest benchmark results.

Aggregated results on 2025-07-24:

```sh
goos: darwin
goarch: arm64
pkg: github.com/iskorotkov/fastjson/benchmarks
cpu: Apple M1 (Virtual)

B/op

* BenchmarkFastJSON-3        	   57513	     20390 ns/op	 592.30 MB/s	    4168 B/op	      23 allocs/op
BenchmarkHambaAvro-3   	        194119	      6195 ns/op	 530.24 MB/s	    6854 B/op	      28 allocs/op
BenchmarkEncodingJSON-3      	   23371	     50309 ns/op	 240.06 MB/s	    9490 B/op	     106 allocs/op
BenchmarkEncodingJSONV2-3    	   34453	     37154 ns/op	 325.05 MB/s	    9493 B/op	     106 allocs/op
BenchmarkProtobuf-3   	         60087	     17118 ns/op	 210.48 MB/s	   11176 B/op	     292 allocs/op
BenchmarkEasyJSON-3   	         48248	     30091 ns/op	 401.35 MB/s	   11216 B/op	     198 allocs/op
BenchmarkJSONIterFastest-3   	   49149	     24826 ns/op	 486.46 MB/s	   11597 B/op	     228 allocs/op
BenchmarkJSONIter-3          	   41522	     30799 ns/op	 392.13 MB/s	   12806 B/op	     324 allocs/op
BenchmarkFastJSONNaive-3     	   29126	     41362 ns/op	 291.98 MB/s	   13784 B/op	      54 allocs/op
BenchmarkGoJSON-3            	   63945	     20575 ns/op	 586.99 MB/s	   19658 B/op	      78 allocs/op

ns/op

BenchmarkHambaAvro-3   	        194119	      6195 ns/op	 530.24 MB/s	    6854 B/op	      28 allocs/op
BenchmarkProtobuf-3   	         60087	     17118 ns/op	 210.48 MB/s	   11176 B/op	     292 allocs/op
* BenchmarkFastJSON-3        	   57513	     20390 ns/op	 592.30 MB/s	    4168 B/op	      23 allocs/op
BenchmarkGoJSON-3            	   63945	     20575 ns/op	 586.99 MB/s	   19658 B/op	      78 allocs/op
BenchmarkJSONIterFastest-3   	   49149	     24826 ns/op	 486.46 MB/s	   11597 B/op	     228 allocs/op
BenchmarkEasyJSON-3   	         48248	     30091 ns/op	 401.35 MB/s	   11216 B/op	     198 allocs/op
BenchmarkJSONIter-3          	   41522	     30799 ns/op	 392.13 MB/s	   12806 B/op	     324 allocs/op
BenchmarkEncodingJSONV2-3    	   34453	     37154 ns/op	 325.05 MB/s	    9493 B/op	     106 allocs/op
BenchmarkFastJSONNaive-3     	   29126	     41362 ns/op	 291.98 MB/s	   13784 B/op	      54 allocs/op
BenchmarkEncodingJSON-3      	   23371	     50309 ns/op	 240.06 MB/s	    9490 B/op	     106 allocs/op
```
