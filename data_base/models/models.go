package models

import "net/http"

type Route struct {
	Info    string
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Handler struct {
	Info   string
	Name   string
	Path   string
	Method string
}

type Users struct {
	Id       int
	Email    string
	Fullname string
	Nickname string
	About    string
}
