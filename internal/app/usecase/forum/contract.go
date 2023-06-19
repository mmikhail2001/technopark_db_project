package forum

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type ForumRepository interface {
	ForumCreate(ctx context.Context, forum *domain.Forum) (*domain.Forum, error)
	GetForumBySlug(ctx context.Context, forum *domain.Forum) (*domain.Forum, error)
}

type UserRepository interface {
	GetUserByNickname(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUsersByForumSlug(ctx context.Context, forum *domain.Forum, params *domain.GetUsersParams) ([]*domain.User, error)
}

type ThreadRepository interface {
	GetThreadsByForumSlug(ctx context.Context, forum *domain.Forum, params *domain.GetThreadsParams) ([]*domain.Thread, error)
}
