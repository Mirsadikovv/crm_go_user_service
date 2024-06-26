package postgres

import (
	"context"
	"fmt"
	"go_user_service/config"
	"go_user_service/storage"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db              *pgxpool.Pool
	teacher         storage.TeacherRepoI
	branch          storage.BranchRepoI
	group           storage.GroupRepoI
	student         storage.StudentRepoI
	event           storage.EventRepoI
	eventRegistrate storage.EventRegistrateRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) Teacher() storage.TeacherRepoI {
	if s.teacher == nil {
		s.teacher = NewTeacherRepo(s.db)
	}
	return s.teacher
}

func (s *Store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}
	return s.branch
}

func (s *Store) Group() storage.GroupRepoI {
	if s.group == nil {
		s.group = NewGroupRepo(s.db)
	}
	return s.group
}

func (s *Store) Student() storage.StudentRepoI {
	if s.student == nil {
		s.student = NewStudentRepo(s.db)
	}
	return s.student
}

func (s *Store) Event() storage.EventRepoI {
	if s.event == nil {
		s.event = NewEventRepo(s.db)
	}
	return s.event
}

func (s *Store) EventRegistrate() storage.EventRegistrateRepoI {
	if s.eventRegistrate == nil {
		s.eventRegistrate = NewEventRegistrateRepo(s.db)
	}
	return s.eventRegistrate
}
