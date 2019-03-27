package models

import (
	"data_base/presentation/logger"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "1209qawsed"
	DBname   = "postgres"
)

type environment struct{
	DatBase *sql.DB
}

var DB environment

func OpenConnectionDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Fatal.Println(err.Error())
		panic(err)
	}

	DB.DatBase = db

	logger.Info.Printf("\nSuccessfully connected to database at: 5432")
}

func CloseConnectionDB(){
	DB.DatBase.Close()
}
