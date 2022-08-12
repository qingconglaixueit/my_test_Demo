package logic

import (
	"context"

	"my_test_demo/my_book_sys/book/api/internal/svc"
	"my_test_demo/my_book_sys/book/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (*types.SearchReply, error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("stuId: %v", l.ctx.Value("stuId"))

	bookInfo, err := l.svcCtx.BookHttpModel.FindOneByName(l.ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &types.SearchReply{
		Name:  bookInfo.Name,
		Count: int(bookInfo.Count),
	}, nil
}
