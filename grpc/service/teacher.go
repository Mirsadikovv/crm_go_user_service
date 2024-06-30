package service

import (
	"context"
	"errors"
	"fmt"
	"go_user_service/config"
	"go_user_service/genproto/teacher_service"
	"go_user_service/pkg"
	"go_user_service/pkg/hash"
	"go_user_service/pkg/jwt"
	"go_user_service/pkg/smtp"
	"time"

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
	redis    storage.IRedisStorage
}

func NewTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI, redis storage.IRedisStorage) *TeacherService {
	return &TeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
		redis:    redis,
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

func (a *TeacherService) Login(ctx context.Context, loginRequest *teacher_service.TeacherLoginRequest) (*teacher_service.TeacherLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.UserLogin)
	teacher, err := a.strg.Teacher().GetByLogin(ctx, loginRequest.UserLogin)
	if err != nil {
		a.log.Error("error while getting teacher credentials by login", logger.Error(err))
		return &teacher_service.TeacherLoginResponse{}, err
	}

	if err = hash.CompareHashAndPassword(teacher.UserPassword, loginRequest.UserPassword); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return &teacher_service.TeacherLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = teacher.Id
	m["user_role"] = config.TEACHER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for teacher login", logger.Error(err))
		return &teacher_service.TeacherLoginResponse{}, err
	}

	return &teacher_service.TeacherLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *TeacherService) Register(ctx context.Context, loginRequest *teacher_service.TeacherRegisterRequest) (*emptypb.Empty, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Mail)

	otpCode := pkg.GenerateOTP()

	msg := fmt.Sprintf("Your otp code is: %v, for registering RENT_CAR. Don't give it to anyone", otpCode)

	err := a.redis.SetX(ctx, loginRequest.Mail, otpCode, time.Minute*2)
	if err != nil {
		a.log.Error("error while setting otpCode to redis teacher register", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	err = smtp.SendMail(loginRequest.Mail, msg)
	if err != nil {
		a.log.Error("error while sending otp code to teacher register", logger.Error(err))
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (a *TeacherService) RegisterConfirm(ctx context.Context, req *teacher_service.TeacherRegisterConfRequest) (*teacher_service.TeacherLoginResponse, error) {
	resp := &teacher_service.TeacherLoginResponse{}

	otp, err := a.redis.Get(ctx, req.Mail)
	if err != nil {
		a.log.Error("error while getting otp code for teacher register confirm", logger.Error(err))
		return resp, err
	}
	if req.Otp != otp {
		a.log.Error("incorrect otp code for teacher register confirm", logger.Error(err))
		return resp, errors.New("incorrect otp code")
	}
	req.Teacher[0].Email = req.Mail

	id, err := a.strg.Teacher().Create(ctx, req.Teacher[0])
	if err != nil {
		a.log.Error("error while creating teacher", logger.Error(err))
		return resp, err
	}
	var m = make(map[interface{}]interface{})

	m["user_id"] = id
	m["user_role"] = config.TEACHER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for teacher register confirm", logger.Error(err))
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}

func (f *TeacherService) ChangePassword(ctx context.Context, pass *teacher_service.TeacherChangePassword) (*teacher_service.TeacherChangePasswordResp, error) {
	f.log.Info("---ChangePassword--->>>", logger.Any("req", pass))

	resp, err := f.strg.Teacher().ChangePassword(ctx, pass)
	if err != nil {
		f.log.Error("---ChangePassword--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
