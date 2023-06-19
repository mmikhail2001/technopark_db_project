package dto

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields forum_details.go

//easyjson:skip
type ForumDetailsRequestDTO struct {
	Slug string
}

func NewForumDetailsRequestDTO() *ForumDetailsRequestDTO {
	return &ForumDetailsRequestDTO{}
}

func (dto *ForumDetailsRequestDTO) Bind(r *http.Request) error {
	vars := mux.Vars(r)
	dto.Slug = vars["slug"]
	return nil
}

func (dto *ForumDetailsRequestDTO) GetModel() *domain.Forum {
	return &domain.Forum{
		Slug: dto.Slug,
	}
}

//easyjson:json
type ForumDetailsResponseDTO struct {
	Title        string `json:"title"`
	User         string `json:"user"`
	Slug         string `json:"slug"`
	CountPosts   int    `json:"posts"`
	CountThreads int    `json:"threads"`
}

func NewForumDetailsResponseDTO(forum *domain.Forum) *ForumDetailsResponseDTO {
	return &ForumDetailsResponseDTO{
		Title:        forum.Title,
		User:         forum.Author,
		Slug:         forum.Slug,
		CountPosts:   forum.CountPosts,
		CountThreads: forum.CountThreads,
	}
}
