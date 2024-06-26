package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/manager_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ManagerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewManagerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ManagerService {
	return &ManagerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *ManagerService) Create(ctx context.Context, req *manager_service.CreateManager) (*manager_service.GetManager, error) {

	f.log.Info("---CreateManager--->>>", logger.Any("req", req))

	resp, err := f.strg.Manager().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateManager--->>>", logger.Error(err))
		return &manager_service.GetManager{}, err
	}

	return resp, nil
}
func (f *ManagerService) Update(ctx context.Context, req *manager_service.UpdateManager) (*manager_service.GetManager, error) {

	f.log.Info("---UpdateManager--->>>", logger.Any("req", req))

	resp, err := f.strg.Manager().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateManager--->>>", logger.Error(err))
		return &manager_service.GetManager{}, err
	}

	return resp, nil
}

func (f *ManagerService) GetList(ctx context.Context, req *manager_service.GetListManagerRequest) (*manager_service.GetListManagerResponse, error) {
	f.log.Info("---GetListManager--->>>", logger.Any("req", req))

	resp, err := f.strg.Manager().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListManager--->>>", logger.Error(err))
		return &manager_service.GetListManagerResponse{}, err
	}

	return resp, nil
}

func (f *ManagerService) GetByID(ctx context.Context, id *manager_service.ManagerPrimaryKey) (*manager_service.GetManager, error) {
	f.log.Info("---GetManager--->>>", logger.Any("req", id))

	resp, err := f.strg.Manager().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetManager--->>>", logger.Error(err))
		return &manager_service.GetManager{}, err
	}

	return resp, nil
}

func (f *ManagerService) Delete(ctx context.Context, req *manager_service.ManagerPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteManager--->>>", logger.Any("req", req))

	_, err := f.strg.Manager().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteManager--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
