package storage

import (
	"context"
	"go_user_service/genproto/teacher_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageI interface {
	CloseDB()
	Teacher() TeacherRepoI
}

type TeacherRepoI interface {
	Create(context.Context, *teacher_service.CreateTeacher) (*teacher_service.GetTeacher, error)
	Update(context.Context, *teacher_service.UpdateTeacher) (*teacher_service.GetTeacher, error)
	GetAll(context.Context, *teacher_service.GetListTeacherRequest) (*teacher_service.GetListTeacherResponse, error)
	GetById(context.Context, *teacher_service.TeacherPrimaryKey) (*teacher_service.GetTeacher, error)
	Delete(context.Context, *teacher_service.TeacherPrimaryKey) (emptypb.Empty, error)
}
