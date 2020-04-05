# gRPC

## Installation Instructions

` brew install protobuf `

`go get -u google.golang.org/grpc`
`go get github.com/golang/protobuf/protoc-gen-go`

`protoc --proto_path=proto --go_out=plugins=grpc:proto service.proto `
