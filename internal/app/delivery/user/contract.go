package user

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type Usecase interface {
	UserCreate(ctx context.Context, user *domain.User) ([]*domain.User, error)
	UserDetails(ctx context.Context, user *domain.User) (*domain.User, error)
	UserUpdate(ctx context.Context, user *domain.User) (*domain.User, error)
}
