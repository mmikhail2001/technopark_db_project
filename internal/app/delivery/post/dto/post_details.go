package dto

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

//go:generate easyjson -all -disallow_unknown_fields post_details.go

//easyjson:json
type PostDetailsRequestDTO struct {
	ID      int
	Related []string
}

func NewPostDetailsRequestDTO() *PostDetailsRequestDTO {
	return &PostDetailsRequestDTO{}
}

func (dto *PostDetailsRequestDTO) Bind(r *http.Request) error {
	vars := mux.Vars(r)
	param := vars["id"]

	value, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	dto.ID = value
	param = r.URL.Query().Get("related")
	dto.Related = strings.Split(param, ",")
	return nil
}

func (dto *PostDetailsRequestDTO) GetModel() *domain.Post {
	return &domain.Post{
		ID: dto.ID,
	}
}

func (dto *PostDetailsRequestDTO) GetParams() *domain.PostDetailsParams {
	return &domain.PostDetailsParams{
		Related: dto.Related,
	}
}

//easyjson:json
type PostDetailsPostResponseDTO struct {
	ID       int    `json:"id"`
	Parent   int    `json:"parent"`
	Author   string `json:"author"`
	Message  string `json:"message"`
	IsEdited bool   `json:"isEdited"`
	Forum    string `json:"forum"`
	Thread   int    `json:"thread"`
	Created  string `json:"created"`
}

//easyjson:json
type PostDetailsAuthorResponseDTO struct {
	Nickname string `json:"nickname"`
	Fullname string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

//easyjson:json
type PostDetailsThreadResponseDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Forum    string `json:"forum"`
	Slug     string `json:"slug,omitempty"`
	Message  string `json:"message"`
	Created  string `json:"created"`
	SumVotes int    `json:"votes"`
}

//easyjson:json
type PostDetailsForumResponseDTO struct {
	Title        string `json:"title"`
	Author       string `json:"user"`
	Slug         string `json:"slug"`
	CountPosts   int    `json:"posts"`
	CountThreads int    `json:"threads"`
}

//easyjson:json
type PostDetailsResponseDTO struct {
	Post   *PostDetailsPostResponseDTO   `json:"post,omitempty"`
	Thread *PostDetailsThreadResponseDTO `json:"thread,omitempty"`
	Author *PostDetailsAuthorResponseDTO `json:"author,omitempty"`
	Forum  *PostDetailsForumResponseDTO  `json:"forum,omitempty"`
}

func NewPostDetailsResponseDTO(postDetails *domain.PostDetails) *PostDetailsResponseDTO {
	res := &PostDetailsResponseDTO{}

	if postDetails.Post.ID != 0 {
		post := PostDetailsPostResponseDTO{
			ID:       postDetails.Post.ID,
			Parent:   postDetails.Post.Parent,
			Author:   postDetails.Post.Author,
			Forum:    postDetails.Post.Forum,
			Thread:   postDetails.Post.Thread,
			Message:  postDetails.Post.Message,
			Created:  postDetails.Post.Created,
			IsEdited: postDetails.Post.IsEdited,
		}

		res.Post = &post
	}

	if postDetails.Author.Nickname != "" {
		author := PostDetailsAuthorResponseDTO{
			Nickname: postDetails.Author.Nickname,
			Fullname: postDetails.Author.Fullname,
			About:    postDetails.Author.About,
			Email:    postDetails.Author.Email,
		}

		res.Author = &author
	}

	if postDetails.Thread.ID != 0 {
		thread := PostDetailsThreadResponseDTO{
			ID:       postDetails.Thread.ID,
			Title:    postDetails.Thread.Title,
			Author:   postDetails.Thread.Author,
			Forum:    postDetails.Thread.Forum,
			Slug:     postDetails.Thread.Slug,
			Message:  postDetails.Thread.Message,
			Created:  postDetails.Thread.Created,
			SumVotes: postDetails.Thread.SumVotes,
		}

		res.Thread = &thread
	}

	if postDetails.Forum.Author != "" {
		forum := PostDetailsForumResponseDTO{
			Title:        postDetails.Forum.Title,
			Author:       postDetails.Forum.Author,
			Slug:         postDetails.Forum.Slug,
			CountPosts:   postDetails.Forum.CountPosts,
			CountThreads: postDetails.Forum.CountThreads,
		}

		res.Forum = &forum
	}

	return res
}
