package service

import (
	"context"
	"errors"
	"fmt"
	"go_user_service/config"
	"go_user_service/genproto/superadmin_service"
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

type SuperadminService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	redis    storage.IRedisStorage
}

func NewSuperadminService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI, redis storage.IRedisStorage) *SuperadminService {
	return &SuperadminService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
		redis:    redis,
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

func (a *SuperadminService) Login(ctx context.Context, loginRequest *superadmin_service.SuperadminLoginRequest) (*superadmin_service.SuperadminLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.UserLogin)
	superadmin, err := a.strg.Superadmin().GetByLogin(ctx, loginRequest.UserLogin)
	if err != nil {
		a.log.Error("error while getting superadmin credentials by login", logger.Error(err))
		return &superadmin_service.SuperadminLoginResponse{}, err
	}

	if err = hash.CompareHashAndPassword(superadmin.UserPassword, loginRequest.UserPassword); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return &superadmin_service.SuperadminLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = superadmin.Id
	m["user_role"] = config.SUPERADMIN_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for superadmin login", logger.Error(err))
		return &superadmin_service.SuperadminLoginResponse{}, err
	}

	return &superadmin_service.SuperadminLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *SuperadminService) Register(ctx context.Context, loginRequest *superadmin_service.SuperadminRegisterRequest) (*emptypb.Empty, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Mail)

	otpCode := pkg.GenerateOTP()

	msg := fmt.Sprintf("Your otp code is: %v, for registering CRM system. Don't give it to anyone", otpCode)

	err := a.redis.SetX(ctx, loginRequest.Mail, otpCode, time.Minute*2)
	if err != nil {
		a.log.Error("error while setting otpCode to redis superadmin register", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	err = smtp.SendMail(loginRequest.Mail, msg)
	if err != nil {
		a.log.Error("error while sending otp code to superadmin register", logger.Error(err))
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (a *SuperadminService) RegisterConfirm(ctx context.Context, req *superadmin_service.SuperadminRegisterConfRequest) (*superadmin_service.SuperadminLoginResponse, error) {
	resp := &superadmin_service.SuperadminLoginResponse{}

	otp, err := a.redis.Get(ctx, req.Mail)
	if err != nil {
		a.log.Error("error while getting otp code for superadmin register confirm", logger.Error(err))
		return resp, err
	}
	if req.Otp != otp {
		a.log.Error("incorrect otp code for superadmin register confirm", logger.Error(err))
		return resp, errors.New("incorrect otp code")
	}
	req.Superadmin[0].Email = req.Mail

	id, err := a.strg.Superadmin().Create(ctx, req.Superadmin[0])
	if err != nil {
		a.log.Error("error while creating superadmin", logger.Error(err))
		return resp, err
	}
	var m = make(map[interface{}]interface{})

	m["user_id"] = id
	m["user_role"] = config.SUPERADMIN_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for superadmin register confirm", logger.Error(err))
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}

func (f *SuperadminService) ChangePassword(ctx context.Context, pass *superadmin_service.SuperadminChangePassword) (*superadmin_service.SuperadminChangePasswordResp, error) {
	f.log.Info("---ChangePassword--->>>", logger.Any("req", pass))

	resp, err := f.strg.Superadmin().ChangePassword(ctx, pass)
	if err != nil {
		f.log.Error("---ChangePassword--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
