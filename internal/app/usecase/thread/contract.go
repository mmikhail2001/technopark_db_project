package thread

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
	ThreadCreate(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
	ThreadUpdate(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
	GetThreadBySlug(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
	GetThreadByID(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
}

type PostRepository interface {
	GetPostsWithSort_Flat(ctx context.Context, thread *domain.Thread, params *domain.GetPostsParams) ([]*domain.Post, error)
	GetPostsWithSort_Tree(ctx context.Context, thread *domain.Thread, params *domain.GetPostsParams) ([]*domain.Post, error)
	GetPostsWithSort_ParentTree(ctx context.Context, thread *domain.Thread, params *domain.GetPostsParams) ([]*domain.Post, error)
}
