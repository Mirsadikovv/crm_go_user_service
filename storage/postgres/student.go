package postgres

import (
	"context"
	"database/sql"
	"fmt"
	tc "go_user_service/genproto/student_service"
	"go_user_service/pkg"
	"go_user_service/pkg/hash"
	"go_user_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type studentRepo struct {
	db *pgxpool.Pool
}

func NewStudentRepo(db *pgxpool.Pool) storage.StudentRepoI {
	return &studentRepo{
		db: db,
	}
}

func generateStudentLogin(db *pgxpool.Pool, ctx context.Context) (string, error) {
	var nextVal int
	err := db.QueryRow(ctx, "SELECT nextval('student_external_id_seq')").Scan(&nextVal)
	if err != nil {
		return "", err
	}
	userLogin := "S" + fmt.Sprintf("%05d", nextVal)
	return userLogin, nil
}

func (c *studentRepo) Create(ctx context.Context, req *tc.CreateStudent) (*tc.GetStudent, error) {

	id := uuid.NewString()
	pasword, err := hash.HashPassword(req.UserPassword)
	if err != nil {
		log.Println("error while hashing password", err)
	}

	userLogin, err := generateStudentLogin(c.db, ctx)
	if err != nil {
		log.Println("error while generating login", err)
	}

	comtag, err := c.db.Exec(ctx, `
		INSERT INTO students (
			id,
			group_id,
			user_login,
			birthday,
			gender,
			fullname,
			email,
			phone,
			user_password,
			paid_sum,
			started_at,
			finished_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
		)`,
		id,
		req.GroupId,
		userLogin,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		pasword,
		req.PaidSum,
		req.StartedAt,
		req.FinishedAt)
	if err != nil {
		log.Println("error while creating student", comtag)
		return nil, err
	}

	student, err := c.GetById(ctx, &tc.StudentPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting student by id")
		return nil, err
	}
	return student, nil
}

func (c *studentRepo) Update(ctx context.Context, req *tc.UpdateStudent) (*tc.GetStudent, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE students SET
		group_id = $1,
		birthday = $2,
		gender = $3,
		fullname = $4,
		email = $5,
		phone = $6,
		paid_sum = $7,
		started_at = $8,
		finished_at = $9,
		updated_at = NOW()
		WHERE id = $10
		`,
		req.GroupId,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		req.PaidSum,
		req.StartedAt,
		req.FinishedAt,
		req.Id)
	if err != nil {
		log.Println("error while updating student")
		return nil, err
	}

	student, err := c.GetById(ctx, &tc.StudentPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting student by id")
		return nil, err
	}
	return student, nil
}

func (c *studentRepo) GetAll(ctx context.Context, req *tc.GetListStudentRequest) (*tc.GetListStudentResponse, error) {
	students := tc.GetListStudentResponse{}
	var (
		created_at  sql.NullString
		updated_at  sql.NullString
		started_at  sql.NullString
		finished_at sql.NullString
	)
	filter_by_name := ""
	offest := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter_by_name = fmt.Sprintf(`AND fullname ILIKE '%%%v%%'`, req.Search)
	}
	query := `SELECT
				id,
				group_id,
				user_login,
				birthday,
				gender,
				fullname,
				email,
				phone,
				paid_sum,
				start_at,
				finished_at,
				created_at,
				updated_at
			FROM students
			WHERE TRUE AND deleted_at is null ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting all students")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			student tc.GetStudent
		)
		if err = rows.Scan(
			&student.Id,
			&student.GroupId,
			&student.UserLogin,
			&student.Birthday,
			&student.Gender,
			&student.Fullname,
			&student.Email,
			&student.Phone,
			&student.PaidSum,
			&started_at,
			&finished_at,
			&created_at,
			&updated_at,
		); err != nil {
			return &students, err
		}
		student.StartedAt = pkg.NullStringToString(started_at)
		student.FinishedAt = pkg.NullStringToString(finished_at)
		student.CreatedAt = pkg.NullStringToString(created_at)
		student.UpdatedAt = pkg.NullStringToString(updated_at)

		students.Students = append(students.Students, &student)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from students WHERE TRUE AND deleted_at is null `+filter_by_name+``).Scan(&students.Count)
	if err != nil {
		return &students, err
	}

	return &students, nil
}

func (c *studentRepo) GetById(ctx context.Context, id *tc.StudentPrimaryKey) (*tc.GetStudent, error) {
	var (
		student     tc.GetStudent
		created_at  sql.NullString
		updated_at  sql.NullString
		started_at  sql.NullString
		finished_at sql.NullString
	)

	query := `SELECT
				id,
				group_id,
				user_login,
				birthday,
				gender,
				fullname,
				email,
				phone,
				paid_sum,
				start_at,
				finished_at,
				created_at,
				updated_at
			FROM students
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&student.Id,
		&student.GroupId,
		&student.UserLogin,
		&student.Birthday,
		&student.Gender,
		&student.Fullname,
		&student.Email,
		&student.Phone,
		&student.PaidSum,
		&started_at,
		&finished_at,
		&created_at,
		&updated_at); err != nil {
		return &student, err
	}
	student.StartedAt = pkg.NullStringToString(started_at)
	student.FinishedAt = pkg.NullStringToString(finished_at)
	student.CreatedAt = pkg.NullStringToString(created_at)
	student.UpdatedAt = pkg.NullStringToString(updated_at)

	return &student, nil
}

func (c *studentRepo) Delete(ctx context.Context, id *tc.StudentPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE students SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
