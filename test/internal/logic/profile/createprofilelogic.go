package profile

import (
	"context"

	"go-zero-pj/test/internal/svc"
	"go-zero-pj/test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProfileLogic {
	return &CreateProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProfileLogic) CreateProfile(req *types.CreateRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
