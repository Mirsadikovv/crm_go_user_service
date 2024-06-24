package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/catalog_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CategoryService {
	return &CategoryService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *CategoryService) Create(ctx context.Context, req *catalog_service.CreateCategory) (*catalog_service.GetCategory, error) {

	f.log.Info("---CreateCategory--->>>", logger.Any("req", req))

	resp, err := f.strg.Category().CreateCategory(ctx, req)
	if err != nil {
		f.log.Error("---CreateCategory--->>>", logger.Error(err))
		return &catalog_service.GetCategory{}, err
	}

	return resp, nil
}
func (f *CategoryService) Update(ctx context.Context, req *catalog_service.UpdateCategory) (*catalog_service.GetCategory, error) {

	f.log.Info("---UpdateCategory--->>>", logger.Any("req", req))

	resp, err := f.strg.Category().UpdateCategory(ctx, req)
	if err != nil {
		f.log.Error("---UpdateCategory--->>>", logger.Error(err))
		return &catalog_service.GetCategory{}, err
	}

	return resp, nil
}

func (f *CategoryService) GetList(ctx context.Context, req *catalog_service.GetListCategoryRequest) (*catalog_service.GetListCategoryResponse, error) {
	f.log.Info("---GetListCategory--->>>", logger.Any("req", req))

	resp, err := f.strg.Category().GetListCategory(ctx, req)
	if err != nil {
		f.log.Error("---GetListCategory--->>>", logger.Error(err))
		return &catalog_service.GetListCategoryResponse{}, err
	}

	return resp, nil
}

func (f *CategoryService) GetByID(ctx context.Context, id *catalog_service.CategoryPrimaryKey) (*catalog_service.GetCategory, error) {
	f.log.Info("---GetCategory--->>>", logger.Any("req", id))

	resp, err := f.strg.Category().GetCategoryById(ctx, id)
	if err != nil {
		f.log.Error("---GetCategory--->>>", logger.Error(err))
		return &catalog_service.GetCategory{}, err
	}

	return resp, nil
}

func (f *CategoryService) Delete(ctx context.Context, req *catalog_service.CategoryPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteCategory--->>>", logger.Any("req", req))

	_, err := f.strg.Category().DeleteCategory(ctx, req)
	if err != nil {
		f.log.Error("---DeleteCategory--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
