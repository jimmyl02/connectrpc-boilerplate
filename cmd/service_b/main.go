package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/jimmyl02/connectrpc-boilerplate/pkg/util"
	service_av1 "github.com/jimmyl02/connectrpc-boilerplate/proto/service_a/v1"
)

func main() {
	client := service_av1.NewServiceAServiceClient(http.DefaultClient, "http://localhost:8080")

	// pretend we are also running our own connectrpc server
	resp, err := client.Register(context.Background(), connect.NewRequest(&service_av1.RegisterRequest{
		Email: "test-email@test.com",
		Name:  util.ToPtr("test-name"),
	}))
	if err != nil {
		panic(err)
	}

	fmt.Println("Response: ", resp.Msg.UserId)
}
