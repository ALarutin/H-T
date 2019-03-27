package user

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

	user.About = r.PostFormValue("about")
	user.Email = r.PostFormValue("email")
	user.Fullname = r.PostFormValue("fullname")
	user.Nickname = nickname

	rows, err := models.DB.DatBase.Query(`SELECT * FROM public."user" WHERE nickname = $1`, nickname)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	if !rows.Next(){
		MyJSON := `{"message":"`+ "some message here" + `"}`

		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(MyJSON))
		if err != nil {
			logger.Error.Println(err.Error())
			return
		}
		return
	}



	return
}
