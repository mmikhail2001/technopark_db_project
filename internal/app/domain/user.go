package domain

import "github.com/pkg/errors"

type User struct {
	Nickname string
	Email    string
	Fullname string
	About    string
}

var (
	ErrSuchUserExist          = errors.New("such user exist")
	ErrSuchUserNotFound       = errors.New("such user not fount")
	ErrUpdateUserDataConflict = errors.New("impossible update such user data")
)
