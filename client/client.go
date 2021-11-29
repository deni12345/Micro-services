package client

import (
	endpoint "lorem-grpc/endpoints"
	pb "lorem-grpc/pb"
	"lorem-grpc/service"
	"lorem-grpc/transport"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// Return new lorem_grpc service
func New(conn *grpc.ClientConn) service.Service {
	loremEndpoint := grpctransport.NewClient(
		conn, "Lorem", "LoremGenerate",
		transport.EncodeRequest,
		transport.DecodeRespond,
		pb.Response{},
	).Endpoint()

	return endpoint.Endpoints{
		LoremEndpoint: loremEndpoint,
	}
}
