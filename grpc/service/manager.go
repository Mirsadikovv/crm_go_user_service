package service

import (
	"context"
	"errors"
	"fmt"
	"go_user_service/config"
	"go_user_service/genproto/manager_service"
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

type ManagerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	redis    storage.IRedisStorage
}

func NewManagerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI, redis storage.IRedisStorage) *ManagerService {
	return &ManagerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
		redis:    redis,
	}
}

func (f *ManagerService) Create(ctx context.Context, req *manager_service.CreateManager) (*manager_service.GetManager, error) {

	f.log.Info("---CreateManager--->>>", logger.Any("req", req))

	resp, err := f.strg.Manager().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateManager--->>>", logger.Error(err))
		return &manager_service.GetManager{}, err
	}

	return resp, nil
}
func (f *ManagerService) Update(ctx context.Context, req *manager_service.UpdateManager) (*manager_service.GetManager, error) {

	f.log.Info("---UpdateManager--->>>", logger.Any("req", req))

	resp, err := f.strg.Manager().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateManager--->>>", logger.Error(err))
		return &manager_service.GetManager{}, err
	}

	return resp, nil
}

func (f *ManagerService) GetList(ctx context.Context, req *manager_service.GetListManagerRequest) (*manager_service.GetListManagerResponse, error) {
	f.log.Info("---GetListManager--->>>", logger.Any("req", req))

	resp, err := f.strg.Manager().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListManager--->>>", logger.Error(err))
		return &manager_service.GetListManagerResponse{}, err
	}

	return resp, nil
}

func (f *ManagerService) GetByID(ctx context.Context, id *manager_service.ManagerPrimaryKey) (*manager_service.GetManager, error) {
	f.log.Info("---GetManager--->>>", logger.Any("req", id))

	resp, err := f.strg.Manager().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetManager--->>>", logger.Error(err))
		return &manager_service.GetManager{}, err
	}

	return resp, nil
}

func (f *ManagerService) Delete(ctx context.Context, req *manager_service.ManagerPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteManager--->>>", logger.Any("req", req))

	_, err := f.strg.Manager().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteManager--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (a *ManagerService) Login(ctx context.Context, loginRequest *manager_service.ManagerLoginRequest) (*manager_service.ManagerLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.UserLogin)
	manager, err := a.strg.Manager().GetByLogin(ctx, loginRequest.UserLogin)
	if err != nil {
		a.log.Error("error while getting manager credentials by login", logger.Error(err))
		return &manager_service.ManagerLoginResponse{}, err
	}

	if err = hash.CompareHashAndPassword(manager.UserPassword, loginRequest.UserPassword); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return &manager_service.ManagerLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = manager.Id
	m["user_role"] = config.MANAGER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for manager login", logger.Error(err))
		return &manager_service.ManagerLoginResponse{}, err
	}

	return &manager_service.ManagerLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *ManagerService) Register(ctx context.Context, loginRequest *manager_service.ManagerRegisterRequest) (*emptypb.Empty, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Mail)

	otpCode := pkg.GenerateOTP()

	msg := fmt.Sprintf("Your otp code is: %v, for registering CRM system. Don't give it to anyone", otpCode)

	err := a.redis.SetX(ctx, loginRequest.Mail, otpCode, time.Minute*2)
	if err != nil {
		a.log.Error("error while setting otpCode to redis manager register", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	err = smtp.SendMail(loginRequest.Mail, msg)
	if err != nil {
		a.log.Error("error while sending otp code to manager register", logger.Error(err))
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (a *ManagerService) RegisterConfirm(ctx context.Context, req *manager_service.ManagerRegisterConfRequest) (*manager_service.ManagerLoginResponse, error) {
	resp := &manager_service.ManagerLoginResponse{}

	otp, err := a.redis.Get(ctx, req.Mail)
	if err != nil {
		a.log.Error("error while getting otp code for manager register confirm", logger.Error(err))
		return resp, err
	}
	if req.Otp != otp {
		a.log.Error("incorrect otp code for manager register confirm", logger.Error(err))
		return resp, errors.New("incorrect otp code")
	}
	req.Manager[0].Email = req.Mail

	id, err := a.strg.Manager().Create(ctx, req.Manager[0])
	if err != nil {
		a.log.Error("error while creating manager", logger.Error(err))
		return resp, err
	}
	var m = make(map[interface{}]interface{})

	m["user_id"] = id
	m["user_role"] = config.MANAGER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for manager register confirm", logger.Error(err))
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}

func (f *ManagerService) ChangePassword(ctx context.Context, pass *manager_service.ManagerChangePassword) (*manager_service.ManagerChangePasswordResp, error) {
	f.log.Info("---ChangePassword--->>>", logger.Any("req", pass))

	resp, err := f.strg.Manager().ChangePassword(ctx, pass)
	if err != nil {
		f.log.Error("---ChangePassword--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
