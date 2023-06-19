package post

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type Usecase interface {
	PostDetails(ctx context.Context, post *domain.Post, params *domain.PostDetailsParams) (*domain.PostDetails, error)
	PostUpdateMessage(ctx context.Context, post *domain.Post) (*domain.Post, error)
	PostsCreate(ctx context.Context, posts []*domain.Post, thread *domain.Thread) ([]domain.Post, error)
}
