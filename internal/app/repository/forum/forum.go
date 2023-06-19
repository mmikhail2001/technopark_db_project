package forum

import (
	"context"
	"database/sql"
	"log"

	"github.com/pkg/errors"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg/sqltools"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) ForumCreate(ctx context.Context, forum *domain.Forum) (*domain.Forum, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		err := conn.QueryRowContext(ctx, queryInsertForum, forum.Slug, forum.Title, forum.Author).Scan(&forum.Title, &forum.Author, &forum.Slug, &forum.CountPosts, &forum.CountThreads)

		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchForumExist
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v], special error: [%s]", queryInsertForum, *forum, err)
		}
		return nil
	})
	return forum, err
}

func (repo *Repository) GetForumBySlug(ctx context.Context, forum *domain.Forum) (*domain.Forum, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		err := conn.QueryRowContext(ctx, querySelectForum, forum.Slug).Scan(&forum.Slug, &forum.Title, &forum.Author, &forum.CountPosts, &forum.CountThreads)
		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchForumNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%s], special error: [%s]", querySelectForum, forum.Slug, err)
		}
		return nil
	})
	return forum, err
}
