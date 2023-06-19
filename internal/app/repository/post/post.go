package post

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
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

func (repo *Repository) GetPostByID(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		err := conn.QueryRowContext(ctx, querySelectPostByID, post.ID).Scan(&post.ID, &post.Parent, &post.Author, &post.Message, &post.IsEdited, &post.Forum, &post.Thread, &post.Created)

		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v], special error: [%s]", querySelectPostByID, *post, err)
		}
		return nil
	})
	return post, err
}

func (repo *Repository) UpdateMessage(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		err := conn.QueryRowContext(ctx, queryUpdateMessage, post.Message, post.ID).Scan(&post.ID, &post.Parent, &post.Author, &post.Message, &post.IsEdited, &post.Forum, &post.Thread, &post.Created)

		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v], special error: [%s]", querySelectPostByID, *post, err)
		}
		return nil
	})
	return post, err
}

func (repo *Repository) PostsCreate(ctx context.Context, posts []*domain.Post, thread *domain.Thread) ([]domain.Post, error) {

	query := queryInsertPosts_begin

	countAttributes := strings.Count(query, ",") + 1

	pos := 0

	countInserts := len(posts)

	values := make([]interface{}, countInserts*countAttributes)

	insertTimeString := time.Now().Format(time.RFC3339)

	for i := 0; i < len(posts); i++ {
		values[pos] = posts[i].Parent
		pos++
		values[pos] = posts[i].Author
		pos++
		values[pos] = posts[i].Message
		pos++
		values[pos] = thread.Forum
		pos++
		values[pos] = thread.ID
		pos++
		values[pos] = insertTimeString
		pos++
	}

	query = sqltools.CreateFullQuery(query, countInserts, countAttributes)

	query += " RETURNING id, is_edited;"

	rows, err := repo.DB.QueryContext(ctx, query, values...)
	defer rows.Close()
	if err != nil || rows.Err() != nil {
		log.Println(err)

		return nil, fmt.Errorf("InsertBatch: [%w] when inserting row into [%s] table \n [%+v]", err, query, values)
	}
	res := make([]domain.Post, len(posts))

	i := 0
	for rows.Next() {
		err = rows.Scan(&res[i].ID, &res[i].IsEdited)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		res[i].Created = insertTimeString
		res[i].Parent = posts[i].Parent
		res[i].Author = posts[i].Author
		res[i].Message = posts[i].Message
		res[i].Forum = thread.Forum
		res[i].Thread = thread.ID

		i++
	}

	// костыль, но QueryContext не возвращает ошибку
	if i == 0 {
		return nil, pkg.ErrPostParentNotFound
	}

	return res, nil
}
func (repo *Repository) GetPostsWithSort_Flat(ctx context.Context, thread *domain.Thread, params *domain.GetPostsParams) ([]*domain.Post, error) {
	query := querySelectPostsWithSort_Flat_Tree_Begin
	var values []interface{}

	switch {
	case params.Since != -1 && params.Desc:
		query += " AND id < $2"
	case params.Since != -1 && !params.Desc:
		query += " AND id > $2"
	case params.Since != -1:
		query += " AND id > $2"
	}

	switch {
	case params.Desc:
		query += " ORDER BY created DESC, id DESC"
	default:
		query += " ORDER BY created, id"
	}

	// TODO: sql injection
	query += fmt.Sprintf(" LIMIT %d", params.Limit)

	if params.Since == -1 {
		values = []interface{}{thread.ID}
	} else {
		values = []interface{}{thread.ID, params.Since}
	}

	res := make([]*domain.Post, 0)

	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		rows, err := conn.QueryContext(ctx, query, values...)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%+v]. Special error: [%s]",
				query, values, err)
		}
		defer rows.Close()

		for rows.Next() {
			post := &domain.Post{}

			timeTmp := time.Time{}

			err = rows.Scan(
				&post.ID,
				&post.Parent,
				&post.Author,
				&post.Message,
				&post.IsEdited,
				&post.Forum,
				&post.Thread,
				&timeTmp)
			if err != nil {
				return err
			}

			post.Created = timeTmp.Format(time.RFC3339)

			res = append(res, post)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *Repository) GetPostsWithSort_Tree(ctx context.Context, thread *domain.Thread, params *domain.GetPostsParams) ([]*domain.Post, error) {
	query := querySelectPostsWithSort_Flat_Tree_Begin

	switch {
	case params.Since != -1 && params.Desc:
		query += " AND path < "
	case params.Since != -1 && !params.Desc:
		query += " AND path > "
	case params.Since != -1:
		query += " AND path > "
	}

	if params.Since != -1 {
		query += fmt.Sprintf(` (SELECT path FROM posts WHERE id = %d) `, params.Since)
	}

	switch {
	case params.Desc:
		query += " ORDER BY path DESC "
	default:
		query += " ORDER BY path "
	}

	query += fmt.Sprintf(" LIMIT %d", params.Limit)

	res := make([]*domain.Post, 0)

	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		rows, err := conn.QueryContext(ctx, query, thread.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%+v]. Special error: [%s]",
				query, thread.ID, err)
		}
		defer rows.Close()

		for rows.Next() {
			post := &domain.Post{}

			timeTmp := time.Time{}

			err = rows.Scan(
				&post.ID,
				&post.Parent,
				&post.Author,
				&post.Message,
				&post.IsEdited,
				&post.Forum,
				&post.Thread,
				&timeTmp)
			if err != nil {
				return err
			}

			post.Created = timeTmp.Format(time.RFC3339)

			res = append(res, post)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *Repository) GetPostsWithSort_ParentTree(ctx context.Context, thread *domain.Thread, params *domain.GetPostsParams) ([]*domain.Post, error) {
	var query string
	var values []interface{}

	if params.Since == -1 {
		if params.Desc {
			query = `
			with root_posts as (
				select id
				  from posts
					where parent = 0 
					  and thread = $1
					  order by id desc
						limit $2
			  ) select id, parent, author, message, is_edited, forum, thread, created 
				  from posts
					where path[1] in (select id from root_posts)
					  order by path[1] desc, path;`
		} else {
			query = `
			with root_posts as (
				select id
				  from posts
					where parent = 0 
					  and thread = $1
					  order by id
						limit $2
			  ) select id, parent, author, message, is_edited, forum, thread, created 
				  from posts
					where path[1] in (select id from root_posts)
					  order by path;`
		}

		values = []interface{}{thread.ID, params.Limit}
	} else {
		if params.Desc {
			query = `
			select id, parent, author, message, is_edited, forum, thread, created 
				from posts
					where path[1] in 
					(select id from posts 
						where thread = $1 and parent = 0 and path[1] <
									(select path[1] from posts where id = $2) 
					order by id desc limit $3) 
						order by path[1] desc, path, id;`
		} else {
			query = `
			select id, parent, author, message, is_edited, forum, thread, created 
				from posts
				where path[1] in 
					(select id from posts 
					where thread = $1 and parent = 0 and path[1] >
								(select path[1] from posts where id = $2) 
					order by id limit $3) 
					order by path, id;`

		}

		values = []interface{}{thread.ID, params.Since, params.Limit}
	}

	res := make([]*domain.Post, 0)

	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		rows, err := conn.QueryContext(ctx, query, values...)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%+v]. Special error: [%s]",
				query, values, err)
		}
		defer rows.Close()

		for rows.Next() {
			post := &domain.Post{}

			timeTmp := time.Time{}

			err = rows.Scan(
				&post.ID,
				&post.Parent,
				&post.Author,
				&post.Message,
				&post.IsEdited,
				&post.Forum,
				&post.Thread,
				&timeTmp)
			if err != nil {
				return err
			}

			post.Created = timeTmp.Format(time.RFC3339)

			res = append(res, post)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil

}
