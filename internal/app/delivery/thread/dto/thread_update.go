package dto

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields thread_update.go

//easyjson:json
type ThreadUpdateRequestDTO struct {
	SlugOrID string
	Title    string `json:"title"`
	Message  string `json:"message"`
}

func NewThreadUpdateRequestDTO() *ThreadUpdateRequestDTO {
	return &ThreadUpdateRequestDTO{}
}

func (dto *ThreadUpdateRequestDTO) Bind(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	easyjson.Unmarshal(body, dto)
	if err != nil {
		return err
	}

	vars := mux.Vars(r)
	dto.SlugOrID = vars["slug_or_id"]
	return nil
}

func (dto *ThreadUpdateRequestDTO) GetModel() *domain.Thread {
	id, err := strconv.Atoi(dto.SlugOrID)
	if err == nil {
		return &domain.Thread{
			ID:      id,
			Message: dto.Message,
			Title:   dto.Title,
		}
	}
	return &domain.Thread{
		Slug:    dto.SlugOrID,
		Message: dto.Message,
		Title:   dto.Title,
	}
}

//easyjson:json
type ThreadUpdateResponseDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Forum    string `json:"forum"`
	Slug     string `json:"slug"`
	Message  string `json:"message"`
	Created  string `json:"created"`
	SumVotes int    `json:"votes"`
}

func NewThreadUpdateResponseDTO(thread *domain.Thread) *ThreadUpdateResponseDTO {
	return &ThreadUpdateResponseDTO{
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
