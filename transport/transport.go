package transport

import (
	"context"

	endpoint "lorem-grpc/endpoints"
	pb "lorem-grpc/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	lorem grpctransport.Handler
	pb.UnimplementedLoremServer
}

func (s *grpcServer) LoremGenerate(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	_, resp, err := s.lorem.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Response), nil
}

func NewGRPCServer(endpoint endpoint.Endpoints) pb.LoremServer {
	return &grpcServer{
		lorem: grpctransport.NewServer(
			endpoint.LoremEndpoint,
			DecodeRequest,
			EncodeRespond,
		),
	}
}
