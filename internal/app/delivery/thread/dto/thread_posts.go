package dto

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields thread_posts.go

//easyjson:skip
type ThreadPostsRequestDTO struct {
	SlugOrID string
	Limit    int
	Since    int
	Sort     string
	Desc     bool
}

func NewThreadPostsRequestDTO() *ThreadPostsRequestDTO {
	return &ThreadPostsRequestDTO{}
}

func (dto *ThreadPostsRequestDTO) Bind(r *http.Request) error {
	vars := mux.Vars(r)
	dto.SlugOrID = vars["slug_or_id"]

	param := r.URL.Query().Get("limit")
	if param != "" {
		value, _ := strconv.Atoi(param)
		dto.Limit = value
	} else {
		dto.Limit = 100
	}

	param = r.URL.Query().Get("since")
	if param != "" {
		value, _ := strconv.Atoi(param)
		dto.Since = value
	} else {
		dto.Since = -1
	}

	param = r.URL.Query().Get("desc")

	if param == "true" {
		dto.Desc = true
	} else {
		dto.Desc = false

	}
	dto.Sort = r.URL.Query().Get("sort")
	if dto.Sort == "" {
		dto.Sort = "flat"
	}

	return nil
}

func (dto *ThreadPostsRequestDTO) GetModel() *domain.Thread {
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

func (dto *ThreadPostsRequestDTO) GetParams() *domain.GetPostsParams {
	return &domain.GetPostsParams{
		Limit: dto.Limit,
		Since: dto.Since,
		Desc:  dto.Desc,
		Sort:  dto.Sort,
	}
}

//easyjson:json
type PostDTO struct {
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
type PostsDTO []*PostDTO

func NewThreadPostsResponseDTO(posts []*domain.Post) PostsDTO {
	res := make([]*PostDTO, len(posts))

	for i, post := range posts {
		res[i] = &PostDTO{
			ID:       post.ID,
			Parent:   post.Parent,
			Author:   post.Author,
			Forum:    post.Forum,
			Thread:   post.Thread,
			Message:  post.Message,
			Created:  post.Created,
			IsEdited: post.IsEdited,
		}
	}

	return res
}
