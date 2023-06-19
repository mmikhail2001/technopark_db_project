package forum

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mmikhail2001/technopark_db_project/internal/app/delivery/forum/dto"
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

func (h *Handler) ForumCreate(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.NewForumCreateRequestDTO()
	requestDTO.Bind(r)

	forum, err := h.usecase.ForumCreate(r.Context(), requestDTO.GetModel())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchForumExist):
			responseDTO := dto.NewForumCreateResponseDTO(forum)
			httptools.Response(r.Context(), w, http.StatusConflict, responseDTO)
		case errors.Is(err, pkg.ErrSuchUserNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewForumCreateResponseDTO(forum)
	httptools.Response(r.Context(), w, http.StatusCreated, responseDTO)
}

func (h *Handler) ForumDetails(w http.ResponseWriter, r *http.Request) {

	requestDTO := dto.NewForumDetailsRequestDTO()
	requestDTO.Bind(r)

	forum, err := h.usecase.ForumDetails(r.Context(), requestDTO.GetModel())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchForumNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewForumDetailsResponseDTO(forum)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}

func (h *Handler) ForumUsers(w http.ResponseWriter, r *http.Request) {

	requestDTO := dto.NewForumUsersRequestDTO()
	requestDTO.Bind(r)

	users, err := h.usecase.ForumUsers(r.Context(), requestDTO.GetModel(), requestDTO.GetParams())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchForumNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewForumUsersResponseDTO(users)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}

func (h *Handler) ForumThreads(w http.ResponseWriter, r *http.Request) {

	requestDTO := dto.NewForumThreadsRequestDTO()
	requestDTO.Bind(r)

	users, err := h.usecase.ForumThreads(r.Context(), requestDTO.GetModel(), requestDTO.GetParams())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchForumNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewForumThreadsResponseDTO(users)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}
