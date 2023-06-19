package vote

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type Usecase struct {
	voteRepo   VoteRepository
	userRepo   UserRepository
	threadRepo ThreadRepository
}

func NewUsecase(voteRepo VoteRepository, userRepo UserRepository, threadRepo ThreadRepository) *Usecase {
	return &Usecase{
		voteRepo:   voteRepo,
		userRepo:   userRepo,
		threadRepo: threadRepo,
	}
}

func (u *Usecase) Vote(ctx context.Context, vote *domain.Vote, thread *domain.Thread) (*domain.Thread, error) {
	var err error
	if thread.Slug != "" {
		thread, err = u.threadRepo.GetThreadBySlug(ctx, thread)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := u.threadRepo.GetThreadByID(ctx, thread)
		if err != nil {
			return nil, err
		}
	}
	vote.Thread = thread.ID
	user := &domain.User{
		Nickname: vote.Author,
	}
	_, err = u.userRepo.GetUserByNickname(ctx, user)
	if err != nil {
		return thread, err
	}
	err = u.voteRepo.AddVoteToThread(ctx, vote)
	if err != nil {
		return thread, err
	}
	thread, err = u.threadRepo.GetThreadByID(ctx, thread)
	if err != nil {
		return nil, err
	}
	return thread, nil
}
