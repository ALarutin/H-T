package models

type User struct {
	About    string
	Email    string
	Fullname string
	Nickname string
}

type Forum struct {
	Posts   int
	Slug    string
	Threads int
	Title   string
	User    string
}

type Branch struct {
	Author  string `json:"author"`
	Created string `json:"created"`
	Forum   string `json:"forum"`
	ID      int    `json:"id"`
	Message string `json:"message"`
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Votes   int    `json:"votes"`
}

type Post struct {
	Author   string `json:"author"`
	Created  string
	Forum    string
	ID       int
	IsEdited bool
	Message  string `json:"message"`
	Parent   int    `json:"parent"`
	Thread   int
}

type Vote struct {
	Thread   string
	Nickname string
	Voice    int
}
