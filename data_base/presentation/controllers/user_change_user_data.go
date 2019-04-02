package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ChangUserDataHandler(w http.ResponseWriter, r *http.Request) {

	var user models.Users

	varMap := mux.Vars(r)
	nickname, found := varMap["nickname"]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println("not found")
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	row := models.DB.DatBase.QueryRow(`SELECT * FROM public."user" WHERE nickname = $1`, nickname)

	err = row.Scan(&user.About, &user.Email, &user.Fullname, &user.Nickname)
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
			logger.Error.Println(err.Error())
			return
		}
		return
	}

	row = models.DB.DatBase.QueryRow(`SELECT * FROM public."user" WHERE email = $1`, r.PostFormValue("email"))

	err = row.Scan(&user.About, &user.Email, &user.Fullname, &user.Nickname)
	if err != nil && err.Error() != ErrorSqlNoRows{
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	if user.Nickname != nickname {
		myJSON := `{"message": "this email "` + user.Email + `" is already taken by another user"}`

		w.WriteHeader(http.StatusConflict)
		_, err = w.Write([]byte(myJSON))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
		return
	}

	user.About = r.PostFormValue("about")
	user.Email = r.PostFormValue("email")
	user.Fullname = r.PostFormValue("fullname")
	user.Nickname = nickname

	_, err = models.DB.DatBase.Exec(
		`UPDATE public."user" 
				SET email = $1, fullname = $2, about = $3  
				WHERE nickname = $4`, user.Email, user.Fullname, user.About, user.Nickname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
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
	return
}
