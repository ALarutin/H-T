package controllers

import (
	"database/sql"
	"net/http"
)

const (
	ErrorSqlNoRows = `sql: no rows in result set`
	ErrorCantFindUser    = `{"message": "cant find user with nickname `
	ErrorCantFindSlug = `{"message": "cant find forum with slug `
	ErrorHaveDuplicates = `pq: duplicate key value violates unique constraint "forum_users_pk"`
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
