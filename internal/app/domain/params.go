package domain

type GetThreadsParams struct {
	Limit int
	Since string
	Desc  bool
}

type GetUsersParams struct {
	Limit int
	Since string
	Desc  bool
}

type GetPostsParams struct {
	Limit int
	Since int
	Desc  bool
	Sort  string
}

type VoteParams struct {
	Nickname string
	Voice    int
}

type PostDetailsParams struct {
	Related []string
}
