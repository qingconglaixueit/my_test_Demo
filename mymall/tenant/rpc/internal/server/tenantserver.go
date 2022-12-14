// Code generated by goctl. DO NOT EDIT!
// Source: tenant.proto

package server

import (
	"context"

	"my_test_demo/mymall/tenant/rpc/internal/logic"
	"my_test_demo/mymall/tenant/rpc/internal/svc"
	"my_test_demo/mymall/tenant/rpc/protoc/tenant"
)

type TenantServer struct {
	svcCtx *svc.ServiceContext
	tenant.UnimplementedTenantServer
}

func NewTenantServer(svcCtx *svc.ServiceContext) *TenantServer {
	return &TenantServer{
		svcCtx: svcCtx,
	}
}

func (s *TenantServer) GetTenant(ctx context.Context, in *tenant.TidReq) (*tenant.TenantRsp, error) {
	l := logic.NewGetTenantLogic(ctx, s.svcCtx)
	return l.GetTenant(in)
}

func (s *TenantServer) AddTenant(ctx context.Context, in *tenant.AddTenantReq) (*tenant.AddTenantRsp, error) {
	l := logic.NewAddTenantLogic(ctx, s.svcCtx)
	return l.AddTenant(in)
}
