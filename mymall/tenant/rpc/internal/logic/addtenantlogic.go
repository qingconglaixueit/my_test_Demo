package logic

import (
	"context"
	"crypto/md5"
	"fmt"

	"my_test_demo/mymall/tenant/rpc/internal/svc"
	"my_test_demo/mymall/tenant/rpc/protoc/tenant"
	"my_test_demo/mymall/tenant/rpc/protoc/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTenantLogic {
	return &AddTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTenantLogic) AddTenant(in *tenant.AddTenantReq) (*tenant.AddTenantRsp, error) {
	// todo: add your logic here and delete this line
	h := md5.New()
	h.Write([]byte(in.Name))
	id := fmt.Sprintf("%x", h.Sum(nil))[:6]

	_, err := l.svcCtx.Model.Insert(l.ctx, &model.TenantInfo{
		Id:   id,
		Name: in.Name,
		Addr: in.Addr})
	if err != nil {
		return nil, err
	}

	l.Logger.Info(" ---------------------- id = ", id)


	return &tenant.AddTenantRsp{Id: id}, nil
}
