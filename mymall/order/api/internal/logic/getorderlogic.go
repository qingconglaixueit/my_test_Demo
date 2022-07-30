package logic

import (
	"context"
	"my_test_demo/mymall/tenant/rpc/tenant"

	"my_test_demo/mymall/order/api/internal/svc"
	"my_test_demo/mymall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line
	rsp , err:=l.svcCtx.TenantRpc.GetTenant(l.ctx, &tenant.TidReq{Name: req.Name})
	if err != nil{
		return nil,err
	}
	return &types.OrderReply{
		Id: rsp.Id,
		Name: rsp.Name,
		Addr: rsp.Addr,
	},nil

}
