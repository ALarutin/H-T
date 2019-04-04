package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetForumInfoHandler(w http.ResponseWriter, r *http.Request) {

	var forum models.Forum

	varMap := mux.Vars(r)
	slug, found := varMap["slug"]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println("not found")
		return
	}

	row := models.DB.DatBase.QueryRow(`SELECT * FROM public."forum" WHERE slug = $1`, slug)

	err := row.Scan(&forum.Slug, &forum.Author, &forum.Title, &forum.Posts, &forum.Threads)
	if err != nil && err.Error() != ErrorSqlNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	if len(forum.Slug) == 0 {

		myJSON := ErrorCantFindSlug + slug + `"}`

		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(myJSON))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
		return
	}

	data, err := json.Marshal(forum)
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
