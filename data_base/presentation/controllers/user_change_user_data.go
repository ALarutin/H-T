package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func ChangUserDataHandler(w http.ResponseWriter, r *http.Request) {

	var user models.Users

	varMap := mux.Vars(r)
	nickname, found := varMap["nickname"]
	if !found {
		return
	}

	err := r.ParseForm()
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	row := models.DB.DatBase.QueryRow(`SELECT * FROM public."user" WHERE nickname = $1`, nickname)

	err = row.Scan(&user.About, &user.Email, &user.Fullname, &user.Nickname)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	if len(user.Nickname) == 0 {
		MyJSON := `{"message":"` + "cant find user with nickname " + nickname + `"}`

		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(MyJSON))
		if err != nil {
			logger.Error.Println(err.Error())
			return
		}
		return
	}

	row = models.DB.DatBase.QueryRow(`SELECT * FROM public."user" WHERE email = $1`, r.PostFormValue("email"))

	err = row.Scan(&user.About, &user.Email, &user.Fullname, &user.Nickname)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	if len(user.Nickname) != 0 && user.Nickname != nickname {
		MyJSON := `{"message":"` + "this email " + user.Email + " is already taken by another user" + `"}`

		w.WriteHeader(http.StatusConflict)
		_, err = w.Write([]byte(MyJSON))
		if err != nil {
			logger.Error.Println(err.Error())
			return
		}
		return
	}




	return
}
