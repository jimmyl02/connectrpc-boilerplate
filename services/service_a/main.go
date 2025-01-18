package main

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"

	"connectrpc.com/connect"
	"github.com/jimmyl02/connectrpc-boilerplate/pkg/cors"
	"github.com/jimmyl02/connectrpc-boilerplate/pkg/util"
	service_av1 "github.com/jimmyl02/connectrpc-boilerplate/proto/service_a/v1"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ServiceAServer struct {
	logger *zap.Logger

	// internal constructs
	count int32
}

// GetCount implements v1.ServiceAServiceHandler.
func (s *ServiceAServer) GetCount(context.Context, *connect.Request[service_av1.GetCountRequest]) (*connect.Response[service_av1.GetCountResponse], error) {
	return connect.NewResponse(&service_av1.GetCountResponse{
		Status: util.ToPtr(service_av1.ResponseStatus{
			Code: service_av1.ResponseCode_RESPONSE_CODE_SUCCESS,
		}),
		Count: atomic.LoadInt32(&s.count),
	}), nil
}

// NewCount implements v1.ServiceAServiceHandler.
func (s *ServiceAServer) NewCount(context.Context, *connect.Request[service_av1.NewCountRequest]) (*connect.Response[service_av1.NewCountResponse], error) {
	atomic.AddInt32(&s.count, 1)
	return connect.NewResponse(&service_av1.NewCountResponse{
		Status: util.ToPtr(service_av1.ResponseStatus{
			Code: service_av1.ResponseCode_RESPONSE_CODE_SUCCESS,
		}),
	}), nil
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

// buildServer builds the server handler and wraps it with the appropriate middleware
func buildServer(
	logger *zap.Logger,
) (path string, handler http.Handler) {
	// declare server
	server := &ServiceAServer{
		logger: logger,
	}

	// construct the path and handler
	path, handler = service_av1.NewServiceAServiceHandler(server)

	// wrap the handler with middleware
	handler = cors.WithCORS(handler, []string{"http://localhost:5173"})

	// return the path and handler
	return path, handler
}

func main() {
	// initialize logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// declare the mux
	mux := http.NewServeMux()
	path, handler := buildServer(logger)
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
