package postgres

import (
	"context"
	br "go_user_service/genproto/event_registrate_service"
	"go_user_service/storage"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type eventRegistrateRepo struct {
	db *pgxpool.Pool
}

func NewEventRegistrateRepo(db *pgxpool.Pool) storage.EventRegistrateRepoI {
	return &eventRegistrateRepo{
		db: db,
	}
}

func (c *eventRegistrateRepo) Create(ctx context.Context, req *br.CreateEventRegistrate) (*br.GetEventRegistrate, error) {
	id := uuid.NewString()

	comtag, err := c.db.Exec(ctx, `
		INSERT INTO event_registrate (
			id,
			event_id,
			student_id
		) VALUES ($1,$2,$3
		)`,
		id,
		req.EventId,
		req.StudentId,
	)
	if err != nil {
		log.Println("error while creating event", comtag)
		return nil, err
	}

	event, err := c.GetById(ctx, &br.EventRegistratePrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting event by id")
		return nil, err
	}
	return event, nil
}

func (c *eventRegistrateRepo) Update(ctx context.Context, req *br.UpdateEventRegistrate) (*br.GetEventRegistrate, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE event_registrate SET
		event_id = $1,
		student_id = $2,
		updated_at = NOW()
		WHERE id = $3
		`,
		req.EventId,
		req.StudentId,
		req.Id,
	)
	if err != nil {
		log.Println("error while updating event_registrate")
		return nil, err
	}

	event_registrate, err := c.GetById(ctx, &br.EventRegistratePrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting event_registrate by id")
		return nil, err
	}
	return event_registrate, nil
}

func (c *eventRegistrateRepo) GetById(ctx context.Context, id *br.EventRegistratePrimaryKey) (*br.GetEventRegistrate, error) {
	var (
		event_registrate br.GetEventRegistrate
	)

	query := `SELECT
				id,
				event_id,
				student_id,
				created_at,
				updated_at
			FROM event_registrate
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&event_registrate.Id,
		&event_registrate.EventId,
		&event_registrate.StudentId,
	); err != nil {
		return &event_registrate, err
	}

	return &event_registrate, nil
}

func (c *eventRegistrateRepo) Delete(ctx context.Context, id *br.EventRegistratePrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE event_registrate SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}
