package user

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreatNewUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.Users

	varMap := mux.Vars(r)
	nickname, found := varMap["nickname"]
	if !found {
		return
	}

	rows, err := models.DB.DatBase.Query(`SELECT * FROM public."user" WHERE nickname = $1`, nickname)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	if rows.Next(){
		err = rows.Scan(&user.About, &user.Email, &user.Fullname, &user.Nickname)
		if err != nil {
			logger.Error.Println(err.Error())
			return
		}

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
		return
	}

	err = r.ParseForm()
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	user.About = r.PostFormValue("about")
	user.Email = r.PostFormValue("email")
	user.Fullname = r.PostFormValue("fullname")
	user.Nickname = nickname

	_, err = models.DB.DatBase.Exec(
				`INSERT INTO public."user" (email, about, fullname, nickname) 
				VALUES ($1, $2, $3, $4)`, user.Email, user.About, user.Fullname, user.Nickname )
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	data, err := json.Marshal(user)
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
