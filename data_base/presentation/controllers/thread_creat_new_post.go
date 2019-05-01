package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
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

	posts := make([]models.PostInput, 0)
	err = json.Unmarshal(body, &posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	var authors []string
	var messages []string
	var parents []int

	for _, post := range posts{
		authors = append(authors, post.Author)
		messages = append(messages, post.Message)
		parents = append(parents, post.Parent)
	}
	len := len(posts)

	logger.Info.Print("//////////////////////")
	logger.Info.Print(authors)
	logger.Info.Print(messages)
	logger.Info.Print(parents)

	p, err := models.GetInstance().CreatePost(authors, messages, parents, slug, id, len)
	logger.Info.Print("//////////////////////")
	logger.Error.Print(p)
	logger.Error.Print(err)
	if err != nil {

		if err.Error() == postParentForeignKey {
			myJSON := fmt.Sprintf(`{"%s%sor%s"}`,
				messageCantFind, cantFindParent, cantFindUser)
			w.WriteHeader(http.StatusConflict)
			_, err = w.Write([]byte(myJSON))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Error.Println(err.Error())
				return
			}
			return
		}

		if err.Error() == errorSqlNoRows{
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

		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	data, err := json.Marshal(p)
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
