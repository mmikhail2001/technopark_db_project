package httptools

//go:generate easyjson -all -disallow_unknown_fields dto.go

//easyjson:json
type ErrResponseDTO struct {
	Message string `json:"message,omitempty"`
}
