package logic

import (
	"context"

	"my_test_demo/mymall/tenant/rpc/internal/svc"
	"my_test_demo/mymall/tenant/rpc/protoc/tenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTenantLogic {
	return &GetTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTenantLogic) GetTenant(in *tenant.TidReq) (*tenant.TenantRsp, error) {
	// todo: add your logic here and delete this line

	res,err :=l.svcCtx.Model.FindOneByName(l.ctx,in.Name)
	if err !=nil{
		return nil,err
	}

	return &tenant.TenantRsp{
		Id: res.Id,
		Name: res.Name,
		Addr: res.Addr,
	},nil
}
