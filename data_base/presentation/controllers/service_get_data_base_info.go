package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"net/http"
)

func GetDataBaseInfoHandler(w http.ResponseWriter, r *http.Request) {

	database, err := models.GetInstance().GetDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	data, err := json.Marshal(database)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
}
