package dto

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields posts_create.go

//easyjson:json
type PostRequestDTO struct {
	Parent  int    `json:"parent"`
	Author  string `json:"author"`
	Message string `json:"message"`
}

//easyjson:json
type PostsRequestDTO []PostRequestDTO

//easyjson:skip
type PostsCreateForThreadRequestDTO struct {
	SlugOrID string
	Posts    PostsRequestDTO
}

func NewPostsCreateForThreadRequestDTO() *PostsCreateForThreadRequestDTO {
	return &PostsCreateForThreadRequestDTO{}
}

func (dto *PostsCreateForThreadRequestDTO) Bind(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	easyjson.Unmarshal(body, &dto.Posts)

	vars := mux.Vars(r)

	param := vars["slug_or_id"]
	dto.SlugOrID = param

	return nil
}

func (dto *PostsCreateForThreadRequestDTO) GetModelThread() *domain.Thread {
	id, err := strconv.Atoi(dto.SlugOrID)
	if err == nil {
		return &domain.Thread{
			ID: id,
		}
	}

	return &domain.Thread{
		Slug: dto.SlugOrID,
	}
}

func (dto *PostsCreateForThreadRequestDTO) GetModelPosts() []*domain.Post {
	res := make([]*domain.Post, len(dto.Posts))

	for i, post := range dto.Posts {
		res[i] = &domain.Post{
			Parent:  post.Parent,
			Message: post.Message,
			Author:  post.Author,
		}
	}

	return res
}

//easyjson:json
type PostResponseDTO struct {
	ID       int    `json:"id"`
	Parent   int    `json:"parent"`
	Author   string `json:"author"`
	Message  string `json:"message"`
	IsEdited bool   `json:"isEdited"`
	Forum    string `json:"forum"`
	Thread   int    `json:"thread"`
	Created  string `json:"created"`
}

//easyjson:json
type PostsResponseDTO []*PostResponseDTO

func NewPostsCreateResponseDTO(posts []domain.Post) PostsResponseDTO {
	res := make([]*PostResponseDTO, len(posts))

	for i, post := range posts {
		res[i] = &PostResponseDTO{
			ID:       post.ID,
			Parent:   post.Parent,
			Author:   post.Author,
			Forum:    post.Forum,
			IsEdited: post.IsEdited,
			Message:  post.Message,
			Created:  post.Created,
			Thread:   post.Thread,
		}
	}

	return res
}
