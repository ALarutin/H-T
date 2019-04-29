package models

type User struct {
	About    string
	Email    string
	Fullname string
	Nickname string
}

type Forum struct {
	Posts   int    `json:"posts"`
	Slug    string `json:"slug"`
	Threads int    `json:"threads"`
	Title   string `json:"title"`
	User    string `json:"user"`
}

type Thread struct {
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
	ThreadSlug string
	Nickname   string
	Voice      int
}
