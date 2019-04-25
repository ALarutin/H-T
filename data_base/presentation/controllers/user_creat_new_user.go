package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
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

	err = models.GetInstance().CreateUser(user)
	if pqErr, ok := err.(*pq.Error); ok {

		if ok && pqErr.Code.Class() == ErrorUniqueViolation {

			users, err := models.GetInstance().SelectUsers(user.Nickname, user.Email)
			if pqErr, ok := err.(*pq.Error); ok {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(pqErr.Code.Class())
				logger.Error.Println(pqErr.Error())
				return
			}

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
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(pqErr.Error())
		logger.Error.Println(pqErr.Code.Class())
		return
	}

	data, err := json.Marshal(user)
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
}
