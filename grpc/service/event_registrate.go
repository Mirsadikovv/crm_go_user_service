package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/event_registrate_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventRegistrateService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewEventRegistrateService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *EventRegistrateService {
	return &EventRegistrateService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *EventRegistrateService) Create(ctx context.Context, req *event_registrate_service.CreateEventRegistrate) (*event_registrate_service.GetEventRegistrate, error) {

	f.log.Info("---CreateEventRegistrate--->>>", logger.Any("req", req))

	resp, err := f.strg.EventRegistrate().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateEventRegistrate--->>>", logger.Error(err))
		return &event_registrate_service.GetEventRegistrate{}, err
	}

	return resp, nil
}
func (f *EventRegistrateService) Update(ctx context.Context, req *event_registrate_service.UpdateEventRegistrate) (*event_registrate_service.GetEventRegistrate, error) {

	f.log.Info("---UpdateEventRegistrate--->>>", logger.Any("req", req))

	resp, err := f.strg.EventRegistrate().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateEventRegistrate--->>>", logger.Error(err))
		return &event_registrate_service.GetEventRegistrate{}, err
	}

	return resp, nil
}

func (f *EventRegistrateService) GetByID(ctx context.Context, id *event_registrate_service.EventRegistratePrimaryKey) (*event_registrate_service.GetEventRegistrate, error) {
	f.log.Info("---GetEventRegistrate--->>>", logger.Any("req", id))

	resp, err := f.strg.EventRegistrate().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetEventRegistrate--->>>", logger.Error(err))
		return &event_registrate_service.GetEventRegistrate{}, err
	}

	return resp, nil
}

func (f *EventRegistrateService) Delete(ctx context.Context, req *event_registrate_service.EventRegistratePrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteEventRegistrate--->>>", logger.Any("req", req))

	_, err := f.strg.EventRegistrate().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteEventRegistrate--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
