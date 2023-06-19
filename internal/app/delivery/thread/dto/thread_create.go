package dto

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields thread_create.go

//easyjson:json
type ThreadCreateRequestDTO struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Created string `json:"created"`
	Forum   string `json:"forum"`
	Slug    string `json:"slug"`
}

func NewThreadCreateRequestDTO() *ThreadCreateRequestDTO {
	return &ThreadCreateRequestDTO{}
}

func (dto *ThreadCreateRequestDTO) Bind(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	easyjson.Unmarshal(body, dto)

	vars := mux.Vars(r)
	dto.Forum = vars["slug"]

	if err != nil {
		return err
	}
	return nil
}

func (dto *ThreadCreateRequestDTO) GetModel() *domain.Thread {
	return &domain.Thread{
		Slug:    dto.Slug,
		Title:   dto.Title,
		Author:  dto.Author,
		Message: dto.Message,
		Created: dto.Created,
		Forum:   dto.Forum,
	}
}

//easyjson:json
type ThreadCreateResponseDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Message  string `json:"message"`
	Created  string `json:"created"`
	Forum    string `json:"forum"`
	Slug     string `json:"slug,omitempty"`
	SumVotes int    `json:"votes"`
}

func NewThreadCreateResponseDTO(thread *domain.Thread) *ThreadCreateResponseDTO {
	return &ThreadCreateResponseDTO{
		ID:       thread.ID,
		Title:    thread.Title,
		Author:   thread.Author,
		Forum:    thread.Forum,
		Message:  thread.Message,
		Created:  thread.Created,
		SumVotes: thread.SumVotes,
		Slug:     thread.Slug,
	}
}
