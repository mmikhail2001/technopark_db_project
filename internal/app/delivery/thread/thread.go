package forum

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mmikhail2001/technopark_db_project/internal/app/delivery/thread/dto"
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

func (h *Handler) ThreadCreate(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.NewThreadCreateRequestDTO()
	requestDTO.Bind(r)

	thread, err := h.usecase.ThreadCreate(r.Context(), requestDTO.GetModel())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchThreadExist):
			responseDTO := dto.NewThreadCreateResponseDTO(thread)
			httptools.Response(r.Context(), w, http.StatusConflict, responseDTO)
		case errors.Is(err, pkg.ErrSuchForumNotFound) || errors.Is(err, pkg.ErrSuchUserNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewThreadCreateResponseDTO(thread)
	httptools.Response(r.Context(), w, http.StatusCreated, responseDTO)
}

func (h *Handler) ThreadDetails(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.NewThreadDetailsRequestDTO()
	requestDTO.Bind(r)

	thread, err := h.usecase.ThreadDetails(r.Context(), requestDTO.GetModel())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchThreadNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewThreadDetailsResponseDTO(thread)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}

func (h *Handler) ThreadUpdate(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.NewThreadUpdateRequestDTO()
	requestDTO.Bind(r)

	thread, err := h.usecase.ThreadUpdate(r.Context(), requestDTO.GetModel())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchThreadNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewThreadUpdateResponseDTO(thread)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}

func (h *Handler) ThreadPosts(w http.ResponseWriter, r *http.Request) {
	requestDTO := dto.NewThreadPostsRequestDTO()
	requestDTO.Bind(r)

	posts, err := h.usecase.ThreadPosts(r.Context(), requestDTO.GetModel(), requestDTO.GetParams())
	if err != nil {
		log.Println(err)
		err = errors.Cause(err)
		switch {
		case errors.Is(err, pkg.ErrSuchThreadNotFound):
			httptools.ResponseError(r.Context(), w, http.StatusNotFound, err)
		case errors.Is(err, pkg.ErrSuchPostNotFound):
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	responseDTO := dto.NewThreadPostsResponseDTO(posts)
	httptools.Response(r.Context(), w, http.StatusOK, responseDTO)
}
