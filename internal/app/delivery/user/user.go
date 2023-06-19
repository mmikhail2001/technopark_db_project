package user

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mmikhail2001/technopark_db_project/internal/app/delivery/user/dto"
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

func (h *Handler) UserCreate(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.NewUserCreateRequestDTO()
	requestDTO.Bind(r)

	users, err := h.usecase.UserCreate(r.Context(), requestDTO.GetModel())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchUserExist):
			responseDTO := dto.NewUsersCreateResponseDTO(users)
			httptools.Response(r.Context(), w, http.StatusConflict, responseDTO)
		case errors.Is(err, pkg.ErrNotEnoughData):
			httptools.ResponseError(r.Context(), w, http.StatusConflict, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewUserCreateResponseDTO(users[0])
	httptools.Response(r.Context(), w, http.StatusCreated, responseDTO)
}

func (h *Handler) UserDetails(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.NewUserDetailsRequestDTO()
	requestDTO.Bind(r)

	thread, err := h.usecase.UserDetails(r.Context(), requestDTO.GetModel())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchUserNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewUserDetailsResponseDTO(thread)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}

func (h *Handler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.NewUserUpdateRequestDTO()
	requestDTO.Bind(r)

	user, err := h.usecase.UserUpdate(r.Context(), requestDTO.GetModel())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchUserNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		case errors.Is(err, pkg.ErrSuchUserExist):
			httptools.ResponseError(r.Context(), w, http.StatusConflict, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewUserUpdateResponseDTO(user)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}
