package service

import (
	"context"
	"go_user_service/config"
	"go_user_service/genproto/event_service"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewEventService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *EventService {
	return &EventService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *EventService) Create(ctx context.Context, req *event_service.CreateEvent) (*event_service.GetEvent, error) {

	f.log.Info("---CreateEvent--->>>", logger.Any("req", req))

	resp, err := f.strg.Event().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateEvent--->>>", logger.Error(err))
		return &event_service.GetEvent{}, err
	}

	return resp, nil
}
func (f *EventService) Update(ctx context.Context, req *event_service.UpdateEvent) (*event_service.GetEvent, error) {

	f.log.Info("---UpdateEvent--->>>", logger.Any("req", req))

	resp, err := f.strg.Event().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateEvent--->>>", logger.Error(err))
		return &event_service.GetEvent{}, err
	}

	return resp, nil
}

func (f *EventService) GetList(ctx context.Context, req *event_service.GetListEventRequest) (*event_service.GetListEventResponse, error) {
	f.log.Info("---GetListEvent--->>>", logger.Any("req", req))

	resp, err := f.strg.Event().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListEvent--->>>", logger.Error(err))
		return &event_service.GetListEventResponse{}, err
	}

	return resp, nil
}

func (f *EventService) GetByID(ctx context.Context, id *event_service.EventPrimaryKey) (*event_service.GetEvent, error) {
	f.log.Info("---GetEvent--->>>", logger.Any("req", id))

	resp, err := f.strg.Event().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetEvent--->>>", logger.Error(err))
		return &event_service.GetEvent{}, err
	}

	return resp, nil
}

func (f *EventService) Delete(ctx context.Context, req *event_service.EventPrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteEvent--->>>", logger.Any("req", req))

	_, err := f.strg.Event().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteEvent--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
