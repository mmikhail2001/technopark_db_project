package dto

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields post_update_message.go

//easyjson:json
type PostUpdateMessageRequestDTO struct {
	ID      int
	Message string `json:"message"`
}

func NewPostUpdateMessageRequestDTO() *PostUpdateMessageRequestDTO {
	return &PostUpdateMessageRequestDTO{}
}

func (dto *PostUpdateMessageRequestDTO) Bind(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	easyjson.Unmarshal(body, dto)

	vars := mux.Vars(r)

	param := vars["id"]
	value, _ := strconv.Atoi(param)
	dto.ID = value

	if err != nil {
		return err
	}
	return nil
}

func (dto *PostUpdateMessageRequestDTO) GetModel() *domain.Post {
	return &domain.Post{
		ID:      dto.ID,
		Message: dto.Message,
	}
}

//easyjson:json
type PostUpdateMessageResponseDTO struct {
	ID       int    `json:"id"`
	Parent   int    `json:"parent"`
	Author   string `json:"author"`
	Message  string `json:"message"`
	IsEdited bool   `json:"isEdited"`
	Forum    string `json:"forum"`
	Thread   int    `json:"thread"`
	Created  string `json:"created"`
}

func NewPostUpdateMessageResponseDTO(post *domain.Post) *PostUpdateMessageResponseDTO {
	return &PostUpdateMessageResponseDTO{
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
