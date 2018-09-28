package main

import (
	"net"
	"log"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"


	"github.com/alphamu/goecho/service"
	"google.golang.org/grpc/reflection"
	gw "github.com/alphamu/goecho/proto"
	pb "github.com/alphamu/goecho/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var (
	echoEndpoint = flag.String("echo", "localhost:9090", "/")
)

const (
	port = ":9090"
)

func runGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterMessagesServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func runHttpServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessagesServiceServer(s, &service.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	flag.Parse()
	defer glog.Flush()
	go runHttpServer()
	if err := runGateway(); err != nil {
		glog.Fatal(err)
	}
}