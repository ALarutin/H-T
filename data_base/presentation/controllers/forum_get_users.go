package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	varMap := mux.Vars(r)
	slug, found := varMap["slug"]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println("not found")
		return
	}

	limit := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	since := r.URL.Query().Get("since")
	sinceInt, err := strconv.Atoi(since)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	desc := r.URL.Query().Get("desc")
	var descBool bool
	if desc == "true" {
		descBool = true
	} else if desc == "false" {
		descBool = false
	}

	users, err := models.GetInstance().GetUsers(slug, sinceInt, descBool, limitInt)
	if err != nil {
		if err.Error() == errorSqlNoRows {
			myJSON := fmt.Sprintf(`{"%s%s%s"}`, messageCantFind, cantFindForum, slug)
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(myJSON))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(err.Error())
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	data, err := json.Marshal(users)
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
