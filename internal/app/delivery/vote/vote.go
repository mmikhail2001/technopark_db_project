package vote

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mmikhail2001/technopark_db_project/internal/app/delivery/vote/dto"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg"
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

func (h *Handler) Vote(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.NewVoteRequestDTO()
	requestDTO.Bind(r)

	user, err := h.usecase.Vote(r.Context(), requestDTO.GetModelVote(), requestDTO.GetModelThread())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchThreadNotFound) || errors.Is(err, pkg.ErrSuchUserNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewVoteThreadResponseDTO(user)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}
