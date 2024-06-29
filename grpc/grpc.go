package grpc

import (
	"go_user_service/config"
	"go_user_service/genproto/administrator_service"
	"go_user_service/genproto/branch_service"
	"go_user_service/genproto/event_service"
	"go_user_service/genproto/group_service"
	"go_user_service/genproto/manager_service"
	"go_user_service/genproto/student_service"
	"go_user_service/genproto/superadmin_service"
	"go_user_service/genproto/support_teacher_service"
	"go_user_service/genproto/teacher_service"

	"go_user_service/grpc/client"
	"go_user_service/grpc/service"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI, redis storage.IRedisStorage) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	teacher_service.RegisterTeacherServiceServer(grpcServer, service.NewTeacherService(cfg, log, strg, srvc, redis))
	support_teacher_service.RegisterSupportTeacherServiceServer(grpcServer, service.NewSupportTeacherService(cfg, log, strg, srvc))
	branch_service.RegisterBranchServiceServer(grpcServer, service.NewBranchService(cfg, log, strg, srvc))
	group_service.RegisterGroupServiceServer(grpcServer, service.NewGroupService(cfg, log, strg, srvc))
	student_service.RegisterStudentServiceServer(grpcServer, service.NewStudentService(cfg, log, strg, srvc))
	event_service.RegisterEventServiceServer(grpcServer, service.NewEventService(cfg, log, strg, srvc))
	manager_service.RegisterManagerServiceServer(grpcServer, service.NewManagerService(cfg, log, strg, srvc))
	administrator_service.RegisterAdministratorServiceServer(grpcServer, service.NewAdministratorService(cfg, log, strg, srvc))
	superadmin_service.RegisterSuperadminServiceServer(grpcServer, service.NewSuperadminService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
