package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/student_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StudentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewStudentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *StudentService {
	return &StudentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *StudentService) Create(ctx context.Context, req *student_service.CreateStudent) (*student_service.GetStudent, error) {

	f.log.Info("---CreateStudent--->>>", logger.Any("req", req))

	resp, err := f.strg.Student().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateStudent--->>>", logger.Error(err))
		return &student_service.GetStudent{}, err
	}

	return resp, nil
}
func (f *StudentService) Update(ctx context.Context, req *student_service.UpdateStudent) (*student_service.GetStudent, error) {

	f.log.Info("---UpdateStudent--->>>", logger.Any("req", req))

	resp, err := f.strg.Student().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateStudent--->>>", logger.Error(err))
		return &student_service.GetStudent{}, err
	}

	return resp, nil
}

func (f *StudentService) GetList(ctx context.Context, req *student_service.GetListStudentRequest) (*student_service.GetListStudentResponse, error) {
	f.log.Info("---GetListStudent--->>>", logger.Any("req", req))

	resp, err := f.strg.Student().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListStudent--->>>", logger.Error(err))
		return &student_service.GetListStudentResponse{}, err
	}

	return resp, nil
}

func (f *StudentService) GetByID(ctx context.Context, id *student_service.StudentPrimaryKey) (*student_service.GetStudent, error) {
	f.log.Info("---GetStudent--->>>", logger.Any("req", id))

	resp, err := f.strg.Student().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetStudent--->>>", logger.Error(err))
		return &student_service.GetStudent{}, err
	}

	return resp, nil
}

func (f *StudentService) Delete(ctx context.Context, req *student_service.StudentPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteStudent--->>>", logger.Any("req", req))

	_, err := f.strg.Student().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteStudent--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
