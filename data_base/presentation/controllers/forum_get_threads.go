package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetThreadsHandler(w http.ResponseWriter, r *http.Request) {

	varMap := mux.Vars(r)
	slug, found := varMap["slug"]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println("not found")
		return
	}

	limit := r.URL.Query().Get("limit")
	since := r.URL.Query().Get("since")
	desc := r.URL.Query().Get("desc")
	if desc == "true" {
		desc = " DESC "
	} else if desc == "false" {
		desc = " ASC "
	}

	threads, err := models.GetInstance().GetThreads(slug, since, desc, limit)
	if err != nil {

		if err.Error() == errorSqlNoRows  {

			myJSON := fmt.Sprintf(`{"%s%s%s"}`, messageCantFind, cantFindForum, slug)

			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(myJSON))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(err.Error())
				return
			}
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	data, err := json.Marshal(threads)
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
