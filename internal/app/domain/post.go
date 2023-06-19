package domain

import (
	"github.com/pkg/errors"
)

type Post struct {
	ID       int
	Parent   int
	Path     []int
	Author   string
	Forum    string
	Thread   int
	Message  string
	IsEdited bool
	Created  string
}

var (
	ErrNoSuchRuleSortPosts = errors.New("no such rule for sort posts")
	ErrSuchPostNotFound    = errors.New("such post not found")
	ErrPostParentNotFound  = errors.New("such post parent not found")
	ErrInvalidParent       = errors.New("parent not valid")
)
