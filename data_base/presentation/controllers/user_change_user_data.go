package controllers
//
//import (
//	"data_base/models"
//	"data_base/presentation/logger"
//	"encoding/json"
//	"github.com/gorilla/mux"
//	"github.com/lib/pq"
//	"net/http"
//)
//
//func ChangUserDataHandler(w http.ResponseWriter, r *http.Request) {
//
//	var user models.Users
//
//	varMap := mux.Vars(r)
//	nickname, found := varMap["nickname"]
//	if !found {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println("not found")
//		return
//	}
//
//	err := r.ParseForm()
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//
//	user.About = r.PostFormValue("about")
//	user.Email = r.PostFormValue("email")
//	user.Fullname = r.PostFormValue("fullname")
//	user.Nickname = nickname
//
//	var id int
//	err = models.DB.DatBase.QueryRow(
//		`UPDATE public."person"
//		SET email = $1, fullname = $2, about = $3
//		WHERE nickname = $4 RETURNING id`,
//		user.Email, user.Fullname, user.About, user.Nickname).
//		Scan(&id)
//	if err, ok := err.(*pq.Error); ok && err.Code.Class() != ErrorUniqueViolation {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		logger.Error.Println(err.Code.Class())
//		return
//	} else if ok && err.Code.Class() == ErrorUniqueViolation {
//		myJSON := `{"message": "this email ` + user.Email + ` is already taken by another user"}` //TODO исправить
//
//		w.WriteHeader(http.StatusConflict)
//		_, _err := w.Write([]byte(myJSON))
//		if _err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			logger.Error.Println(_err.Error())
//			return
//		}
//		return
//	} else if id == 0 {
//		myJSON := ErrorCantFindUser + nickname + `"}`
//
//		w.WriteHeader(http.StatusNotFound)
//		_, _err := w.Write([]byte(myJSON))
//		if _err != nil {
//			logger.Error.Println(_err.Error())
//			return
//		}
//		return
//	}
//
//	data, err := json.Marshal(user)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	_, err = w.Write(data)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//	return
//}
