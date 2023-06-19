package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

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

func (repo *Repository) GetUserByNickname(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		err := conn.QueryRowContext(ctx, querySelectUserByNickname, user.Nickname).Scan(&user.Nickname, &user.Fullname, &user.About, &user.Email)
		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchUserNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%s], special error: [%s]", querySelectUserByNickname, user.Nickname, err)
		}
		return nil
	})
	return user, err
}

func (repo *Repository) GetUserByEmail(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		err := conn.QueryRowContext(ctx, querySelectUserByEmail, user.Email).Scan(&user.Nickname, &user.Fullname, &user.About, &user.Email)
		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchUserNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%s], special error: [%s]", querySelectUserByNickname, user.Nickname, err)
		}
		return nil
	})
	return user, err
}

func (repo *Repository) GetUsersByNicknameOrEmail(ctx context.Context, user *domain.User) ([]*domain.User, error) {
	res := make([]*domain.User, 0)
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		rows, err := conn.QueryContext(ctx, querySelectUserByNicknameOrEmail, user.Nickname, user.Email)
		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchUserNotFound
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%s], special error: [%s]", querySelectUserByNickname, user.Nickname, err)
		}
		defer rows.Close()
		for rows.Next() {
			user := &domain.User{}
			err := rows.Scan(&user.Nickname, &user.Fullname, &user.About, &user.Email)
			if err != nil {
				return err
			}
			res = append(res, user)
		}

		if len(res) == 0 {
			return pkg.ErrSuchUserNotFound
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return res, err
}

func (repo *Repository) GetUsersByForumSlug(ctx context.Context, forum *domain.Forum, params *domain.GetUsersParams) ([]*domain.User, error) {
	query := querySelectUsersByForumSlug_begin

	// TODO: sql injections
	switch {
	case params.Desc && params.Since != "":
		query += fmt.Sprintf(" AND uf.nickname < '%s'", params.Since)
	case params.Since != "":
		query += fmt.Sprintf(" AND uf.nickname > '%s'", params.Since)
	}

	query += " ORDER BY uf.nickname "

	if params.Desc {
		query += "DESC"
	}

	query += fmt.Sprintf(" LIMIT %d", params.Limit)

	res := make([]*domain.User, 0, 5)

	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		rows, err := conn.QueryContext(ctx, query, forum.Slug)
		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrUsersByForumSlugNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				query, forum.Slug, err)
		}
		defer rows.Close()

		for rows.Next() {
			user := &domain.User{}

			err = rows.Scan(
				&user.Nickname,
				&user.Fullname,
				&user.About,
				&user.Email)
			if err != nil {
				return err
			}

			res = append(res, user)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *Repository) UserCreate(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		err := conn.QueryRowContext(ctx, queryInsertUser, user.Nickname, user.Fullname, user.About, user.Email).Scan(&user.Nickname, &user.Fullname, &user.About, &user.Email)

		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchUserExist
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v], special error: [%s]", queryInsertUser, *user, err)
		}
		return nil
	})
	return user, err
}

func (repo *Repository) UserUpdate(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := sqltools.RunQuery(ctx, repo.DB, func(ctx context.Context, conn *sql.Conn) error {
		err := conn.QueryRowContext(ctx, queryUpdateUser, user.Fullname, user.About, user.Email, user.Nickname).Scan(&user.Fullname, &user.About, &user.Email, &user.Nickname)

		if err != nil {
			log.Println(err)
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchUserNotFound
			}
			if strings.Contains(err.Error(), "users_email_key") {
				return pkg.ErrSuchUserExist
			}
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query: [%s], values: [%+v], special error: [%s]", queryInsertUser, *user, err)
		}
		return nil
	})
	return user, err
}
