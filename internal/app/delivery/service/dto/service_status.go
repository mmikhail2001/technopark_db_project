package dto

import (
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields service_status.go

//easyjson:json
type ServiceStatusResponseDTO struct {
	User   int `json:"user"`
	Forum  int `json:"forum"`
	Thread int `json:"thread"`
	Post   int `json:"post"`
}

func NewServiceStatusResponseDTO(service *domain.StatusService) *ServiceStatusResponseDTO {
	return &ServiceStatusResponseDTO{
		User:   service.User,
		Forum:  service.Forum,
		Thread: service.Thread,
		Post:   service.Post,
	}
}
