package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"net/http"
)

func CreatForumHandler(w http.ResponseWriter, r *http.Request) {

	var forum models.Forum
	var user models.Users

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	nickname := r.PostFormValue("user")

	row := models.DB.DatBase.QueryRow(`SELECT * FROM public."user" WHERE nickname = $1`, nickname)

	err = row.Scan(&user.About, &user.Email, &user.Fullname, &user.Nickname)
	if err != nil && err.Error() != ErrorSqlNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	if len(user.Nickname) == 0 {

		myJSON := ErrorCantFindUser + nickname + `"}`

		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(myJSON))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
		return
	}

	slug := r.PostFormValue("slug")
	username := r.PostFormValue("user")

	row = models.DB.DatBase.QueryRow(`SELECT * FROM public."forum" WHERE slug = $1 AND username = $2`, slug, username)

	err = row.Scan(&forum.Posts, &forum.Slug, &forum.Threads, &forum.Title, &forum.User)
	if err != nil && err.Error() != ErrorSqlNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
	if forum.Slug == slug {

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

	forum.Slug = r.PostFormValue("slug")
	forum.Title = r.PostFormValue("title")
	forum.User = r.PostFormValue("user")

	_, err = models.DB.DatBase.Exec(
		`INSERT INTO public."forum" (slug, title, username, posts, threads) 
				VALUES ($1, $2, $3, $4, $5)`, forum.Slug, forum.Title, forum.User, forum.Posts, forum.Threads)
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

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}
}
