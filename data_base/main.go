package main

import (
	"data_base/models"
	"data_base/presentation/core/router"
	"data_base/presentation/logger"
	"net/http"
)

func main() {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	models.Host, models.Port, models.User, models.Password, models.DBname)
	//
	//DB, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	logger.Fatal.Println(err.Error())
	//	panic(err)
	//}
	//defer DB.Close()
	//logger.Info.Printf("\nSuccessfully connected to database at: 5432")

	models.OpenConnectionDB()
	defer models.CloseConnectionDB()

	r := router.GetRouter()
	logger.Info.Printf("\nStarted listening at: 5000")
	logger.Fatal.Println(http.ListenAndServe(":5000", r))


}




