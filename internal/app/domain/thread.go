package domain

import (
	"github.com/pkg/errors"
)

type Thread struct {
	ID       int
	Author   string
	Slug     string
	Forum    string
	Title    string
	Message  string
	Created  string
	SumVotes int
}

var (
	ErrSuchThreadNotFound = errors.New("such thread not fount")
	ErrSuchThreadExist    = errors.New("such thread exist")
)
