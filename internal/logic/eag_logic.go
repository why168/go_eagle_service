package logic

import (
	"context"
	"sync"

	"go_eagle_service/internal/svc"
	"go_eagle_service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var once sync.Once
var logic *EagLogic

type EagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EagLogic {
	once.Do(func() {
		logic = &EagLogic{
			Logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	})
	return logic
}

func (l *EagLogic) GoEagleService(req *types.Request) (*types.Response, error) {
	response := &types.Response{
		Message: "Hello, " + req.Name,
	}
	return response, nil
}
