package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {

	var user models.Users

	varMap := mux.Vars(r)
	nickname, found := varMap["nickname"]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println("not found")
		return
	}

	row := models.DB.DatBase.QueryRow(`SELECT nickname, email, fullname, about FROM public."person" WHERE nickname = $1`, nickname)
	err := row.Scan(&user.Nickname, &user.Email, &user.Fullname, &user.About)
	if err != nil && err.Error() != ErrorSqlNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	if len(user.Nickname) == 0 {
		myJSON := ErrorCantFindUser + nickname + `"}`

		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(myJSON))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
		return
	}

	data, err := json.Marshal(user)
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
