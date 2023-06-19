package post

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type ForumRepository interface {
	GetForumBySlug(ctx context.Context, forum *domain.Forum) (*domain.Forum, error)
}

type UserRepository interface {
	GetUserByNickname(ctx context.Context, user *domain.User) (*domain.User, error)
}

type ThreadRepository interface {
	GetThreadByID(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
	GetThreadBySlug(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
}

type PostRepository interface {
	GetPostByID(ctx context.Context, post *domain.Post) (*domain.Post, error)
	UpdateMessage(ctx context.Context, post *domain.Post) (*domain.Post, error)
	PostsCreate(ctx context.Context, posts []*domain.Post, thread *domain.Thread) ([]domain.Post, error)
}
