package user

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByNickname(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUsersByNicknameOrEmail(ctx context.Context, user *domain.User) ([]*domain.User, error)
	UserCreate(ctx context.Context, user *domain.User) (*domain.User, error)
	UserUpdate(ctx context.Context, user *domain.User) (*domain.User, error)
}
