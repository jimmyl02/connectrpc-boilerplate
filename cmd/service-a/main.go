package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/jimmyl02/bazel-playground-connectrpc/proto/testproto"
)

func main() {
	client := testproto.NewTestClient(http.DefaultClient, "http://localhost:8080")
	age := int32(10)
	resp, err := client.SayHi(context.Background(), connect.NewRequest(&testproto.SayHiRequest{
		Name: "myname",
		Age:  &age,
	}))
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Msg.GetResponse())
	fmt.Println(resp.Header())
}
