package postgres

import (
	"context"
	"database/sql"
	"fmt"
	tc "go_user_service/genproto/teacher_service"
	"go_user_service/pkg"
	"go_user_service/pkg/hash"
	"go_user_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type teacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacherRepo(db *pgxpool.Pool) storage.TeacherRepoI {
	return &teacherRepo{
		db: db,
	}
}

func generateLogin(db *pgxpool.Pool, ctx context.Context) (string, error) {
	var nextVal int
	err := db.QueryRow(ctx, "SELECT nextval('teacher_external_id_seq')").Scan(&nextVal)
	if err != nil {
		return "", err
	}
	userLogin := "" + fmt.Sprintf("%06d", nextVal)
	return userLogin, nil
}

func (c *teacherRepo) Create(ctx context.Context, req *tc.CreateTeacher) (*tc.GetTeacher, error) {

	id := uuid.NewString()
	pasword, err := hash.HashPassword(req.UserPassword)
	if err != nil {
		log.Println("error while hashing password", err)
	}

	userLogin, err := generateLogin(c.db, ctx)
	if err != nil {
		log.Println("error while generating login", err)
	}
	comtag, err := c.db.Exec(ctx, `
		INSERT INTO teachers (
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
		req.EndWorking)
	if err != nil {
		log.Println("error while creating teacher", comtag)
		return nil, err
	}

	teacher, err := c.GetById(ctx, &tc.TeacherPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting teacher by id")
		return nil, err
	}
	return teacher, nil
}

func (c *teacherRepo) Update(ctx context.Context, req *tc.UpdateTeacher) (*tc.GetTeacher, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE teachers SET
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
		req.EndWorking,
		req.Id)
	if err != nil {
		log.Println("error while updating teacher")
		return nil, err
	}

	teacher, err := c.GetById(ctx, &tc.TeacherPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting teacher by id")
		return nil, err
	}
	return teacher, nil
}

func (c *teacherRepo) GetAll(ctx context.Context, req *tc.GetListTeacherRequest) (*tc.GetListTeacherResponse, error) {
	teachers := tc.GetListTeacherResponse{}
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
				user_password,
				salary,
				ielts_score,
				ielts_attempts_count,
				start_working,
				end_working,
				created_at,
				updated_at
			FROM teachers
			WHERE TRUE AND deleted_at is null ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting all teachers")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			teacher tc.GetTeacher
		)
		if err = rows.Scan(
			&teacher.Id,
			&teacher.BranchId,
			&teacher.UserLogin,
			&teacher.Birthday,
			&teacher.Gender,
			&teacher.Fullname,
			&teacher.Email,
			&teacher.Phone,
			&teacher.UserPassword,
			&teacher.Salary,
			&teacher.IeltsScore,
			&teacher.IeltsAttemptsCount,
			&start_working,
			&end_working,
			&created_at,
			&updated_at,
		); err != nil {
			return &teachers, err
		}
		teacher.StartWorking = pkg.NullStringToString(start_working)
		teacher.EndWorking = pkg.NullStringToString(end_working)
		teacher.CreatedAt = pkg.NullStringToString(created_at)
		teacher.UpdatedAt = pkg.NullStringToString(updated_at)

		teachers.Teachers = append(teachers.Teachers, &teacher)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from teacher WHERE TRUE AND deleted_at is null `+filter_by_name+``).Scan(&teachers.Count)
	if err != nil {
		return &teachers, err
	}

	return &teachers, nil
}

func (c *teacherRepo) GetById(ctx context.Context, id *tc.TeacherPrimaryKey) (*tc.GetTeacher, error) {
	var (
		teacher       tc.GetTeacher
		created_at    sql.NullString
		updated_at    sql.NullString
		start_working sql.NullString
		end_working   sql.NullString
	)

	query := `SELECT
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
				end_working,
				created_at,
				updated_at
			FROM teachers
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&teacher.Id,
		&teacher.BranchId,
		&teacher.UserLogin,
		&teacher.Birthday,
		&teacher.Gender,
		&teacher.Fullname,
		&teacher.Email,
		&teacher.Phone,
		&teacher.UserPassword,
		&teacher.Salary,
		&teacher.IeltsScore,
		&teacher.IeltsAttemptsCount,
		&start_working,
		&end_working,
		&created_at,
		&updated_at); err != nil {
		return &teacher, err
	}
	teacher.CreatedAt = pkg.NullStringToString(created_at)
	teacher.UpdatedAt = pkg.NullStringToString(updated_at)
	teacher.CreatedAt = pkg.NullStringToString(created_at)
	teacher.UpdatedAt = pkg.NullStringToString(updated_at)

	return &teacher, nil
}

func (c *teacherRepo) Delete(ctx context.Context, id *tc.TeacherPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE teachers SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
