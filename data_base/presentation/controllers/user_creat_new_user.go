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

	user.About = r.PostFormValue("about")
	user.Email = r.PostFormValue("email")
	user.Fullname = r.PostFormValue("fullname")
	user.Nickname = nickname

	_, err = models.DB.DatBase.Exec(
		`INSERT INTO public."person" (email, about, fullname, nickname) 
				VALUES ($1, $2, $3, $4)`, user.Email, user.About, user.Fullname, user.Nickname)
	if err, ok := err.(*pq.Error); ok && err.Code.Class() != ErrorUniqueViolation {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		logger.Error.Println(err.Code.Class())
		return
	} else if ok && err.Code.Class() == ErrorUniqueViolation {

		rows, _err := models.DB.DatBase.Query(`SELECT nickname, email, fullname, about FROM public."person" WHERE nickname = $1 OR email = $2`, nickname, r.PostFormValue("email"))
		if _err, ok := _err.(*pq.Error); ok {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(_err.Error())
			logger.Error.Println(_err.Code.Class())
			return
		}

		users := make([]models.Users, 0)
		var i int

		for rows.Next() {
			i++
			_err := rows.Scan(&user.Nickname, &user.Email, &user.Fullname, &user.About)
			if _err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(_err.Error())
				return
			}
			users = append(users, user)
		}

		if i != 0 {
			data, _err := json.Marshal(users)
			if _err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(_err.Error())
				return
			}

			w.WriteHeader(http.StatusConflict)
			_, _err = w.Write(data)
			if _err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(_err.Error())
				return
			}
			return
		}
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
