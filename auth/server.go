package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/totoview/centrex/pb"
)

type authServer struct {
}

func (s *authServer) Login(ctx context.Context, req pb.Login) (pb.LoginRsp, error) {
	return pb.LoginRsp{}, nil
}

func MakeLoginEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.Login)
		return pb.LoginRsp{RequestId: req.RequestId, ErrorCode: 0}, nil
	}
}
