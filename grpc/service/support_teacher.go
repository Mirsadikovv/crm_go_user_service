package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/support_teacher_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SupportTeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewSupportTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SupportTeacherService {
	return &SupportTeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *SupportTeacherService) Create(ctx context.Context, req *support_teacher_service.CreateSupportTeacher) (*support_teacher_service.GetSupportTeacher, error) {

	f.log.Info("---CreateSupportTeacher--->>>", logger.Any("req", req))

	resp, err := f.strg.SupportTeacher().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateSupportTeacher--->>>", logger.Error(err))
		return &support_teacher_service.GetSupportTeacher{}, err
	}

	return resp, nil
}
func (f *SupportTeacherService) Update(ctx context.Context, req *support_teacher_service.UpdateSupportTeacher) (*support_teacher_service.GetSupportTeacher, error) {

	f.log.Info("---UpdateSupportTeacher--->>>", logger.Any("req", req))

	resp, err := f.strg.SupportTeacher().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateSupportTeacher--->>>", logger.Error(err))
		return &support_teacher_service.GetSupportTeacher{}, err
	}

	return resp, nil
}

func (f *SupportTeacherService) GetList(ctx context.Context, req *support_teacher_service.GetListSupportTeacherRequest) (*support_teacher_service.GetListSupportTeacherResponse, error) {
	f.log.Info("---GetListSupportTeacher--->>>", logger.Any("req", req))

	resp, err := f.strg.SupportTeacher().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListSupportTeacher--->>>", logger.Error(err))
		return &support_teacher_service.GetListSupportTeacherResponse{}, err
	}

	return resp, nil
}

func (f *SupportTeacherService) GetByID(ctx context.Context, id *support_teacher_service.SupportTeacherPrimaryKey) (*support_teacher_service.GetSupportTeacher, error) {
	f.log.Info("---GetSupportTeacher--->>>", logger.Any("req", id))

	resp, err := f.strg.SupportTeacher().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetSupportTeacher--->>>", logger.Error(err))
		return &support_teacher_service.GetSupportTeacher{}, err
	}

	return resp, nil
}

func (f *SupportTeacherService) Delete(ctx context.Context, req *support_teacher_service.SupportTeacherPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteSupportTeacher--->>>", logger.Any("req", req))

	_, err := f.strg.SupportTeacher().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteSupportTeacher--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
