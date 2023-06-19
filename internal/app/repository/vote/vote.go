package vote

import (
	"context"
	"database/sql"
	"log"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg/sqltools"
	"github.com/pkg/errors"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) AddVoteToThread(ctx context.Context, vote *domain.Vote) error {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		_, err := conn.ExecContext(ctx, queryInsertVote, vote.Author, vote.Thread, vote.Vote)

		if err != nil {
			log.Println(err)
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v], special error: [%s]", queryInsertVote, *vote, err)
		}
		return nil
	})
	return err
}
