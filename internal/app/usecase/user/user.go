package user

import (
	"context"
	"errors"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg"
)

type Usecase struct {
	userRepo UserRepository
}

func NewUsecase(userRepo UserRepository) *Usecase {
	return &Usecase{
		userRepo: userRepo,
	}
}

func (u *Usecase) UserCreate(ctx context.Context, user *domain.User) ([]*domain.User, error) {
	res, err := u.userRepo.GetUsersByNicknameOrEmail(ctx, user)
	if err == nil {
		return res, pkg.ErrSuchUserExist
	}

	user, err = u.userRepo.UserCreate(ctx, user)
	if err != nil {
		return nil, err
	}

	resOne := []*domain.User{user}

	return resOne, nil
}

func (u *Usecase) UserDetails(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.userRepo.GetUserByNickname(ctx, user)
}

func (u *Usecase) UserUpdate(ctx context.Context, user *domain.User) (*domain.User, error) {
	_, err := u.userRepo.GetUserByEmail(ctx, user)
	if err == nil {
		return nil, pkg.ErrSuchUserExist
	} else if err != nil && !errors.Is(err, pkg.ErrSuchUserNotFound) {
		return user, err
	}
	return u.userRepo.UserUpdate(ctx, user)
}
