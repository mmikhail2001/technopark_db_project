package dto

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields user_create.go

//easyjson:json
type UserCreateRequestDTO struct {
	Nickname string
	Fullname string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

func NewUserCreateRequestDTO() *UserCreateRequestDTO {
	return &UserCreateRequestDTO{}
}

func (dto *UserCreateRequestDTO) Bind(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	easyjson.Unmarshal(body, dto)
	if err != nil {
		return err
	}

	vars := mux.Vars(r)
	dto.Nickname = vars["nickname"]
	return nil
}

func (dto *UserCreateRequestDTO) GetModel() *domain.User {
	return &domain.User{
		Nickname: dto.Nickname,
		Fullname: dto.Fullname,
		About:    dto.About,
		Email:    dto.Email,
	}
}

//easyjson:json
type UserDTO struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
}

//easyjson:json
type UsersDTO []UserDTO

func NewUserCreateResponseDTO(user *domain.User) *UserDTO {
	return &UserDTO{
		About:    user.About,
		Email:    user.Email,
		Fullname: user.Fullname,
		Nickname: user.Nickname,
	}
}

func NewUsersCreateResponseDTO(users []*domain.User) UsersDTO {
	res := make([]UserDTO, len(users))

	for idx, value := range users {
		res[idx] = UserDTO{
			Nickname: value.Nickname,
			Fullname: value.Fullname,
			About:    value.About,
			Email:    value.Email,
		}
	}

	return res
}
