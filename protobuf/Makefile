all: gen

.PHONY : gen

gen:
	protoc perf.golang.proto --go_out=plugins=grpc:./golang --go_opt=paths=source_relative
	protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf perf.gogo.proto --gogofast_out=plugins=grpc:./gogo --gogofast_opt=paths=source_relative
