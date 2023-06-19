package forum

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type Usecase interface {
	ThreadCreate(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
	ThreadDetails(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
	ThreadUpdate(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
	ThreadPosts(ctx context.Context, thread *domain.Thread, params *domain.GetPostsParams) ([]*domain.Post, error)
}
