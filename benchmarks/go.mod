module github.com/iskorotkov/fastjson/benchmarks

go 1.24

replace github.com/iskorotkov/fastjson => ../

require (
	github.com/goccy/go-json v0.10.5
	github.com/hamba/avro/v2 v2.29.0
	github.com/iskorotkov/fastjson v0.0.0-00010101000000-000000000000
	github.com/json-iterator/go v1.1.12
	github.com/mailru/easyjson v0.9.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
)
