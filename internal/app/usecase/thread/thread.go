package thread

import (
	"context"
	"errors"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg"
)

type Usecase struct {
	threadRepo ThreadRepository
	userRepo   UserRepository
	forumRepo  ForumRepository
	postRepo   PostRepository
}

func NewUsecase(threadRepo ThreadRepository, forumRepo ForumRepository, userRepo UserRepository, postRepo PostRepository) *Usecase {
	return &Usecase{
		forumRepo:  forumRepo,
		userRepo:   userRepo,
		threadRepo: threadRepo,
		postRepo:   postRepo,
	}
}

func (u *Usecase) ThreadCreate(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	user := &domain.User{
		Nickname: thread.Author,
	}
	_, err := u.userRepo.GetUserByNickname(ctx, user)
	if err != nil {
		return thread, err
	}
	forum := &domain.Forum{
		Slug: thread.Forum,
	}
	forum, err = u.forumRepo.GetForumBySlug(ctx, forum)
	if err != nil {
		return thread, err
	}
	thread.Forum = forum.Slug
	thread, err = u.threadRepo.ThreadCreate(ctx, thread)
	if err != nil {
		if !errors.Is(err, pkg.ErrSuchThreadExist) {
			return thread, err
		}
		thread, err = u.threadRepo.GetThreadBySlug(ctx, thread)
		if err != nil {
			return thread, err
		}
		return thread, pkg.ErrSuchThreadExist
	}
	return thread, nil
}

func (u *Usecase) ThreadDetails(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	if thread.Slug != "" {
		thread, err := u.threadRepo.GetThreadBySlug(ctx, thread)
		if err != nil {
			return thread, err
		}
	} else {
		thread, err := u.threadRepo.GetThreadByID(ctx, thread)
		if err != nil {
			return thread, err
		}
	}
	return thread, nil
}

func (u *Usecase) ThreadUpdate(ctx context.Context, thread *domain.Thread) (*domain.Thread, error) {
	if thread.Slug != "" {
		var err error
		threadTmp := &domain.Thread{}
		threadTmp.Slug = thread.Slug
		threadTmp, err = u.threadRepo.GetThreadBySlug(ctx, threadTmp)
		if err != nil {
			return thread, err
		}
		thread.ID = threadTmp.ID
	}
	return u.threadRepo.ThreadUpdate(ctx, thread)
}

func (u *Usecase) ThreadPosts(ctx context.Context, thread *domain.Thread, params *domain.GetPostsParams) ([]*domain.Post, error) {
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
	switch params.Sort {
	case "flat":
		return u.postRepo.GetPostsWithSort_Flat(ctx, thread, params)
	case "tree":
		return u.postRepo.GetPostsWithSort_Tree(ctx, thread, params)
	case "parent_tree":
		return u.postRepo.GetPostsWithSort_ParentTree(ctx, thread, params)
	}
	return nil, pkg.ErrNoSuchRuleSortPosts
}
