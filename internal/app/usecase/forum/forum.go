package forum

import (
	"context"
	"errors"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg"
)

type Usecase struct {
	forumRepo  ForumRepository
	userRepo   UserRepository
	threadRepo ThreadRepository
}

func NewUsecase(forumRepo ForumRepository, userRepo UserRepository, threadRepo ThreadRepository) *Usecase {
	return &Usecase{
		forumRepo:  forumRepo,
		userRepo:   userRepo,
		threadRepo: threadRepo,
	}
}

func (u *Usecase) ForumCreate(ctx context.Context, forum *domain.Forum) (*domain.Forum, error) {
	user := &domain.User{
		Nickname: forum.Author,
	}
	user, err := u.userRepo.GetUserByNickname(ctx, user)
	if err != nil {
		return forum, err
	}
	forum.Author = user.Nickname
	forum, err = u.forumRepo.ForumCreate(ctx, forum)
	if err != nil {
		if !errors.Is(err, pkg.ErrSuchForumExist) {
			return forum, err
		}
		forum, err = u.forumRepo.GetForumBySlug(ctx, forum)
		if err != nil {
			return forum, err
		}
		return forum, pkg.ErrSuchForumExist
	}
	return forum, nil
}

func (u *Usecase) ForumDetails(ctx context.Context, forum *domain.Forum) (*domain.Forum, error) {
	forum, err := u.forumRepo.GetForumBySlug(ctx, forum)
	if err != nil {
		return forum, err
	}
	return forum, nil
}

func (u *Usecase) ForumUsers(ctx context.Context, forum *domain.Forum, params *domain.GetUsersParams) ([]*domain.User, error) {
	forum, err := u.forumRepo.GetForumBySlug(ctx, forum)
	if err != nil {
		return nil, err
	}
	users, err := u.userRepo.GetUsersByForumSlug(ctx, forum, params)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *Usecase) ForumThreads(ctx context.Context, forum *domain.Forum, params *domain.GetThreadsParams) ([]*domain.Thread, error) {
	forum, err := u.forumRepo.GetForumBySlug(ctx, forum)
	if err != nil {
		return nil, err
	}
	threads, err := u.threadRepo.GetThreadsByForumSlug(ctx, forum, params)
	if err != nil {
		return nil, err
	}
	return threads, nil
}
