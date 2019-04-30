package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreatNewUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

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

	user.About = r.PostFormValue("about")
	user.Email = r.PostFormValue("email")
	user.Fullname = r.PostFormValue("fullname")
	user.Nickname = nickname

	users, err := models.GetInstance().CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	if users[0].IsNew == true {

		data, err := json.Marshal(users[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
		return
	} else {
		data, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusConflict)
		_, err = w.Write(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
	}
}
