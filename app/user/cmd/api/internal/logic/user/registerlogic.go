package logic

import (
	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/app/user/model"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	//验证
	if res, _ := l.svcCtx.RedisClient.Get(req.Email); req.Captcha != res {
		return nil, errors.Wrapf(nil, "验证码错误")
	}

	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Email:    req.Email,
		Password: req.Password,
		AuthKey:  req.Email,
		AuthType: model.UserAuthTypeSystem,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "UserRpc.Register err req: %+v", req)
	}
	var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)
	return &resp, nil
}
