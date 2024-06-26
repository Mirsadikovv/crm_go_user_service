package postgres

import (
	"context"
	"database/sql"
	"fmt"
	br "go_user_service/genproto/group_service"
	"go_user_service/pkg"
	"go_user_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type groupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) storage.GroupRepoI {
	return &groupRepo{
		db: db,
	}
}

func (c *groupRepo) Create(ctx context.Context, req *br.CreateGroup) (*br.GetGroup, error) {
	id := uuid.NewString()

	comtag, err := c.db.Exec(ctx, `
		INSERT INTO groups (
			id,
			branch_id,
			teacher_id,
			support_teacher_id,
			group_name,
			started_at,
			finished_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7
		)`,
		id,
		req.BranchId,
		req.TeacherId,
		req.SupportTeacherId,
		req.GroupName,
		req.StartedAt,
		req.FinishedAt,
	)
	if err != nil {
		log.Println("error while creating group", comtag)
		return nil, err
	}

	group, err := c.GetById(ctx, &br.GroupPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting group by id")
		return nil, err
	}
	return group, nil
}

func (c *groupRepo) Update(ctx context.Context, req *br.UpdateGroup) (*br.GetGroup, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE groups SET
		branch_id = $1,
		teacher_id = $2,
		support_teacher_id = $3,
		group_name = $4,
		started_at = $5,
		finished_at = $6,
		updated_at = NOW()
		WHERE id = $7
		`,
		req.Id,
		req.BranchId,
		req.TeacherId,
		req.SupportTeacherId,
		req.GroupName,
		req.StartedAt,
		req.FinishedAt,
	)
	if err != nil {
		log.Println("error while updating group")
		return nil, err
	}

	group, err := c.GetById(ctx, &br.GroupPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting group by id")
		return nil, err
	}
	return group, nil
}

func (c *groupRepo) GetAll(ctx context.Context, req *br.GetListGroupRequest) (*br.GetListGroupResponse, error) {
	groups := br.GetListGroupResponse{}
	var (
		created_at  sql.NullString
		updated_at  sql.NullString
		started_at  sql.NullString
		finished_at sql.NullString
	)
	filter_by_name := ""
	offest := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter_by_name = fmt.Sprintf(`AND group_name ILIKE '%%%v%%'`, req.Search)
	}
	query := `SELECT
				id,
				branch_id,
				teacher_id,
				support_teacher_id,
				group_name,
				started_at,
				finished_at,
				created_at,
				updated_at
			FROM groups
			WHERE TRUE AND deleted_at is null ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting all groups")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			group br.GetGroup
		)
		if err = rows.Scan(
			&group.Id,
			&group.BranchId,
			&group.TeacherId,
			&group.SupportTeacherId,
			&group.GroupName,
			&started_at,
			&finished_at,
			&created_at,
			&updated_at,
		); err != nil {
			return &groups, err
		}
		group.StartedAt = pkg.NullStringToString(started_at)
		group.FinishedAt = pkg.NullStringToString(finished_at)
		group.CreatedAt = pkg.NullStringToString(created_at)
		group.UpdatedAt = pkg.NullStringToString(updated_at)

		groups.Groups = append(groups.Groups, &group)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from groups WHERE TRUE AND deleted_at is null `+filter_by_name+``).Scan(&groups.Count)
	if err != nil {
		return &groups, err
	}

	return &groups, nil
}

func (c *groupRepo) GetById(ctx context.Context, id *br.GroupPrimaryKey) (*br.GetGroup, error) {
	var (
		group       br.GetGroup
		created_at  sql.NullString
		updated_at  sql.NullString
		started_at  sql.NullString
		finished_at sql.NullString
	)

	query := `SELECT
				id,
				branch_id,
				teacher_id,
				support_teacher_id,
				group_name,
				started_at,
				finished_at,
				created_at,
				updated_at
			FROM groups
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&group.Id,
		&group.BranchId,
		&group.TeacherId,
		&group.SupportTeacherId,
		&group.GroupName,
		&started_at,
		&finished_at,
		&created_at,
		&updated_at); err != nil {
		return &group, err
	}
	group.StartedAt = pkg.NullStringToString(started_at)
	group.FinishedAt = pkg.NullStringToString(finished_at)
	group.CreatedAt = pkg.NullStringToString(created_at)
	group.UpdatedAt = pkg.NullStringToString(updated_at)

	return &group, nil
}

func (c *groupRepo) Delete(ctx context.Context, id *br.GroupPrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE groups SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
