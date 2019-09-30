package main

import (
	"context"
	"fmt"
	"time"

	api "bitbucket.org/agrostar/grpc_test/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:1200", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	testClient := api.NewSomeStructClient(conn)

	a := api.MSG1{
		A: "keyur",
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	res, _ := testClient.SomeFunction(ctx, &a)
	fmt.Println(res)
}
