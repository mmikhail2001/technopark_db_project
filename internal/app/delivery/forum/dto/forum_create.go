package dto

import (
	"io"
	"net/http"

	"github.com/mailru/easyjson"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields forum_create.go

//easyjson:json
type ForumCreateRequestDTO struct {
	Title string `json:"title"`
	User  string `json:"user"`
	Slug  string `json:"slug"`
}

func NewForumCreateRequestDTO() *ForumCreateRequestDTO {
	return &ForumCreateRequestDTO{}
}

func (dto *ForumCreateRequestDTO) Bind(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	easyjson.Unmarshal(body, dto)

	if err != nil {
		return err
	}
	return nil
}

func (dto *ForumCreateRequestDTO) GetModel() *domain.Forum {
	return &domain.Forum{
		Title:  dto.Title,
		Author: dto.User,
		Slug:   dto.Slug,
	}
}

//easyjson:json
type ForumCreateResponseDTO struct {
	Slug         string `json:"slug"`
	Title        string `json:"title"`
	User         string `json:"user"`
	CountPosts   int    `json:"posts,omitempty"`
	CountThreads int    `json:"threads,omitempty"`
}

func NewForumCreateResponseDTO(forum *domain.Forum) *ForumCreateResponseDTO {
	return &ForumCreateResponseDTO{
		Title:        forum.Title,
		User:         forum.Author,
		Slug:         forum.Slug,
		CountPosts:   forum.CountPosts,
		CountThreads: forum.CountThreads,
	}
}
