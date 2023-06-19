package domain

import "github.com/pkg/errors"

type Forum struct {
	Slug         string
	Title        string
	Author       string
	CountPosts   int
	CountThreads int
}

var (
	ErrSuchForumNotFound = errors.New("such forum not fount")
	ErrSuchForumExist    = errors.New("such forum exist")
)
