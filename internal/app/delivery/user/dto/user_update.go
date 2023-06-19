package dto

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields user_update.go

//easyjson:json
type UserUpdateRequestDTO struct {
	Nickname_ string
	Fullname  string `json:"fullname"`
	About     string `json:"about"`
	Email     string `json:"email"`
}

func NewUserUpdateRequestDTO() *UserUpdateRequestDTO {
	return &UserUpdateRequestDTO{}
}

func (dto *UserUpdateRequestDTO) Bind(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	easyjson.Unmarshal(body, dto)
	if err != nil {
		return err
	}

	vars := mux.Vars(r)
	dto.Nickname_ = vars["nickname"]
	return nil
}

func (dto *UserUpdateRequestDTO) GetModel() *domain.User {
	return &domain.User{
		Nickname: dto.Nickname_,
		Fullname: dto.Fullname,
		About:    dto.About,
		Email:    dto.Email,
	}
}

//easyjson:json
type UserUpdateResponseDTO struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
}

func NewUserUpdateResponseDTO(user *domain.User) *UserUpdateResponseDTO {
	return &UserUpdateResponseDTO{
		Nickname: user.Nickname,
		Fullname: user.Fullname,
		About:    user.About,
		Email:    user.Email,
	}
}
