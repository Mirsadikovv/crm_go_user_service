package postgres

import (
	"context"
	"database/sql"
	"fmt"
	tc "go_user_service/genproto/administrator_service"
	"go_user_service/pkg"
	"go_user_service/pkg/hash"
	"go_user_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type administratorRepo struct {
	db *pgxpool.Pool
}

func NewAdministratorRepo(db *pgxpool.Pool) storage.AdministratorRepoI {
	return &administratorRepo{
		db: db,
	}
}

func generateAdministratorLogin(db *pgxpool.Pool, ctx context.Context) (string, error) {
	var nextVal int
	err := db.QueryRow(ctx, "SELECT nextval('administrator_external_id_seq')").Scan(&nextVal)
	if err != nil {
		return "", err
	}
	userLogin := "A" + fmt.Sprintf("%05d", nextVal)
	return userLogin, nil
}

func (c *administratorRepo) Create(ctx context.Context, req *tc.CreateAdministrator) (*tc.GetAdministrator, error) {
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

	userLogin, err := generateAdministratorLogin(c.db, ctx)
	if err != nil {
		log.Println("error while generating login", err)
	}
	comtag, err := c.db.Exec(ctx, `
		INSERT INTO administrators (
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
			start_working,
			end_working
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
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
		req.StartWorking,
		end_working)
	if err != nil {
		log.Println("error while creating administrator", comtag)
		return nil, err
	}

	administrator, err := c.GetById(ctx, &tc.AdministratorPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting administrator by id")
		return nil, err
	}
	return administrator, nil
}

func (c *administratorRepo) Update(ctx context.Context, req *tc.UpdateAdministrator) (*tc.GetAdministrator, error) {
	var end_working sql.NullString
	if req.EndWorking == "" {
		end_working = sql.NullString{Valid: false}
	} else {
		end_working = sql.NullString{String: req.EndWorking, Valid: true}
	}
	_, err := c.db.Exec(ctx, `
		UPDATE administrators SET
		branch_id = $1,
		birthday = $2,
		gender = $3,
		fullname = $4,
		email = $5,
		phone = $6,
		salary = $7,
		start_working = $8,    
		end_working = $9,
		updated_at = NOW()
		WHERE id = $10
		`,
		req.BranchId,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		req.Salary,
		req.StartWorking,
		end_working,
		req.Id)
	if err != nil {
		log.Println("error while updating administrator")
		return nil, err
	}

	administrator, err := c.GetById(ctx, &tc.AdministratorPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting administrator by id")
		return nil, err
	}
	return administrator, nil
}

func (c *administratorRepo) GetAll(ctx context.Context, req *tc.GetListAdministratorRequest) (*tc.GetListAdministratorResponse, error) {
	administrators := tc.GetListAdministratorResponse{}
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
				start_working,
				end_working,
				created_at,
				updated_at
			FROM administrators
			WHERE TRUE AND deleted_at is null ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting all administrators")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			administrator tc.GetAdministrator
		)
		if err = rows.Scan(
			&administrator.Id,
			&administrator.BranchId,
			&administrator.UserLogin,
			&administrator.Birthday,
			&administrator.Gender,
			&administrator.Fullname,
			&administrator.Email,
			&administrator.Phone,
			&administrator.Salary,
			&start_working,
			&end_working,
			&created_at,
			&updated_at,
		); err != nil {
			return &administrators, err
		}
		administrator.StartWorking = pkg.NullStringToString(start_working)
		administrator.EndWorking = pkg.NullStringToString(end_working)
		administrator.CreatedAt = pkg.NullStringToString(created_at)
		administrator.UpdatedAt = pkg.NullStringToString(updated_at)

		administrators.Administrators = append(administrators.Administrators, &administrator)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from administrators WHERE TRUE AND deleted_at is null `+filter_by_name+``).Scan(&administrators.Count)
	if err != nil {
		return &administrators, err
	}

	return &administrators, nil
}

func (c *administratorRepo) GetById(ctx context.Context, id *tc.AdministratorPrimaryKey) (*tc.GetAdministrator, error) {
	var (
		administrator tc.GetAdministrator
		created_at    sql.NullString
		updated_at    sql.NullString
		start_working sql.NullString
		end_working   sql.NullString
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
				start_working,
				end_working,
				created_at,
				updated_at
			FROM administrators
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&administrator.Id,
		&administrator.BranchId,
		&administrator.Birthday,
		&administrator.Gender,
		&administrator.Fullname,
		&administrator.Email,
		&administrator.Phone,
		&administrator.Salary,
		&start_working,
		&end_working,
		&created_at,
		&updated_at); err != nil {
		return &administrator, err
	}
	administrator.StartWorking = pkg.NullStringToString(start_working)
	administrator.EndWorking = pkg.NullStringToString(end_working)
	administrator.CreatedAt = pkg.NullStringToString(created_at)
	administrator.UpdatedAt = pkg.NullStringToString(updated_at)

	return &administrator, nil
}

func (c *administratorRepo) Delete(ctx context.Context, id *tc.AdministratorPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE administrators SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
