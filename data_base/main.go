package main

import (
	"data_base/presentation/core/router"
	"data_base/presentation/logger"
	"database/sql"
	"net/http"
)

func main() {
	connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		logger.Fatal.Printf("%s", err.Error())
	}
	defer db.Close()
	r := router.GetRouter()
	logger.Info.Printf("\nStarted listening at: 8001")
	logger.Fatal.Println(http.ListenAndServe(":5000", r))
}
