package controllers

import (
	"database/sql"
	"net/http"
)

const (
	ErrorSqlNoRows = `sql: no rows in result set`
	ErrorCantFindUser    = `{"message": "cant find user with nickname `
)

type Handler struct {
	DB *sql.DB
}

type Route struct {
	Info    string
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}
