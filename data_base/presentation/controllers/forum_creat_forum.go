package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"net/http"
)

func CreatForumHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	var forum models.Forum

	forum.Slug = r.PostFormValue("slug")
	forum.Title = r.PostFormValue("title")
	forum.User = r.PostFormValue("user")

	err = models.GetInstance().CreateForum(forum)
	if pqErr, ok := err.(*pq.Error); ok {

		if pqErr.Code.Class() == ErrorUniqueViolation {

			if pqErr.Constraint == ForumUserForeignKey {

				myJSON := fmt.Sprintf("{\"%s%s\"}", ErrorCantFindUser, forum.User)

				w.WriteHeader(http.StatusNotFound)
				_, err := w.Write([]byte(myJSON))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error.Println(err.Error())
					return
				}
				return
			}

			if pqErr.Constraint == ForumPrimaryKey {

				forum, err := models.GetInstance().SelectForum(forum.Slug)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error.Println(err.Error())
					return
				}

				data, err := json.Marshal(forum)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error.Println(err.Error())
					return
				}

				w.WriteHeader(http.StatusConflict)
				_, err = w.Write(data)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error.Println(err.Error())
					return
				}
				return
			}
		}

		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(pqErr.Error())
		logger.Error.Println(pqErr.Code.Class())
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
