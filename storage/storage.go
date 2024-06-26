package storage

import (
	"context"
	"go_user_service/genproto/branch_service"
	"go_user_service/genproto/event_registrate_service"
	"go_user_service/genproto/event_service"
	"go_user_service/genproto/group_service"
	"go_user_service/genproto/student_service"
	"go_user_service/genproto/teacher_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageI interface {
	CloseDB()
	Teacher() TeacherRepoI
	Branch() BranchRepoI
	Group() GroupRepoI
	Student() StudentRepoI
	Event() EventRepoI
	EventRegistrate() EventRegistrateRepoI
}

type TeacherRepoI interface {
	Create(context.Context, *teacher_service.CreateTeacher) (*teacher_service.GetTeacher, error)
	Update(context.Context, *teacher_service.UpdateTeacher) (*teacher_service.GetTeacher, error)
	GetAll(context.Context, *teacher_service.GetListTeacherRequest) (*teacher_service.GetListTeacherResponse, error)
	GetById(context.Context, *teacher_service.TeacherPrimaryKey) (*teacher_service.GetTeacher, error)
	Delete(context.Context, *teacher_service.TeacherPrimaryKey) (emptypb.Empty, error)
}

type BranchRepoI interface {
	Create(context.Context, *branch_service.CreateBranch) (*branch_service.GetBranch, error)
	Update(context.Context, *branch_service.UpdateBranch) (*branch_service.GetBranch, error)
	GetAll(context.Context, *branch_service.GetListBranchRequest) (*branch_service.GetListBranchResponse, error)
	GetById(context.Context, *branch_service.BranchPrimaryKey) (*branch_service.GetBranch, error)
	Delete(context.Context, *branch_service.BranchPrimaryKey) (emptypb.Empty, error)
}

type GroupRepoI interface {
	Create(context.Context, *group_service.CreateGroup) (*group_service.GetGroup, error)
	Update(context.Context, *group_service.UpdateGroup) (*group_service.GetGroup, error)
	GetAll(context.Context, *group_service.GetListGroupRequest) (*group_service.GetListGroupResponse, error)
	GetById(context.Context, *group_service.GroupPrimaryKey) (*group_service.GetGroup, error)
	Delete(context.Context, *group_service.GroupPrimaryKey) (emptypb.Empty, error)
}

type StudentRepoI interface {
	Create(context.Context, *student_service.CreateStudent) (*student_service.GetStudent, error)
	Update(context.Context, *student_service.UpdateStudent) (*student_service.GetStudent, error)
	GetAll(context.Context, *student_service.GetListStudentRequest) (*student_service.GetListStudentResponse, error)
	GetById(context.Context, *student_service.StudentPrimaryKey) (*student_service.GetStudent, error)
	Delete(context.Context, *student_service.StudentPrimaryKey) (emptypb.Empty, error)
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
}
