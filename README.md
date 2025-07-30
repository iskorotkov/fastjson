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

## Future features

### Core features

- unmarshal:

  - optionally skip unknown fields
  - ignore unexported fields

- marshal:

  - implement fast marshaling
  - support json tag options
  - support custom types: time.Time, time.Duration, net.IP, net.IPNet, url.URL
  - support custom marshalers
  - add tests

### Docs

- write benchmarks below
- fill README.md
- fill LICENSE

### Additional features

- cgo:

  - use rust for tokenization

- generate:

  - generate unmarshaling code
  - generate marshaling code

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
