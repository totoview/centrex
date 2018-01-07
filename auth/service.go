package auth

import (
	"context"

	"github.com/totoview/centrex/pb"
)

// Service is authentication service
type Service interface {
	Login(context.Context, pb.Login) (pb.LoginRsp, error)
}
