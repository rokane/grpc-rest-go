package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rokane/grpc-rest-template-go/pkg/api"
	"github.com/rokane/grpc-rest-template-go/pkg/pingv1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Setup gRPC API
	grpcServer := grpc.NewServer()
	pingServer := api.NewPingServer()
	pingv1.RegisterPingAPIServer(grpcServer, pingServer)
	reflection.Register(grpcServer)

	// Setup gRPC-REST Gateway
	mux := runtime.NewServeMux()
	httpServer := &http.Server{Handler: mux}
	if err := pingv1.RegisterPingAPIHandlerServer(context.Background(), mux, pingServer); err != nil {
		log.Fatal("Failed to setup http server")
	}

	// Start Routines
	g := new(errgroup.Group)
	g.Go(func() error { return listenAndServe("grpc", ":8080", grpcServer.Serve) })
	g.Go(func() error { return listenAndServe("http", ":8081", httpServer.Serve) })
	log.Fatal(g.Wait())
}

func listenAndServe(name, port string, serve func(net.Listener) error) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", port, err)
	}
	log.Println("Listening on server", name, port)
	return serve(listener)
}
