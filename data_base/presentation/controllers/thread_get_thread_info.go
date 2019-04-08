package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetThreadInfoThreadHandler(w http.ResponseWriter, r *http.Request) {

	varMap := mux.Vars(r)
	slug, found := varMap["slug_or_id"]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println("not found")
		return
	}

	i, err := strconv.Atoi(slug)
	if err != nil{
		i = 0
		logger.Error.Println(err.Error())
	}

	row := models.DB.DatBase.QueryRow(`SELECT * FROM public."thread" WHERE slug = $1 OR id = $2`, slug, i)

	var thread models.Branch

	err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title, &thread.Message, &thread.Votes, &thread.Created)
	if err != nil && err.Error() != ErrorSqlNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	if len(thread.Slug) == 0 {

		myJSON := ErrorCantFindThread + slug + `"}`

		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(myJSON))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
		return
	}

	data, err := json.Marshal(thread)
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
}
