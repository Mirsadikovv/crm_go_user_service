package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/group_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GroupService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewGroupService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *GroupService {
	return &GroupService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *GroupService) Create(ctx context.Context, req *group_service.CreateGroup) (*group_service.GetGroup, error) {

	f.log.Info("---CreateGroup--->>>", logger.Any("req", req))

	resp, err := f.strg.Group().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateGroup--->>>", logger.Error(err))
		return &group_service.GetGroup{}, err
	}

	return resp, nil
}
func (f *GroupService) Update(ctx context.Context, req *group_service.UpdateGroup) (*group_service.GetGroup, error) {

	f.log.Info("---UpdateGroup--->>>", logger.Any("req", req))

	resp, err := f.strg.Group().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateGroup--->>>", logger.Error(err))
		return &group_service.GetGroup{}, err
	}

	return resp, nil
}

func (f *GroupService) GetList(ctx context.Context, req *group_service.GetListGroupRequest) (*group_service.GetListGroupResponse, error) {
	f.log.Info("---GetListGroup--->>>", logger.Any("req", req))

	resp, err := f.strg.Group().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListGroup--->>>", logger.Error(err))
		return &group_service.GetListGroupResponse{}, err
	}

	return resp, nil
}

func (f *GroupService) GetByID(ctx context.Context, id *group_service.GroupPrimaryKey) (*group_service.GetGroup, error) {
	f.log.Info("---GetGroup--->>>", logger.Any("req", id))

	resp, err := f.strg.Group().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetGroup--->>>", logger.Error(err))
		return &group_service.GetGroup{}, err
	}

	return resp, nil
}

func (f *GroupService) Delete(ctx context.Context, req *group_service.GroupPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteGroup--->>>", logger.Any("req", req))

	_, err := f.strg.Group().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteGroup--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
