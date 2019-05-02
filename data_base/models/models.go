package models

type Database struct {
	Forum  int `json:"forum"`
	Post   int `json:"post"`
	Thread int `json:"thread"`
	User   int `json:"user"`
}

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
}

type PostInfo struct {
	Person *User   `json:"author,omitempty"`
	Forum  *Forum  `json:"forum,omitempty"`
	Post   Post    `json:"post"`
	Thread *Thread `json:"thread,omitempty"`
}

type Vote struct {
	ThreadSlug string
	Nickname   string
	Voice      int
}
