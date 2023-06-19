package pkg

import "errors"

var (
	// Common delivery
	ErrBadBodyRequest                      = errors.New("bad body request")
	ErrJSONUnexpectedEnd                   = errors.New("unexpected end of JSON input")
	ErrContentTypeUndefined                = errors.New("content-type undefined")
	ErrUnsupportedMediaType                = errors.New("unsupported media type")
	ErrEmptyBody                           = errors.New("empty body")
	ErrConvertQueryType                    = errors.New("bad input query")
	ErrQueryRequiredEmpty                  = errors.New("miss query params")
	ErrBadRequestParams                    = errors.New("bad query params")
	ErrBadRequestParamsEmptyRequiredFields = errors.New("bad params, empty required field")
	ErrGetEasyJSON                         = errors.New("err get easyjson")

	// Common repository
	ErrNotFoundInDB             = errors.New("not found")
	ErrWorkDatabase             = errors.New("error sql")
	ErrGetParamsConvert         = errors.New("err get sql params")
	ErrUnsupportedSortParameter = errors.New("unsupported sort parameter")

	// Middleware
	ErrBigRequest    = errors.New("big request")
	ErrConvertLength = errors.New("getting content-length failed")

	// User
	ErrNotEnoughData            = errors.New("not enough data for create user")
	ErrSuchUserExist            = errors.New("such user exist")
	ErrSuchUserNotFound         = errors.New("such user not found")
	ErrUsersByForumSlugNotFound = errors.New("users by forum slug not found")
	ErrUpdateUserDataConflict   = errors.New("impossible update such user data")

	// Thread
	ErrSuchThreadNotFound = errors.New("such thread not found")
	ErrSuchThreadExist    = errors.New("such thread exist")

	// Post
	ErrNoSuchRuleSortPosts = errors.New("no such rule for sort posts")
	ErrSuchPostNotFound    = errors.New("such post not found")
	ErrPostsNotGiven       = errors.New("posts not given")
	ErrPostParentNotFound  = errors.New("such post parent not found")
	ErrInvalidParent       = errors.New("parent not valid")

	// Forum
	ErrSuchForumNotFound = errors.New("such forum not found")
	ErrSuchForumExist    = errors.New("such forum exist")
)
