package models

import (
	"net/http"
)

type Route struct {
	Info    string
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Users struct {
	About    string
	Email    string
	Fullname string
	Nickname string
}
