package profile

import (
	"context"

	"go-zero-pj/test/internal/svc"
	"go-zero-pj/test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProfileLogic) GetProfile(req *types.GetRequest) (resp *types.GetResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
