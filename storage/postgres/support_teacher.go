package postgres

import (
	"context"
	"database/sql"
	"fmt"
	stc "go_user_service/genproto/support_teacher_service"
	"go_user_service/pkg"
	"go_user_service/pkg/hash"
	"go_user_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type support_teacherRepo struct {
	db *pgxpool.Pool
}

func NewSupportTeacherRepo(db *pgxpool.Pool) storage.SupportTeacherRepoI {
	return &support_teacherRepo{
		db: db,
	}
}

func generateSupportTeacherLogin(db *pgxpool.Pool, ctx context.Context) (string, error) {
	var nextVal int
	err := db.QueryRow(ctx, "SELECT nextval('support_teacher_external_id_seq')").Scan(&nextVal)
	if err != nil {
		return "", err
	}
	userLogin := "ST" + fmt.Sprintf("%05d", nextVal)
	return userLogin, nil
}

func (c *support_teacherRepo) Create(ctx context.Context, req *stc.CreateSupportTeacher) (*stc.GetSupportTeacher, error) {
	var end_working sql.NullString
	if req.EndWorking == "" {
		end_working = sql.NullString{Valid: false}
	} else {
		end_working = sql.NullString{String: req.EndWorking, Valid: true}
	}
	id := uuid.NewString()
	pasword, err := hash.HashPassword(req.UserPassword)
	if err != nil {
		log.Println("error while hashing password", err)
	}

	userLogin, err := generateSupportTeacherLogin(c.db, ctx)
	if err != nil {
		log.Println("error while generating login", err)
	}
	comtag, err := c.db.Exec(ctx, `
		INSERT INTO support_teachers (
			id,
			branch_id,
			user_login,
			birthday,
			gender,
			fullname,
			email,
			phone,
			user_password,
			salary,
			ielts_score,
			ielts_attempts_count,
			start_working,
			end_working
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14
		)`,
		id,
		req.BranchId,
		userLogin,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		pasword,
		req.Salary,
		req.IeltsScore,
		req.IeltsAttemptsCount,
		req.StartWorking,
		end_working)
	if err != nil {
		log.Println("error while creating support_teacher", comtag)
		return nil, err
	}

	support_teacher, err := c.GetById(ctx, &stc.SupportTeacherPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting support_teacher by id")
		return nil, err
	}
	return support_teacher, nil
}

func (c *support_teacherRepo) Update(ctx context.Context, req *stc.UpdateSupportTeacher) (*stc.GetSupportTeacher, error) {
	var end_working sql.NullString
	if req.EndWorking == "" {
		end_working = sql.NullString{Valid: false}
	} else {
		end_working = sql.NullString{String: req.EndWorking, Valid: true}
	}
	_, err := c.db.Exec(ctx, `
		UPDATE support_teachers SET
		branch_id = $1,
		birthday = $2,
		gender = $3,
		fullname = $4,
		email = $5,
		phone = $6,
		salary = $7,
		ielts_score = $8,
		ielts_attempts_count = $9,
		start_working = $10,
		end_working = $11,
		updated_at = NOW()
		WHERE id = $12
		`,
		req.BranchId,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		req.Salary,
		req.IeltsScore,
		req.IeltsAttemptsCount,
		req.StartWorking,
		end_working,
		req.Id)
	if err != nil {
		log.Println("error while updating support_teacher")
		return nil, err
	}

	support_teacher, err := c.GetById(ctx, &stc.SupportTeacherPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting support_teacher by id")
		return nil, err
	}
	return support_teacher, nil
}

func (c *support_teacherRepo) GetAll(ctx context.Context, req *stc.GetListSupportTeacherRequest) (*stc.GetListSupportTeacherResponse, error) {
	support_teachers := stc.GetListSupportTeacherResponse{}
	var (
		created_at    sql.NullString
		updated_at    sql.NullString
		start_working sql.NullString
		end_working   sql.NullString
	)
	filter_by_name := ""
	offest := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter_by_name = fmt.Sprintf(`AND fullname ILIKE '%%%v%%'`, req.Search)
	}
	query := `SELECT
				id,
				branch_id,
				user_login,
				birthday,
				gender,
				fullname,
				email,
				phone,
				salary,
				ielts_score,
				ielts_attempts_count,
				start_working,
				end_working,
				created_at,
				updated_at
			FROM support_teachers
			WHERE TRUE AND deleted_at is null ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting all support_teachers")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			support_teacher stc.GetSupportTeacher
		)
		if err = rows.Scan(
			&support_teacher.Id,
			&support_teacher.BranchId,
			&support_teacher.UserLogin,
			&support_teacher.Birthday,
			&support_teacher.Gender,
			&support_teacher.Fullname,
			&support_teacher.Email,
			&support_teacher.Phone,
			&support_teacher.Salary,
			&support_teacher.IeltsScore,
			&support_teacher.IeltsAttemptsCount,
			&start_working,
			&end_working,
			&created_at,
			&updated_at,
		); err != nil {
			return &support_teachers, err
		}
		support_teacher.StartWorking = pkg.NullStringToString(start_working)
		support_teacher.EndWorking = pkg.NullStringToString(end_working)
		support_teacher.CreatedAt = pkg.NullStringToString(created_at)
		support_teacher.UpdatedAt = pkg.NullStringToString(updated_at)

		support_teachers.SupportTeachers = append(support_teachers.SupportTeachers, &support_teacher)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from support_teachers WHERE TRUE AND deleted_at is null `+filter_by_name+``).Scan(&support_teachers.Count)
	if err != nil {
		return &support_teachers, err
	}

	return &support_teachers, nil
}

func (c *support_teacherRepo) GetById(ctx context.Context, id *stc.SupportTeacherPrimaryKey) (*stc.GetSupportTeacher, error) {
	var (
		support_teacher stc.GetSupportTeacher
		created_at      sql.NullString
		updated_at      sql.NullString
		start_working   sql.NullString
		end_working     sql.NullString
	)

	query := `SELECT
				id,
				branch_id,
				birthday,
				gender,
				fullname,
				email,
				phone,
				salary,
				ielts_score,
				ielts_attempts_count,
				start_working,
				end_working,
				created_at,
				updated_at
			FROM support_teachers
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&support_teacher.Id,
		&support_teacher.BranchId,
		&support_teacher.Birthday,
		&support_teacher.Gender,
		&support_teacher.Fullname,
		&support_teacher.Email,
		&support_teacher.Phone,
		&support_teacher.Salary,
		&support_teacher.IeltsScore,
		&support_teacher.IeltsAttemptsCount,
		&start_working,
		&end_working,
		&created_at,
		&updated_at); err != nil {
		return &support_teacher, err
	}
	support_teacher.StartWorking = pkg.NullStringToString(start_working)
	support_teacher.EndWorking = pkg.NullStringToString(end_working)
	support_teacher.CreatedAt = pkg.NullStringToString(created_at)
	support_teacher.UpdatedAt = pkg.NullStringToString(updated_at)

	return &support_teacher, nil
}

func (c *support_teacherRepo) Delete(ctx context.Context, id *stc.SupportTeacherPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE support_teachers SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
