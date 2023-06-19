package post

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mmikhail2001/technopark_db_project/internal/app/delivery/post/dto"
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

func (h *Handler) PostDetails(w http.ResponseWriter, r *http.Request) {

	requestDTO := dto.NewPostDetailsRequestDTO()
	requestDTO.Bind(r)

	posts, err := h.usecase.PostDetails(r.Context(), requestDTO.GetModel(), requestDTO.GetParams())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchPostNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewPostDetailsResponseDTO(posts)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}

func (h *Handler) PostUpdateMessage(w http.ResponseWriter, r *http.Request) {

	requestDTO := dto.NewPostUpdateMessageRequestDTO()
	requestDTO.Bind(r)

	posts, err := h.usecase.PostUpdateMessage(r.Context(), requestDTO.GetModel())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchPostNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewPostUpdateMessageResponseDTO(posts)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}

func (h *Handler) PostsCreate(w http.ResponseWriter, r *http.Request) {

	requestDTO := dto.NewPostsCreateForThreadRequestDTO()
	requestDTO.Bind(r)
	posts, err := h.usecase.PostsCreate(r.Context(), requestDTO.GetModelPosts(), requestDTO.GetModelThread())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrPostsNotGiven):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("[]"))
		case errors.Is(err, pkg.ErrSuchThreadNotFound) || errors.Is(err, pkg.ErrSuchUserNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		case errors.Is(err, pkg.ErrPostParentNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusConflict, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewPostsCreateResponseDTO(posts)
	httptools.Response(r.Context(), w, http.StatusCreated, responseDTO)
}
