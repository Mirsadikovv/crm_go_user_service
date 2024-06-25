package grpc

import (
	"go_user_service/config"
	"go_user_service/genproto/teacher_service"

	"go_user_service/grpc/client"
	"go_user_service/grpc/service"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	teacher_service.RegisterTeacherServiceServer(grpcServer, service.NewTeacherService(cfg, log, strg, srvc))
	// product_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg, srvc))
	// review_service.RegisterReviewServiceServer(grpcServer, service.NewReviewService(cfg, log, strg, srvc))
	// product_categories_service.RegisterProductCategoriesServiceServer(grpcServer, service.NewProductCategoriesService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
