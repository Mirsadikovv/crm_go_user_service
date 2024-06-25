package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/teacher_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TeacherService {
	return &TeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *TeacherService) Create(ctx context.Context, req *teacher_service.CreateTeacher) (*teacher_service.GetTeacher, error) {

	f.log.Info("---CreateTeacher--->>>", logger.Any("req", req))

	resp, err := f.strg.Teacher().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateTeacher--->>>", logger.Error(err))
		return &teacher_service.GetTeacher{}, err
	}

	return resp, nil
}
func (f *TeacherService) Update(ctx context.Context, req *teacher_service.UpdateTeacher) (*teacher_service.GetTeacher, error) {

	f.log.Info("---UpdateTeacher--->>>", logger.Any("req", req))

	resp, err := f.strg.Teacher().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateTeacher--->>>", logger.Error(err))
		return &teacher_service.GetTeacher{}, err
	}

	return resp, nil
}

func (f *TeacherService) GetList(ctx context.Context, req *teacher_service.GetListTeacherRequest) (*teacher_service.GetListTeacherResponse, error) {
	f.log.Info("---GetListTeacher--->>>", logger.Any("req", req))

	resp, err := f.strg.Teacher().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListTeacher--->>>", logger.Error(err))
		return &teacher_service.GetListTeacherResponse{}, err
	}

	return resp, nil
}

func (f *TeacherService) GetByID(ctx context.Context, id *teacher_service.TeacherPrimaryKey) (*teacher_service.GetTeacher, error) {
	f.log.Info("---GetTeacher--->>>", logger.Any("req", id))

	resp, err := f.strg.Teacher().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetTeacher--->>>", logger.Error(err))
		return &teacher_service.GetTeacher{}, err
	}

	return resp, nil
}

func (f *TeacherService) Delete(ctx context.Context, req *teacher_service.TeacherPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteTeacher--->>>", logger.Any("req", req))

	_, err := f.strg.Teacher().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteTeacher--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
