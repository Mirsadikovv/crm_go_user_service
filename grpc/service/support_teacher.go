package service

import (
	"context"
	"errors"
	"fmt"
	"go_user_service/config"
	"go_user_service/genproto/support_teacher_service"
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

type SupportTeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	redis    storage.IRedisStorage
}

func NewSupportTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI, redis storage.IRedisStorage) *SupportTeacherService {
	return &SupportTeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
		redis:    redis,
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

func (a *SupportTeacherService) Login(ctx context.Context, loginRequest *support_teacher_service.SupportTeacherLoginRequest) (*support_teacher_service.SupportTeacherLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.UserLogin)
	support_teacher, err := a.strg.SupportTeacher().GetByLogin(ctx, loginRequest.UserLogin)
	if err != nil {
		a.log.Error("error while getting support_teacher credentials by login", logger.Error(err))
		return &support_teacher_service.SupportTeacherLoginResponse{}, err
	}

	if err = hash.CompareHashAndPassword(support_teacher.UserPassword, loginRequest.UserPassword); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return &support_teacher_service.SupportTeacherLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = support_teacher.Id
	m["user_role"] = config.SUPPORT_TEACHER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for support_teacher login", logger.Error(err))
		return &support_teacher_service.SupportTeacherLoginResponse{}, err
	}

	return &support_teacher_service.SupportTeacherLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *SupportTeacherService) Register(ctx context.Context, loginRequest *support_teacher_service.SupportTeacherRegisterRequest) (*emptypb.Empty, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Mail)

	otpCode := pkg.GenerateOTP()

	msg := fmt.Sprintf("Your otp code is: %v, for registering CRM system. Don't give it to anyone", otpCode)

	err := a.redis.SetX(ctx, loginRequest.Mail, otpCode, time.Minute*2)
	if err != nil {
		a.log.Error("error while setting otpCode to redis support_teacher register", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	err = smtp.SendMail(loginRequest.Mail, msg)
	if err != nil {
		a.log.Error("error while sending otp code to support_teacher register", logger.Error(err))
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (a *SupportTeacherService) RegisterConfirm(ctx context.Context, req *support_teacher_service.SupportTeacherRegisterConfRequest) (*support_teacher_service.SupportTeacherLoginResponse, error) {
	resp := &support_teacher_service.SupportTeacherLoginResponse{}

	otp, err := a.redis.Get(ctx, req.Mail)
	if err != nil {
		a.log.Error("error while getting otp code for support_teacher register confirm", logger.Error(err))
		return resp, err
	}
	if req.Otp != otp {
		a.log.Error("incorrect otp code for support_teacher register confirm", logger.Error(err))
		return resp, errors.New("incorrect otp code")
	}
	req.SupportTeacher[0].Email = req.Mail

	id, err := a.strg.SupportTeacher().Create(ctx, req.SupportTeacher[0])
	if err != nil {
		a.log.Error("error while creating support_teacher", logger.Error(err))
		return resp, err
	}
	var m = make(map[interface{}]interface{})

	m["user_id"] = id
	m["user_role"] = config.SUPPORT_TEACHER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for support_teacher register confirm", logger.Error(err))
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}

func (f *SupportTeacherService) ChangePassword(ctx context.Context, pass *support_teacher_service.SupportTeacherChangePassword) (*support_teacher_service.SupportTeacherChangePasswordResp, error) {
	f.log.Info("---ChangePassword--->>>", logger.Any("req", pass))

	resp, err := f.strg.SupportTeacher().ChangePassword(ctx, pass)
	if err != nil {
		f.log.Error("---ChangePassword--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
