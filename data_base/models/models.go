package models

type Users struct {
	About    string
	Email    string
	Fullname string
	Nickname string
}

type Forum struct {
	Posts   int64
	Threads int64
	Title   string
	Author  string
	Slug    string
}

type Branch struct {
	Author  string
	Created string
	Forum   string
	ID      int64
	Message string
	Slug    string
	Title   string
	Votes   int64
}
