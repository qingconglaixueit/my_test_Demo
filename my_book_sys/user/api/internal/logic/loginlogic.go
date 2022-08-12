package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"

	"my_test_demo/my_book_sys/user/api/internal/svc"
	"my_test_demo/my_book_sys/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// 校验基本入参
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("login params invalid")
	}
	// 校验 DB 用户密码
	userInfo, err := l.svcCtx.UserHttpModel.FindOneByName(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if userInfo.Password != req.Password {
		return nil, errors.New("username or password error")
	}
	// 生成 jwt 响应
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.StuId)
	if err != nil {
		return nil, err
	}

	return &types.LoginReply{
		Id:           userInfo.StuId,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
	}, nil
}

func (l *LoginLogic) getJwtToken(secret string, nowDate int64, accessExpire int64, StuId string) (jwtToken string, err error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = nowDate + accessExpire
	claims["iat"] = nowDate
	claims["stuId"] = StuId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
