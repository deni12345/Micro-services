package endpoint

import (
	"context"
	"errors"
	"log"
	"lorem-grpc/service"

	"github.com/go-kit/kit/endpoint"
)

type Request struct {
	RequestType string
	Min         int32
	Max         int32
}

type Respond struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	log.Panicln("dd")
	return Endpoints{
		LoremEndpoint: MakeLoremEndpoint(s),
	}
}

func MakeLoremEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)

		var min, max int

		min = int(req.Min)
		max = int(req.Max)
		txt, err := svc.LoremGenerate(ctx, req.RequestType, min, max)
		if err != nil {
			return nil, err
		}

		return Respond{Message: txt}, nil
	}
}

func (e Endpoints) LoremGenerate(ctx context.Context, requestType string, min, max int) (string, error) {
	req := Request{
		RequestType: requestType,
		Min:         int32(min),
		Max:         int32(max),
	}

	resp, err := e.LoremEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	loremResp := resp.(Respond)
	if loremResp.Err != "" {
		return "", errors.New(loremResp.Err)
	}

	return loremResp.Message, nil
}
