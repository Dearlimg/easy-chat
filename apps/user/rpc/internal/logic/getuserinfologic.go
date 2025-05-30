package logic

import (
	"context"
	"easy-chat/apps/user/models"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNotFound = errors.New("这个用户没有")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	userEntiy, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	var resp user.UserEntity
	copier.Copy(&resp, &userEntiy)

	fmt.Println()
	return &user.GetUserInfoResp{
		User: &resp,
	}, nil
}
