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
	User    string
	Slug    string
}
