package bootstrap

import "net/http"

type Router struct {
	Info    string
	Name    string
	Path    string
	Method string
	Handler http.HandlerFunc
}

type Handler struct {
	Info    string
	Name    string
	Path    string
	Method string
}
