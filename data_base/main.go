package main

import (
	"data_base/core/router"
	"data_base/logger"
	"net/http"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}
	logger.Info.Printf("Started listening at: %s", PORT)
	r := router.GetRouter()
	logger.Fatal.Println(http.ListenAndServe(":"+PORT, r))
}
