package postgres

import (
	"context"
	"database/sql"
	"fmt"
	tc "go_user_service/genproto/superadmin_service"
	"go_user_service/pkg"
	"go_user_service/pkg/hash"
	"go_user_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type superadminRepo struct {
	db *pgxpool.Pool
}

func NewSuperadminRepo(db *pgxpool.Pool) storage.SuperadminRepoI {
	return &superadminRepo{
		db: db,
	}
}

func generateSuperadminLogin(db *pgxpool.Pool, ctx context.Context) (string, error) {
	var nextVal int
	err := db.QueryRow(ctx, "SELECT nextval('superadmin_external_id_seq')").Scan(&nextVal)
	if err != nil {
		return "", err
	}
	userLogin := "SA" + fmt.Sprintf("%05d", nextVal)
	return userLogin, nil
}

func (c *superadminRepo) Create(ctx context.Context, req *tc.CreateSuperadmin) (*tc.GetSuperadmin, error) {

	id := uuid.NewString()
	pasword, err := hash.HashPassword(req.UserPassword)
	if err != nil {
		log.Println("error while hashing password", err)
	}

	userLogin, err := generateSuperadminLogin(c.db, ctx)
	if err != nil {
		log.Println("error while generating login", err)
	}
	comtag, err := c.db.Exec(ctx, `
		INSERT INTO superadmins (
			id,
			user_login,
			birthday,
			gender,
			fullname,
			email,
			phone,
			user_password,
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8
		)`,
		id,
		userLogin,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		pasword)
	if err != nil {
		log.Println("error while creating superadmin", comtag)
		return nil, err
	}

	superadmin, err := c.GetById(ctx, &tc.SuperadminPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting superadmin by id")
		return nil, err
	}
	return superadmin, nil
}

func (c *superadminRepo) Update(ctx context.Context, req *tc.UpdateSuperadmin) (*tc.GetSuperadmin, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE superadmins SET
		birthday = $1,
		gender = $2,
		fullname = $3,
		email = $4,
		phone = $5,
		updated_at = NOW()
		WHERE id = $6
		`,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		req.Id)
	if err != nil {
		log.Println("error while updating superadmin")
		return nil, err
	}

	superadmin, err := c.GetById(ctx, &tc.SuperadminPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting superadmin by id")
		return nil, err
	}
	return superadmin, nil
}

func (c *superadminRepo) GetById(ctx context.Context, id *tc.SuperadminPrimaryKey) (*tc.GetSuperadmin, error) {
	var (
		superadmin tc.GetSuperadmin
		created_at sql.NullString
		updated_at sql.NullString
	)

	query := `SELECT
				id,
				birthday,
				gender,
				fullname,
				email,
				phone,
				created_at,
				updated_at
			FROM superadmins
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&superadmin.Id,
		&superadmin.Birthday,
		&superadmin.Gender,
		&superadmin.Fullname,
		&superadmin.Email,
		&superadmin.Phone,
		&created_at,
		&updated_at); err != nil {
		return &superadmin, err
	}

	superadmin.CreatedAt = pkg.NullStringToString(created_at)
	superadmin.UpdatedAt = pkg.NullStringToString(updated_at)

	return &superadmin, nil
}

func (c *superadminRepo) Delete(ctx context.Context, id *tc.SuperadminPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE superadmins SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
