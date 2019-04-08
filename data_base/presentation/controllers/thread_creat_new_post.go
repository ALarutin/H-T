package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreatNewPostHandler(w http.ResponseWriter, r *http.Request) {

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

	row := models.DB.DatBase.QueryRow(`SELECT id, slug, forum FROM public."thread" WHERE slug = $1 OR id = $2`, slug, i)

	var thread models.Branch
	var forum string

	err = row.Scan(&thread.ID, &thread.Slug, &forum)
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	posts := make([]models.Post, 0)
	err = json.Unmarshal(body, &posts)

	var p models.Post

	for _, post := range posts {

		if post.Parent == 0 {
			continue
		}

		row = models.DB.DatBase.QueryRow(`SELECT id FROM public."post" WHERE id = $1 AND slug = $2`, post.Parent, thread.Slug)

		err := row.Scan(&p.ID)
		if err != nil && err.Error() != ErrorSqlNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}
		if p.ID == 0 {

			myJSON := ErrorCantFindParent + string(post.Parent) + `"}`

			w.WriteHeader(http.StatusConflict)
			_, err = w.Write([]byte(myJSON))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(err.Error())
				return
			}
			return
		}
	}

	ps := make([]models.Post, 0)

	for _, post := range posts {

		var p models.Post

		p.Message = post.Message
		p.Author = post.Author
		p.Parent = post.Parent
		p.Forum = forum
		p.Thread = thread.ID

		err = models.DB.DatBase.QueryRow(
			`INSERT INTO public."post" (author, thread, forum, message, parent)
				VALUES ($1, $2, $3, $4, $5) RETURNING id, created`, post.Author, p.Thread, p.Forum, post.Message, post.Parent).
			Scan(&p.ID, &p.Created)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(err.Error())
			return
		}

		ps = append(ps, p)
	}

	data, err := json.Marshal(ps)
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
