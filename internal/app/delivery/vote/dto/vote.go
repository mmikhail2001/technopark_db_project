package dto

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	easyjson "github.com/mailru/easyjson"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields vote.go

//easyjson:json
type VoteRequestDTO struct {
	ThreadSlugOrID string
	Nickname       string `json:"nickname"`
	Vote           int    `json:"voice"`
}

func NewVoteRequestDTO() *VoteRequestDTO {
	return &VoteRequestDTO{}
}

func (dto *VoteRequestDTO) Bind(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	easyjson.Unmarshal(body, dto)
	if err != nil {
		return err
	}

	vars := mux.Vars(r)
	dto.ThreadSlugOrID = vars["slug_or_id"]
	return nil
}

func (dto *VoteRequestDTO) GetModelVote() *domain.Vote {
	return &domain.Vote{
		Author: dto.Nickname,
		Vote:   dto.Vote,
	}
}

func (dto *VoteRequestDTO) GetModelThread() *domain.Thread {
	id, err := strconv.Atoi(dto.ThreadSlugOrID)
	if err == nil {
		return &domain.Thread{
			ID: id,
		}
	}
	return &domain.Thread{
		Slug: dto.ThreadSlugOrID,
	}
}

//easyjson:json
type VoteThreadResponseDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Forum    string `json:"forum"`
	Slug     string `json:"slug"`
	Message  string `json:"message"`
	Created  string `json:"created"`
	SumVotes int    `json:"votes"`
}

func NewVoteThreadResponseDTO(thread *domain.Thread) *VoteThreadResponseDTO {
	return &VoteThreadResponseDTO{
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
