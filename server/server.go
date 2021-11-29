package main

import (
	"flag"
	"fmt"
	"log"
	endpoint "lorem-grpc/endpoints"
	pb "lorem-grpc/pb"
	"lorem-grpc/service"
	"lorem-grpc/transport"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	gRPCAddr := flag.String("grpc", ":8081", "gRPC listen address")
	flag.Parse()

	var svc service.Service
	svc = service.LoremService{}
	errChan := make(chan error)

	log.Println("server")
	// creating Endpoints struct
	endpoints := endpoint.MakeEndpoints(svc)
	// execute grpc server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := transport.NewGRPCServer(endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterLoremServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
