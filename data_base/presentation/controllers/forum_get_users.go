package controllers
//
//import (
//	"data_base/models"
//	"data_base/presentation/logger"
//	"encoding/json"
//	"fmt"
//	"github.com/gorilla/mux"
//	"net/http"
//)
//
//func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
//
//	var forum models.Forum
//
//	varMap := mux.Vars(r)
//	slug, found := varMap["slug"]
//	if !found {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println("not found")
//		return
//	}
//
//	row := models.DB.DatBase.QueryRow(`SELECT slug FROM public."forum" WHERE slug = $1`, slug)
//
//	err := row.Scan(&forum.Slug)
//	if err != nil && err.Error() != ErrorSqlNoRows {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//	if len(forum.Slug) == 0 {
//
//		myJSON := ErrorCantFindSlug + slug + `"}`
//
//		w.WriteHeader(http.StatusNotFound)
//		_, err = w.Write([]byte(myJSON))
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			logger.Error.Println(err.Error())
//			return
//		}
//		return
//	}
//
//	limit := r.URL.Query().Get("limit")
//	since := r.URL.Query().Get("since")
//	desc := r.URL.Query().Get("desc")
//	if desc == "true" {
//		desc = " DESC "
//	} else if desc == "false" {
//		desc = " ASC "
//	}
//
//	tempSelect:= fmt.Sprintf(`SELECT * FROM public."person" WHERE nickname in
//					(SELECT user_nickname
//					FROM (SELECT * FROM public."forum_users" GROUP BY forum_slug, user_nickname) as m
//					WHERE forum_slug = '%s')`, slug)
//	rows, err := models.DB.DatBase.Query(
//		fmt.Sprintf(`SELECT nickname, email, fullname, about FROM (%s) as p WHERE id >= '%s' ORDER BY id`,
//			tempSelect, since) +
//			desc +
//			fmt.Sprintf(`LIMIT %v`, limit))
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//
//	users := make([]models.Users, 0)
//	var user models.Users
//
//	for rows.Next() {
//		err = rows.Scan(&user.Nickname, &user.Email, &user.Fullname, &user.About)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			logger.Error.Println(err.Error())
//			return
//		}
//		users = append(users, user)
//	}
//
//	data, err := json.Marshal(users)
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
//}
