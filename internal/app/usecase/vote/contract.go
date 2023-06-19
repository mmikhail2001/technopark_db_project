package vote

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type VoteRepository interface {
	AddVoteToThread(ctx context.Context, vote *domain.Vote) error
}

type UserRepository interface {
	GetUserByNickname(ctx context.Context, user *domain.User) (*domain.User, error)
}

type ThreadRepository interface {
	GetThreadBySlug(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
	GetThreadByID(ctx context.Context, thread *domain.Thread) (*domain.Thread, error)
}
