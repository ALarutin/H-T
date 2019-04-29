package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
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

	id, err := strconv.Atoi(slug)
	if err != nil{
		id = 0
		logger.Error.Println(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	posts := make([]models.Post, 0)
	err = json.Unmarshal(body, &posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	ps := make([]models.Post, 0)
	for _, post := range posts {

		p, err := models.GetInstance().CreatePost(post, slug, id)
		if pqErr, ok := err.(*pq.Error); ok {

			if pqErr.Code.Class() == errorUniqueViolation{

				if pqErr.Constraint == postParentForeignKey || pqErr.Constraint == postAuthorForeignKey {
					myJSON := fmt.Sprintf(`{"%s%s%v or %s%s"}`,
						messageCantFind, cantFindParent, post.Parent, cantFindUser, post.Author)
					w.WriteHeader(http.StatusConflict)
					_, err = w.Write([]byte(myJSON))
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						logger.Error.Println(err.Error())
						return
					}
					return
				}

				if pqErr.Constraint == "" {
					myJSON := fmt.Sprintf(`{"%s%s%s/%d"}`, messageCantFind, cantFindThread, slug, id)
					w.WriteHeader(http.StatusNotFound)
					_, err = w.Write([]byte(myJSON))
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						logger.Error.Println(err.Error())
						return
					}
					return
				}
			}

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
