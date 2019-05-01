package models

type User struct {
	IsNew    bool   `json:"-"`
	ID       int    `json:"-"`
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
}

type Forum struct {
	IsNew   bool   `json:"-"`
	ID      int    `json:"-"`
	Posts   int    `json:"posts"`
	Slug    string `json:"slug"`
	Threads int    `json:"threads"`
	Title   string `json:"title"`
	User    string `json:"user"`
}

type Thread struct {
	IsNew   bool   `json:"-"`
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
	Created  string `json:"created"`
	Forum    string `json:"forum"`
	ID       int    `json:"id"`
	IsEdited bool   `json:"isEdited"`
	Message  string `json:"message"`
	Parent   int    `json:"parent"`
	Thread   int    `json:"threads"`
	Path     []int  `json:"-"`
}

type PostInput struct{
	Author   string `sql:"author"`
	Message  string `sql:"message"`
	Parent   int    `sql:"parent"`
}

type Vote struct {
	ThreadSlug string
	Nickname   string
	Voice      int
}
