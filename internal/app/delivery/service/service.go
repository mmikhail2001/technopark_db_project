package service

import (
	"log"
	"net/http"

	"github.com/mmikhail2001/technopark_db_project/internal/app/delivery/service/dto"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg/httptools"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) ServiceStatus(w http.ResponseWriter, r *http.Request) {
	status, err := h.usecase.ServiceStatus(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	responseDTO := dto.NewServiceStatusResponseDTO(status)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}

func (h *Handler) ServiceClear(w http.ResponseWriter, r *http.Request) {
	err := h.usecase.ServiceClear(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
