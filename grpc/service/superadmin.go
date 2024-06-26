package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/superadmin_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SuperadminService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewSuperadminService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SuperadminService {
	return &SuperadminService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *SuperadminService) Create(ctx context.Context, req *superadmin_service.CreateSuperadmin) (*superadmin_service.GetSuperadmin, error) {

	f.log.Info("---CreateSuperadmin--->>>", logger.Any("req", req))

	resp, err := f.strg.Superadmin().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateSuperadmin--->>>", logger.Error(err))
		return &superadmin_service.GetSuperadmin{}, err
	}

	return resp, nil
}
func (f *SuperadminService) Update(ctx context.Context, req *superadmin_service.UpdateSuperadmin) (*superadmin_service.GetSuperadmin, error) {

	f.log.Info("---UpdateSuperadmin--->>>", logger.Any("req", req))

	resp, err := f.strg.Superadmin().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateSuperadmin--->>>", logger.Error(err))
		return &superadmin_service.GetSuperadmin{}, err
	}

	return resp, nil
}

func (f *SuperadminService) GetByID(ctx context.Context, id *superadmin_service.SuperadminPrimaryKey) (*superadmin_service.GetSuperadmin, error) {
	f.log.Info("---GetSuperadmin--->>>", logger.Any("req", id))

	resp, err := f.strg.Superadmin().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetSuperadmin--->>>", logger.Error(err))
		return &superadmin_service.GetSuperadmin{}, err
	}

	return resp, nil
}

func (f *SuperadminService) Delete(ctx context.Context, req *superadmin_service.SuperadminPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteSuperadmin--->>>", logger.Any("req", req))

	_, err := f.strg.Superadmin().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteSuperadmin--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
