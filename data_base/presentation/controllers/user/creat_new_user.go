package user

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

func CreatNewUserHandler(w http.ResponseWriter, r *http.Request) {

	varMap := mux.Vars(r)
	nickname, found := varMap["nickname"]
	if !found {
		return
	}

	user := models.Users{}

	rows, err := models.DB.DatBase.Query(`SELECT * FROM homework_DB."user" WHERE nickname = $1`, nickname)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	for rows.Next(){
		err = rows.Scan(&user.About, &user.Email, &user.Fullname, &user.Nickname)
		if err != nil {
			logger.Error.Println(err.Error())
			return
		}
	}

	if !user.IsEmpty(){
		data, err := json.Marshal(user)
		if err != nil {
			logger.Error.Println(err.Error())
			return
		}
		w.WriteHeader(http.StatusConflict)
		_, err = w.Write(data)
		if err != nil {
			logger.Error.Println(err.Error())
			return
		}
	}

	err = r.ParseForm()
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	about := r.PostFormValue("about")
	email := r.PostFormValue("email")
	fullname := r.PostFormValue("fullname")

	_, err = models.DB.DatBase.Exec(
				`INSERT INTO homework_DB."user" (email, about, fullname, nickname) 
				VALUES ($1, $2, $3, $4)`, email, about, fullname, nickname )
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	form := url.Values{}
	form.Add("about", about)
	form.Add("email", email)
	form.Add("fullname", fullname)
	form.Add("nickname", nickname)

	data, err := json.Marshal(form)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

}
