package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"net/http"
)

func ChangUserDataHandler(w http.ResponseWriter, r *http.Request) {

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

	err = models.GetInstance().UpdateUser(user)
	if pqErr, ok := err.(*pq.Error); ok || err != nil {
		logger.Error.Print(err.Error())
		if ok && pqErr.Code.Class() == ErrorUniqueViolation {

			myJSON := fmt.Sprintf(`{"message": "email %s has already taken by another user"}`, user.Email)

			w.WriteHeader(http.StatusConflict)
			_, err := w.Write([]byte(myJSON))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(err.Error())
				return
			}
			return
		}

		if err.Error() == ErrorSqlNoRows {

			myJSON := fmt.Sprintf("{\"%s%s\"}", ErrorCantFindUser, nickname)

			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(myJSON))
			if err != nil {
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

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	return
}
