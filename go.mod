module github.com/cloudwan/watchdog-sdk

go 1.16

require (
	github.com/cloudwan/edgelq-sdk v0.12.36
	github.com/cloudwan/goten-sdk v0.9.17
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.5.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/iancoleman/strcase v0.0.0-20180726023541-3605ed457bf7
	github.com/spf13/pflag v1.0.5
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.26.0
)

replace google.golang.org/protobuf => github.com/cloudwan/goten-protobuf v1.26.0
