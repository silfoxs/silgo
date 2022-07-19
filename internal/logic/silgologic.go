package logic

import (
	"context"

	"silgo/internal/svc"
	"silgo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SilgoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSilgoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SilgoLogic {
	return &SilgoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SilgoLogic) Silgo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: "Hello, world!",
	}, nil
}
