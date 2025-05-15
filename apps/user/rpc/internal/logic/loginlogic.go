package logic

import (
	"context"
	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/ctxdata"
	"easy-chat/pkg/encrypt"
	"easy-chat/pkg/xerr"
	"githu
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneNotRegistered = xerr.New(xerr.SERVER_COMMON_ERROR, "手机号未注册")
	ErrUserPwdErr         = xerr.New(xerr.SERVER_COMMON_ERROR, "密码错误")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errors.WithStack(ErrPhoneNotRegistered)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user by phone failed,err: %v,req %v", err, in.Phone)
	}

	if !encrypt.ValidatePasswordHash(in.Password, userEntity.Password.String) {
		return nil, errors.WithStack(ErrUserPwdErr)
	}

	// 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire,
		userEntity.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "get jwt token failed,err: %v,req %v", err, in.Phone)
	}

	return &user.LoginResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
