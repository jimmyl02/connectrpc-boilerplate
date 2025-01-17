package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/jimmyl02/connectrpc-boilerplate/pkg/util"
	service_av1 "github.com/jimmyl02/connectrpc-boilerplate/proto/service_a/v1"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ServiceAServer struct {
	logger *zap.Logger
}

// Register implements v1.ServiceAServiceHandler.
func (s *ServiceAServer) Register(ctx context.Context, req *connect.Request[service_av1.RegisterRequest]) (*connect.Response[service_av1.RegisterResponse], error) {
	fmt.Println("Registering user: ", req.Msg.Email)

	if req.Msg.Name != nil {
		fmt.Println("Name: ", *req.Msg.Name)
	}
	if req.Msg.Age != nil {
		fmt.Println("Age: ", *req.Msg.Age)
	}

	return connect.NewResponse(&service_av1.RegisterResponse{
		Status: util.ToPtr(service_av1.ResponseStatus{
			Code: service_av1.ResponseCode_RESPONSE_CODE_SUCCESS,
		}),
		UserId: "test-user-id",
	}), nil
}

func main() {
	// initialize logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// declare server
	server := &ServiceAServer{
		logger: logger,
	}

	// begin serving
	mux := http.NewServeMux()
	path, handler := service_av1.NewServiceAServiceHandler(server)
	mux.Handle(path, handler)

	// listen and service
	logger.Info("Listening on :8080")
	err = http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		panic(err)
	}
}
