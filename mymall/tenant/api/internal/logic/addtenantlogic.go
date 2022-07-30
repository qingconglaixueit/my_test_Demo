package logic

import (
	"context"
	"fmt"
	"my_test_demo/mymall/tenant/rpc/tenant"

	"my_test_demo/mymall/tenant/api/internal/svc"
	"my_test_demo/mymall/tenant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTenantLogic {
	return &AddTenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTenantLogic) AddTenant(req *types.AddTenantReq) (*types.AddTenantRsp, error) {
	// todo: add your logic here and delete this line

	fmt.Println(req)

	rsp,err :=l.svcCtx.TenantRpc.AddTenant(l.ctx, &tenant.AddTenantReq{
		Name: req.Name,
		Addr: req.Addr,
	})
	if err !=nil{
		return nil,err
	}
	return &types.AddTenantRsp{Id: rsp.Id},nil
}
