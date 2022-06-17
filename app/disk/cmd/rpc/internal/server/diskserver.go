// Code generated by goctl. DO NOT EDIT!
// Source: disk.proto

package server

import (
	"context"

	"cloud-disk/app/disk/cmd/rpc/internal/logic"
	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"
)

type DiskServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedDiskServer
}

func NewDiskServer(svcCtx *svc.ServiceContext) *DiskServer {
	return &DiskServer{
		svcCtx: svcCtx,
	}
}

// store 详情
func (s *DiskServer) DetailStore(ctx context.Context, in *pb.StoreDetailReq) (*pb.StoreDetailResp, error) {
	l := logic.NewDetailStoreLogic(ctx, s.svcCtx)
	return l.DetailStore(in)
}

// 修改store大小
func (s *DiskServer) UpdateStore(ctx context.Context, in *pb.UpdateStoreReq) (*pb.UpdateStoreResp, error) {
	l := logic.NewUpdateStoreLogic(ctx, s.svcCtx)
	return l.UpdateStore(in)
}

// 获取路径下的文件和文件夹
func (s *DiskServer) ListFolders(ctx context.Context, in *pb.ListFolderReq) (*pb.ListFolderResp, error) {
	l := logic.NewListFoldersLogic(ctx, s.svcCtx)
	return l.ListFolders(in)
}

// 更新file信息
func (s *DiskServer) UpdateFile(ctx context.Context, in *pb.UpdateFileReq) (*pb.UpdateFileResp, error) {
	l := logic.NewUpdateFileLogic(ctx, s.svcCtx)
	return l.UpdateFile(in)
}

// 更新folder信息
func (s *DiskServer) UpdateFolder(ctx context.Context, in *pb.UpdateFolderReq) (*pb.UpdateFolderResp, error) {
	l := logic.NewUpdateFolderLogic(ctx, s.svcCtx)
	return l.UpdateFolder(in)
}
