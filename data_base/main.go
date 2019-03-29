package main

import (
	"data_base/models"
	"data_base/presentation/logger"
	"data_base/presentation/router"
	"net/http"
)

func main() {
	models.OpenConnectionDB()
	defer models.CloseConnectionDB()

	r := router.GetRouter()
	logger.Info.Printf("\nStarted listening at: 5000")
	logger.Fatal.Println(http.ListenAndServe(":5000", r))

}
