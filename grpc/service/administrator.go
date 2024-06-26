package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/administrator_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AdministratorService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewAdministratorService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *AdministratorService {
	return &AdministratorService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *AdministratorService) Create(ctx context.Context, req *administrator_service.CreateAdministrator) (*administrator_service.GetAdministrator, error) {

	f.log.Info("---CreateAdministrator--->>>", logger.Any("req", req))

	resp, err := f.strg.Administrator().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateAdministrator--->>>", logger.Error(err))
		return &administrator_service.GetAdministrator{}, err
	}

	return resp, nil
}
func (f *AdministratorService) Update(ctx context.Context, req *administrator_service.UpdateAdministrator) (*administrator_service.GetAdministrator, error) {

	f.log.Info("---UpdateAdministrator--->>>", logger.Any("req", req))

	resp, err := f.strg.Administrator().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateAdministrator--->>>", logger.Error(err))
		return &administrator_service.GetAdministrator{}, err
	}

	return resp, nil
}

func (f *AdministratorService) GetList(ctx context.Context, req *administrator_service.GetListAdministratorRequest) (*administrator_service.GetListAdministratorResponse, error) {
	f.log.Info("---GetListAdministrator--->>>", logger.Any("req", req))

	resp, err := f.strg.Administrator().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListAdministrator--->>>", logger.Error(err))
		return &administrator_service.GetListAdministratorResponse{}, err
	}

	return resp, nil
}

func (f *AdministratorService) GetByID(ctx context.Context, id *administrator_service.AdministratorPrimaryKey) (*administrator_service.GetAdministrator, error) {
	f.log.Info("---GetAdministrator--->>>", logger.Any("req", id))

	resp, err := f.strg.Administrator().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetAdministrator--->>>", logger.Error(err))
		return &administrator_service.GetAdministrator{}, err
	}

	return resp, nil
}

func (f *AdministratorService) Delete(ctx context.Context, req *administrator_service.AdministratorPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteAdministrator--->>>", logger.Any("req", req))

	_, err := f.strg.Administrator().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteAdministrator--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
