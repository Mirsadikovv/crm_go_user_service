package storage

import (
	"context"
	"go_user_service/genproto/administrator_service"
	"go_user_service/genproto/branch_service"
	"go_user_service/genproto/event_registrate_service"
	"go_user_service/genproto/event_service"
	"go_user_service/genproto/group_service"
	"go_user_service/genproto/manager_service"
	"go_user_service/genproto/student_service"
	"go_user_service/genproto/superadmin_service"
	"go_user_service/genproto/support_teacher_service"
	"go_user_service/genproto/teacher_service"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageI interface {
	CloseDB()
	Teacher() TeacherRepoI
	SupportTeacher() SupportTeacherRepoI
	Manager() ManagerRepoI
	Administrator() AdministratorRepoI
	Superadmin() SuperadminRepoI
	Branch() BranchRepoI
	Group() GroupRepoI
	Student() StudentRepoI
	Event() EventRepoI
	EventRegistrate() EventRegistrateRepoI
	Redis() IRedisStorage
}

type TeacherRepoI interface {
	Create(context.Context, *teacher_service.CreateTeacher) (*teacher_service.GetTeacher, error)
	Update(context.Context, *teacher_service.UpdateTeacher) (*teacher_service.GetTeacher, error)
	GetAll(context.Context, *teacher_service.GetListTeacherRequest) (*teacher_service.GetListTeacherResponse, error)
	GetById(context.Context, *teacher_service.TeacherPrimaryKey) (*teacher_service.GetTeacher, error)
	Delete(context.Context, *teacher_service.TeacherPrimaryKey) (emptypb.Empty, error)
	ChangePassword(context.Context, *teacher_service.TeacherChangePassword) (*teacher_service.TeacherChangePasswordResp, error)
	GetByLogin(context.Context, string) (*teacher_service.GetTeacherByLogin, error)
	GetPassword(context.Context, string) (string, error)
}

type SupportTeacherRepoI interface {
	Create(context.Context, *support_teacher_service.CreateSupportTeacher) (*support_teacher_service.GetSupportTeacher, error)
	Update(context.Context, *support_teacher_service.UpdateSupportTeacher) (*support_teacher_service.GetSupportTeacher, error)
	GetAll(context.Context, *support_teacher_service.GetListSupportTeacherRequest) (*support_teacher_service.GetListSupportTeacherResponse, error)
	GetById(context.Context, *support_teacher_service.SupportTeacherPrimaryKey) (*support_teacher_service.GetSupportTeacher, error)
	Delete(context.Context, *support_teacher_service.SupportTeacherPrimaryKey) (emptypb.Empty, error)
	ChangePassword(context.Context, *support_teacher_service.SupportTeacherChangePassword) (*support_teacher_service.SupportTeacherChangePasswordResp, error)
	GetByLogin(context.Context, string) (*support_teacher_service.GetSupportTeacherByLogin, error)
	GetPassword(context.Context, string) (string, error)
}

type BranchRepoI interface {
	Create(context.Context, *branch_service.CreateBranch) (*branch_service.GetBranch, error)
	Update(context.Context, *branch_service.UpdateBranch) (*branch_service.GetBranch, error)
	GetAll(context.Context, *branch_service.GetListBranchRequest) (*branch_service.GetListBranchResponse, error)
	GetById(context.Context, *branch_service.BranchPrimaryKey) (*branch_service.GetBranch, error)
	Delete(context.Context, *branch_service.BranchPrimaryKey) (emptypb.Empty, error)
}

type ManagerRepoI interface {
	Create(context.Context, *manager_service.CreateManager) (*manager_service.GetManager, error)
	Update(context.Context, *manager_service.UpdateManager) (*manager_service.GetManager, error)
	GetAll(context.Context, *manager_service.GetListManagerRequest) (*manager_service.GetListManagerResponse, error)
	GetById(context.Context, *manager_service.ManagerPrimaryKey) (*manager_service.GetManager, error)
	Delete(context.Context, *manager_service.ManagerPrimaryKey) (emptypb.Empty, error)
	ChangePassword(context.Context, *manager_service.ManagerChangePassword) (*manager_service.ManagerChangePasswordResp, error)
	GetByLogin(context.Context, string) (*manager_service.GetManagerByLogin, error)
	GetPassword(context.Context, string) (string, error)
}

type AdministratorRepoI interface {
	Create(context.Context, *administrator_service.CreateAdministrator) (*administrator_service.GetAdministrator, error)
	Update(context.Context, *administrator_service.UpdateAdministrator) (*administrator_service.GetAdministrator, error)
	GetAll(context.Context, *administrator_service.GetListAdministratorRequest) (*administrator_service.GetListAdministratorResponse, error)
	GetById(context.Context, *administrator_service.AdministratorPrimaryKey) (*administrator_service.GetAdministrator, error)
	Delete(context.Context, *administrator_service.AdministratorPrimaryKey) (emptypb.Empty, error)
	ChangePassword(context.Context, *administrator_service.AdministratorChangePassword) (*administrator_service.AdministratorChangePasswordResp, error)
	GetByLogin(context.Context, string) (*administrator_service.GetAdministratorByLogin, error)
	GetPassword(context.Context, string) (string, error)
}

type SuperadminRepoI interface {
	Create(context.Context, *superadmin_service.CreateSuperadmin) (*superadmin_service.GetSuperadmin, error)
	Update(context.Context, *superadmin_service.UpdateSuperadmin) (*superadmin_service.GetSuperadmin, error)
	GetById(context.Context, *superadmin_service.SuperadminPrimaryKey) (*superadmin_service.GetSuperadmin, error)
	Delete(context.Context, *superadmin_service.SuperadminPrimaryKey) (emptypb.Empty, error)
	ChangePassword(context.Context, *superadmin_service.SuperadminChangePassword) (*superadmin_service.SuperadminChangePasswordResp, error)
	GetByLogin(context.Context, string) (*superadmin_service.GetSuperadminByLogin, error)
	GetPassword(context.Context, string) (string, error)
}

type GroupRepoI interface {
	Create(context.Context, *group_service.CreateGroup) (*group_service.GetGroup, error)
	Update(context.Context, *group_service.UpdateGroup) (*group_service.GetGroup, error)
	GetAll(context.Context, *group_service.GetListGroupRequest) (*group_service.GetListGroupResponse, error)
	GetById(context.Context, *group_service.GroupPrimaryKey) (*group_service.GetGroup, error)
	Delete(context.Context, *group_service.GroupPrimaryKey) (emptypb.Empty, error)
	Check(context.Context, *group_service.GroupPrimaryKey) (*group_service.CheckGroupResp, error)
	GetTBS(context.Context, *group_service.GroupPrimaryKey) (*group_service.GetTBSresp, error)
}

type StudentRepoI interface {
	Create(context.Context, *student_service.CreateStudent) (*student_service.GetStudent, error)
	Update(context.Context, *student_service.UpdateStudent) (*student_service.GetStudent, error)
	GetAll(context.Context, *student_service.GetListStudentRequest) (*student_service.GetListStudentResponse, error)
	GetById(context.Context, *student_service.StudentPrimaryKey) (*student_service.GetStudent, error)
	Delete(context.Context, *student_service.StudentPrimaryKey) (emptypb.Empty, error)
	Check(context.Context, *student_service.StudentPrimaryKey) (*student_service.CheckStudentResp, error)
	ChangePassword(context.Context, *student_service.StudentChangePassword) (*student_service.StudentChangePasswordResp, error)
	GetByLogin(context.Context, string) (*student_service.GetStudentByLogin, error)
	GetPassword(context.Context, string) (string, error)
}

type EventRepoI interface {
	Create(context.Context, *event_service.CreateEvent) (*event_service.GetEvent, error)
	Update(context.Context, *event_service.UpdateEvent) (*event_service.GetEvent, error)
	GetAll(context.Context, *event_service.GetListEventRequest) (*event_service.GetListEventResponse, error)
	GetById(context.Context, *event_service.EventPrimaryKey) (*event_service.GetEvent, error)
	Delete(context.Context, *event_service.EventPrimaryKey) (emptypb.Empty, error)
}

type EventRegistrateRepoI interface {
	Create(context.Context, *event_registrate_service.CreateEventRegistrate) (*event_registrate_service.GetEventRegistrate, error)
	Update(context.Context, *event_registrate_service.UpdateEventRegistrate) (*event_registrate_service.GetEventRegistrate, error)
	GetById(context.Context, *event_registrate_service.EventRegistratePrimaryKey) (*event_registrate_service.GetEventRegistrate, error)
	Delete(context.Context, *event_registrate_service.EventRegistratePrimaryKey) (emptypb.Empty, error)
	GetStudentEvent(context.Context, *event_registrate_service.GetListEventRegistrateRequest) (*event_registrate_service.GetListEventRegistrateResponse, error)
}

type IRedisStorage interface {
	SetX(context.Context, string, interface{}, time.Duration) error
	Get(context.Context, string) (interface{}, error)
	Del(context.Context, string) error
}
