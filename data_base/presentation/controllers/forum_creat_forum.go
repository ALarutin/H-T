package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
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

	nickname := r.PostFormValue("user")

	var forum models.Forum

	forum.Slug = r.PostFormValue("slug")
	forum.Title = r.PostFormValue("title")
	forum.User = r.PostFormValue("user")

	_, err = models.DB.DatBase.Exec(
		`INSERT INTO public."forum" (author, slug, title,  posts, threads) 
				VALUES ($1, $2, $3, $4, $5)`,forum.User, forum.Slug, forum.Title, forum.Posts, forum.Threads)
	if err, ok := err.(*pq.Error); ok && err.Code.Class() != ErrorUniqueViolation {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		logger.Error.Println(err.Code.Class())
		return
	} else if ok && err.Code.Class() == ErrorUniqueViolation && err.Constraint == ForumUserForeignKey {

		myJSON := ErrorCantFindUser + nickname + `"}`

		w.WriteHeader(http.StatusNotFound)
		_, _err := w.Write([]byte(myJSON))
		if _err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(_err.Error())
			return
		}
		return
	} else if ok && err.Code.Class() == ErrorUniqueViolation && err.Constraint == ForumPrimaryKey {

		row := models.DB.DatBase.QueryRow(`SELECT * FROM public."forum" WHERE slug = $1`, forum.Slug)

		_err := row.Scan(&forum.Slug, &forum.User, &forum.Title, &forum.Posts, &forum.Threads)
		if _err != nil && _err.Error() != ErrorSqlNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(_err.Error())
			return
		}

		data, _err := json.Marshal(forum)
		if _err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(_err.Error())
			return
		}

		w.WriteHeader(http.StatusConflict)
		_, _err = w.Write(data)
		if _err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(_err.Error())
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
