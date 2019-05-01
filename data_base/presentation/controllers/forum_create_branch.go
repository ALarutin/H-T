package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func CreateBranchHandler(w http.ResponseWriter, r *http.Request) {

	varMap := mux.Vars(r)
	slugUrl, found := varMap["slug"]
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

	nickname := r.PostFormValue("author")

	slugBody := r.PostFormValue("slug")
	if len(slugBody) == 0 {
		title := r.PostFormValue("title")
		slugBody = strings.Replace(strings.ToLower(title), " ", "_", -1)
	}

	var thread models.Thread

	thread.Author = nickname
	thread.Created = r.PostFormValue("created")
	thread.Forum = slugUrl
	thread.Message = r.PostFormValue("message")
	thread.Title = r.PostFormValue("title")
	thread.Slug = slugBody

	t, err := models.GetInstance().CreateThread(thread)
	if err != nil {
		if err.Error() == errorSqlNoRows {
			myJSON := fmt.Sprintf(`{"%s%s%s or %s%s"}`,
				messageCantFind, cantFindUser, thread.Author, cantFindForum, thread.Forum)
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

	data, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	if t.IsNew == false {
		w.WriteHeader(http.StatusConflict)
		_, err = w.Write(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
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
