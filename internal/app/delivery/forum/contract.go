package forum

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type Usecase interface {
	ForumCreate(ctx context.Context, forum *domain.Forum) (*domain.Forum, error)
	ForumDetails(ctx context.Context, forum *domain.Forum) (*domain.Forum, error)
	ForumUsers(ctx context.Context, forum *domain.Forum, params *domain.GetUsersParams) ([]*domain.User, error)
	ForumThreads(ctx context.Context, forum *domain.Forum, params *domain.GetThreadsParams) ([]*domain.Thread, error)
}
