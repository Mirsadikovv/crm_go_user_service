package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/branch_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BranchService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BranchService {
	return &BranchService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *BranchService) Create(ctx context.Context, req *branch_service.CreateBranch) (*branch_service.GetBranch, error) {

	f.log.Info("---CreateBranch--->>>", logger.Any("req", req))

	resp, err := f.strg.Branch().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateBranch--->>>", logger.Error(err))
		return &branch_service.GetBranch{}, err
	}

	return resp, nil
}
func (f *BranchService) Update(ctx context.Context, req *branch_service.UpdateBranch) (*branch_service.GetBranch, error) {

	f.log.Info("---UpdateBranch--->>>", logger.Any("req", req))

	resp, err := f.strg.Branch().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateBranch--->>>", logger.Error(err))
		return &branch_service.GetBranch{}, err
	}

	return resp, nil
}

func (f *BranchService) GetList(ctx context.Context, req *branch_service.GetListBranchRequest) (*branch_service.GetListBranchResponse, error) {
	f.log.Info("---GetListBranch--->>>", logger.Any("req", req))

	resp, err := f.strg.Branch().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListBranch--->>>", logger.Error(err))
		return &branch_service.GetListBranchResponse{}, err
	}

	return resp, nil
}

func (f *BranchService) GetByID(ctx context.Context, id *branch_service.BranchPrimaryKey) (*branch_service.GetBranch, error) {
	f.log.Info("---GetBranch--->>>", logger.Any("req", id))

	resp, err := f.strg.Branch().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetBranch--->>>", logger.Error(err))
		return &branch_service.GetBranch{}, err
	}

	return resp, nil
}

func (f *BranchService) Delete(ctx context.Context, req *branch_service.BranchPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteBranch--->>>", logger.Any("req", req))

	_, err := f.strg.Branch().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteBranch--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
