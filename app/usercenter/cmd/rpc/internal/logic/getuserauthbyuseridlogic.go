package logic

import (
	"cloud-disk/app/usercenter/cmd/rpc/usercenter"
	"cloud-disk/app/usercenter/model"
	"cloud-disk/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"cloud-disk/app/usercenter/cmd/rpc/internal/svc"
	"cloud-disk/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *pb.GetUserAuthByUserIdReq) (*pb.GetUserAuthByUserIdResp, error) {

	userAuth, err := l.svcCtx.UserAuthModel.FindOneByUserIdAuthType(l.ctx, in.UserId, in.AuthType)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth  failed"), "err : %v , in : %+v", err, in)
	}

	var respUserAuth usercenter.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)
	return &pb.GetUserAuthByUserIdResp{
		UserAuth: &respUserAuth,
	}, nil
}
