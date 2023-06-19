package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg/sqltools"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) GetCountTuples(ctx context.Context) (*domain.StatusService, error) {
	status := &domain.StatusService{}
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		errUsers := conn.QueryRowContext(ctx, queryCountUsers).Scan(&status.User)
		errRest := conn.QueryRowContext(ctx, queryCountForumsThreadsPosts).Scan(&status.Forum, &status.Thread, &status.Post)
		if errUsers != nil || errRest != nil {
			log.Println(errUsers, errRest)
			return fmt.Errorf("%w %w", errUsers, errRest)
		}
		return nil
	})
	return status, err
}

func (repo *Repository) DeleteAllEntities(ctx context.Context) error {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		_, err := conn.ExecContext(ctx, queryTrancate)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	return err
}
