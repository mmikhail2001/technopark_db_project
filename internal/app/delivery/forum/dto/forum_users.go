package dto

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields forum_users.go

//easyjson:skip
type ForumUsersRequestDTO struct {
	Slug  string
	Limit int
	Since string
	Desc  bool
}

func NewForumUsersRequestDTO() *ForumUsersRequestDTO {
	return &ForumUsersRequestDTO{}
}

func (dto *ForumUsersRequestDTO) Bind(r *http.Request) error {
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

func (dto *ForumUsersRequestDTO) GetModel() *domain.Forum {
	return &domain.Forum{
		Slug: dto.Slug,
	}
}

func (dto *ForumUsersRequestDTO) GetParams() *domain.GetUsersParams {
	return &domain.GetUsersParams{
		Limit: dto.Limit,
		Since: dto.Since,
		Desc:  dto.Desc,
	}
}

//easyjson:json
type UserDTO struct {
	Nickname string `json:"nickname"`
	Fullname string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

//easyjson:json
type UsersDTO []*UserDTO

func NewForumUsersResponseDTO(users []*domain.User) UsersDTO {
	response := make([]*UserDTO, len(users))

	for i, user := range users {
		response[i] = &UserDTO{
			Nickname: user.Nickname,
			Fullname: user.Fullname,
			About:    user.About,
			Email:    user.Email,
		}
	}
	return response
}
