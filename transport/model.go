package transport

import (
	"context"
	endpoint "lorem-grpc/endpoints"
	pb "lorem-grpc/pb"
)

func EncodeRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoint.Request)
	return &pb.Request{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

func DecodeRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.Request)
	return endpoint.Request{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

func EncodeRespond(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.Respond)
	return &pb.Response{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}

func DecodeRespond(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.Response)
	return endpoint.Respond{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}
