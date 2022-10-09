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

func (s *DiskServer) FileUploadPrepare(ctx context.Context, in *pb.FileUploadPrepareRep) (*pb.FileUploadPrepareResp, error) {
	l := logic.NewFileUploadPrepareLogic(ctx, s.svcCtx)
	return l.FileUploadPrepare(in)
}

func (s *DiskServer) UpdateFile(ctx context.Context, in *pb.UpdateFileReq) (*pb.UpdateFileResp, error) {
	l := logic.NewUpdateFileLogic(ctx, s.svcCtx)
	return l.UpdateFile(in)
}

//func (s *DiskServer) Statistics(ctx context.Context, in *pb.StatisticsReq) (*pb.StatisticsResp, error) {
//	l := logic.NewStatisticsLogic(ctx, s.svcCtx)
//	return l.Statistics(in)
//}

func (s *DiskServer) ListFile(ctx context.Context, in *pb.ListFileReq) (*pb.ListFileResp, error) {
	l := logic.NewListFileLogic(ctx, s.svcCtx)
	return l.ListFile(in)
}
