package main

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/sureshkorlapati/grpc/sample-app/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	req := &proto.Request{A: 10, B: 20}

	res, err := client.Add(context.Background(), req)
	if err != nil {
		panic(err)
	}
	log.Infof("add result: %v", res.Result)

	res, err = client.Multiply(context.Background(), req)
	if err != nil {
		panic(err)
	}
	log.Infof("multiply result: %v", res.Result)

}
