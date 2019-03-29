package controllers

import "database/sql"

const (
	//ErrorDuplicateKey = `pq: duplicate key value violates unique constraint "user_pkey"`
	//ErrorCantFind     = `Can't find user with id`
)

type Handler struct{
	DB *sql.DB
}