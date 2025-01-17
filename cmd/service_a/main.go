package main

import (
	"net/http"

	service_a_v1 "github.com/jimmyl02/connectrpc-boilerplate/proto/service_a/v1"
)

func main() {
	client := service_a_v1.NewTestClient(http.DefaultClient, "http://localhost:8080")
}
