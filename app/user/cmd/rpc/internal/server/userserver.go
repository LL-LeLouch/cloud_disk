// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"cloud-disk/app/user/cmd/rpc/internal/logic"
	"cloud-disk/app/user/cmd/rpc/internal/svc"
	"cloud-disk/app/user/cmd/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// rpc login(LoginReq) returns(LoginResp);
func (s *UserServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserServer) GetUserAuthByAuthKey(ctx context.Context, in *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthByAuthKeyResp, error) {
	l := logic.NewGetUserAuthByAuthKeyLogic(ctx, s.svcCtx)
	return l.GetUserAuthByAuthKey(in)
}

func (s *UserServer) GetUserAuthByIdentity(ctx context.Context, in *pb.GetUserAuthByIdentityReq) (*pb.GetUserAuthByIdentityResp, error) {
	l := logic.NewGetUserAuthByIdentityLogic(ctx, s.svcCtx)
	return l.GetUserAuthByIdentity(in)
}

func (s *UserServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}
