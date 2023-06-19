package post

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg"
)

type Usecase struct {
	postRepo   PostRepository
	threadRepo ThreadRepository
	userRepo   UserRepository
	forumRepo  ForumRepository
}

func NewUsecase(postRepo PostRepository, threadRepo ThreadRepository, forumRepo ForumRepository, userRepo UserRepository) *Usecase {
	return &Usecase{
		postRepo:   postRepo,
		forumRepo:  forumRepo,
		userRepo:   userRepo,
		threadRepo: threadRepo,
	}
}

func (u *Usecase) PostDetails(ctx context.Context, post *domain.Post, params *domain.PostDetailsParams) (*domain.PostDetails, error) {
	postDetails := &domain.PostDetails{}
	postDetails.Post.ID = post.ID

	post, err := u.postRepo.GetPostByID(ctx, post)
	if err != nil {
		return postDetails, err
	}
	postDetails.Post = *post

	for _, relate := range params.Related {
		switch relate {
		case "user":
			user := &domain.User{
				Nickname: post.Author,
			}
			user, err := u.userRepo.GetUserByNickname(ctx, user)
			if err != nil {
				return postDetails, err
			}
			postDetails.Author = *user
		case "forum":
			forum := &domain.Forum{
				Slug: post.Forum,
			}
			forum, err := u.forumRepo.GetForumBySlug(ctx, forum)
			if err != nil {
				return postDetails, err
			}
			postDetails.Forum = *forum
		case "thread":
			thread := &domain.Thread{
				ID: post.Thread,
			}
			thread, err := u.threadRepo.GetThreadByID(ctx, thread)
			if err != nil {
				return postDetails, err
			}
			postDetails.Thread = *thread
		}
	}
	return postDetails, nil
}

func (u *Usecase) PostUpdateMessage(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	return u.postRepo.UpdateMessage(ctx, post)
}

func (u *Usecase) PostsCreate(ctx context.Context, posts []*domain.Post, thread *domain.Thread) ([]domain.Post, error) {
	if thread.Slug != "" {
		_, err := u.threadRepo.GetThreadBySlug(ctx, thread)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := u.threadRepo.GetThreadByID(ctx, thread)
		if err != nil {
			return nil, err
		}
	}

	if len(posts) == 0 {
		return nil, pkg.ErrPostsNotGiven
	}

	user := &domain.User{
		Nickname: posts[0].Author,
	}
	_, err := u.userRepo.GetUserByNickname(ctx, user)
	if err != nil {
		return nil, err
	}

	res, err := u.postRepo.PostsCreate(ctx, posts, thread)
	if err != nil {
		return nil, err
	}
	return res, nil
}
