package dto

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields forum_threads.go

//easyjson:skip
type ForumThreadsRequestDTO struct {
	Slug  string
	Limit int
	Since string
	Desc  bool
}

func NewForumThreadsRequestDTO() *ForumThreadsRequestDTO {
	return &ForumThreadsRequestDTO{}
}

func (dto *ForumThreadsRequestDTO) Bind(r *http.Request) error {
	vars := mux.Vars(r)
	dto.Slug = vars["slug"]

	param := r.URL.Query().Get("limit")
	if param != "" {
		value, _ := strconv.Atoi(param)
		dto.Limit = value
	} else {
		dto.Limit = 100
	}

	dto.Since = r.URL.Query().Get("since")

	param = r.URL.Query().Get("desc")

	if param == "true" {
		dto.Desc = true
	} else {
		dto.Desc = false

	}

	return nil
}

func (dto *ForumThreadsRequestDTO) GetModel() *domain.Forum {
	return &domain.Forum{
		Slug: dto.Slug,
	}
}

func (dto *ForumThreadsRequestDTO) GetParams() *domain.GetThreadsParams {
	return &domain.GetThreadsParams{
		Limit: dto.Limit,
		Since: dto.Since,
		Desc:  dto.Desc,
	}
}

//easyjson:json
type ThreadDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Forum    string `json:"forum"`
	Slug     string `json:"slug,omitempty"`
	Message  string `json:"message"`
	Created  string `json:"created"`
	SumVotes int    `json:"votes"`
}

//easyjson:json
type ThreadsDTO []*ThreadDTO

func NewForumThreadsResponseDTO(threads []*domain.Thread) ThreadsDTO {
	response := make([]*ThreadDTO, len(threads))

	for i, value := range threads {
		response[i] = &ThreadDTO{
			ID:       value.ID,
			Title:    value.Title,
			Author:   value.Author,
			Forum:    value.Forum,
			Slug:     value.Slug,
			Message:  value.Message,
			Created:  value.Created,
			SumVotes: value.SumVotes,
		}
	}
	return response
}
