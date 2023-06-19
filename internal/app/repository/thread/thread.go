package thread

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

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

func setSlugByNullString(thread *domain.Thread, slug sql.NullString) {
	if slug.Valid == false {
		thread.Slug = ""
		return
	}
	thread.Slug = slug.String
}

func (repo *Repository) ThreadCreate(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	if thread.Created == "" {
		thread.Created = time.Now().Format(time.RFC3339)
	}
	var ThreadSlugOutput sql.NullString
	var ThreadSlugInput sql.NullString
	if thread.Slug == "" {
		ThreadSlugInput.Valid = false
	} else {
		ThreadSlugInput.Valid = true
		ThreadSlugInput.String = thread.Slug
	}
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		err := conn.QueryRowContext(ctx, queryInsertThread,
			thread.Title,
			thread.Message,
			thread.Author,
			thread.Forum,
			ThreadSlugInput,
			thread.Created).Scan(
			&thread.ID,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.SumVotes,
			&ThreadSlugOutput,
			&thread.Created)

		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchThreadExist
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v], special error: [%s]", queryInsertThread, *thread, err)
		}
		setSlugByNullString(thread, ThreadSlugOutput)
		return nil
	})
	return thread, err
}

func (repo *Repository) GetThreadBySlug(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		var ThreadSlugOutput sql.NullString
		err := conn.QueryRowContext(ctx, querySelectThreadBySlug, thread.Slug).Scan(&thread.ID, &thread.Title, &thread.Author, &thread.Forum, &thread.Message, &thread.SumVotes, &ThreadSlugOutput, &thread.Created)

		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v], special error: [%s]", querySelectThreadBySlug, *thread, err)
		}
		setSlugByNullString(thread, ThreadSlugOutput)
		return nil
	})
	return thread, err
}

func (repo *Repository) GetThreadByID(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		var ThreadSlugOutput sql.NullString
		err := conn.QueryRowContext(ctx, querySelectThreadByID, thread.ID).Scan(&thread.ID, &thread.Title, &thread.Author, &thread.Forum, &thread.Message, &thread.SumVotes, &ThreadSlugOutput, &thread.Created)

		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v], special error: [%s]", querySelectThreadByID, *thread, err)
		}
		setSlugByNullString(thread, ThreadSlugOutput)
		return nil
	})
	return thread, err
}

func (repo *Repository) GetThreadsByForumSlug(ctx context.Context, forum *domain.Forum, params *domain.GetThreadsParams) ([]*domain.Thread, error) {
	query := querySelectThreadsByForumSlug_begin

	orderBy := "ORDER BY t.created "
	if params.Desc {
		orderBy += "DESC"
	}

	if params.Limit > 0 {
		orderBy += fmt.Sprintf(" LIMIT %d", params.Limit)
	}

	var querySince string
	switch {
	case params.Since != "" && params.Desc:
		querySince = " AND t.created <= $2 "
	case params.Since != "" && !params.Desc:
		querySince = " AND t.created >= $2 "
	}

	var values []interface{}

	if params.Since != "" {
		query += querySince + orderBy

		values = []interface{}{forum.Slug, params.Since}
	} else {
		query += orderBy

		values = []interface{}{forum.Slug}
	}

	res := make([]*domain.Thread, 0)

	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		rows, err := conn.QueryContext(ctx, query, values...)
		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v, %+v], special error: [%s]", query, values, err)
		}
		defer rows.Close()

		for rows.Next() {
			thread := &domain.Thread{}
			var ThreadSlugOutput sql.NullString
			err = rows.Scan(
				&thread.ID,
				&thread.Title,
				&thread.Author,
				&thread.Forum,
				&thread.Message,
				&thread.SumVotes,
				&ThreadSlugOutput,
				&thread.Created)
			if err != nil {
				return err
			}

			setSlugByNullString(thread, ThreadSlugOutput)

			res = append(res, thread)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo *Repository) ThreadUpdate(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		var ThreadSlugOutput sql.NullString
		err := conn.QueryRowContext(ctx, queryUpdateThread, thread.Title, thread.Message, thread.ID).Scan(&thread.ID, &thread.Title, &thread.Author, &thread.Forum, &thread.Message, &thread.SumVotes, &ThreadSlugOutput, &thread.Created)

		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+#v], special error: [%s]", queryUpdateThread, *thread, err)
		}
		setSlugByNullString(thread, ThreadSlugOutput)
		return nil
	})
	return thread, err
}
