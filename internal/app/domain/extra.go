package domain

type PostDetails struct {
	Post   Post
	Author User
	Thread Thread
	Forum  Forum
}

type StatusService struct {
	User   int
	Forum  int
	Thread int
	Post   int
}
