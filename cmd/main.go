package main

import (
	"flag"
	"fmt"
	"log"
	"lorem-grpc/service"
	"strconv"
	"time"

	grpcClient "lorem-grpc/client"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	grpcAddr := flag.String("addr", ":8081",
		"gRPC address")
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(),
		grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()

	loremService := grpcClient.New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)

	switch cmd {
	case "lorem":
		var requestType, minStr, maxStr string

		requestType, args = pop(args)
		minStr, args = pop(args)
		maxStr, args = pop(args)

		min, _ := strconv.Atoi(minStr)
		max, _ := strconv.Atoi(maxStr)
		lorem(ctx, loremService, requestType, min, max)
	default:
		log.Fatalln("unknown command", cmd)
	}
}

// parse command line argument one by one
func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

// call lorem service
func lorem(ctx context.Context, service service.Service, requestType string, min int, max int) {
	log.Println(requestType, " ", min, " ", max)
	mesg, err := service.LoremGenerate(ctx, requestType, min, max)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(mesg)
}
