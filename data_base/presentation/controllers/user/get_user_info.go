package user

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
	} else{
		MyJSON := `{"message":"`+ "some message here" + `"}`

		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(MyJSON))
		if err != nil {
			logger.Error.Println(err.Error())
			return
		}
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
}
