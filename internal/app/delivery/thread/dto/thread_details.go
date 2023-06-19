package dto

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields thread_details.go

//easyjson:json
type ThreadDetailsRequestDTO struct {
	SlugOrID string
}

func NewThreadDetailsRequestDTO() *ThreadDetailsRequestDTO {
	return &ThreadDetailsRequestDTO{}
}

func (dto *ThreadDetailsRequestDTO) Bind(r *http.Request) error {
	vars := mux.Vars(r)
	dto.SlugOrID = vars["slug_or_id"]
	return nil
}

func (dto *ThreadDetailsRequestDTO) GetModel() *domain.Thread {
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

//easyjson:json
type ThreadDetailsResponseDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Forum    string `json:"forum"`
	Slug     string `json:"slug"`
	Message  string `json:"message"`
	Created  string `json:"created"`
	SumVotes int    `json:"votes"`
}

func NewThreadDetailsResponseDTO(thread *domain.Thread) *ThreadDetailsResponseDTO {
	return &ThreadDetailsResponseDTO{
		ID:       thread.ID,
		Title:    thread.Title,
		Author:   thread.Author,
		Forum:    thread.Forum,
		Slug:     thread.Slug,
		Message:  thread.Message,
		Created:  thread.Created,
		SumVotes: thread.SumVotes,
	}
}
