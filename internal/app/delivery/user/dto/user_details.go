package dto

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields user_details.go

//easyjson:json
type UserDetailsRequestDTO struct {
	Nickname string
}

func NewUserDetailsRequestDTO() *UserDetailsRequestDTO {
	return &UserDetailsRequestDTO{}
}

func (dto *UserDetailsRequestDTO) Bind(r *http.Request) error {
	vars := mux.Vars(r)
	dto.Nickname = vars["nickname"]
	return nil
}

func (dto *UserDetailsRequestDTO) GetModel() *domain.User {
	return &domain.User{
		Nickname: dto.Nickname,
	}
}

//easyjson:json
type UserDetailsResponseDTO struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
}

func NewUserDetailsResponseDTO(user *domain.User) *UserDetailsResponseDTO {
	return &UserDetailsResponseDTO{
		Nickname: user.Nickname,
		Fullname: user.Fullname,
		About:    user.About,
		Email:    user.Email,
	}
}
