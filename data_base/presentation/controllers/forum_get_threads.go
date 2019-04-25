package controllers
//
//import (
//	"data_base/models"
//	"data_base/presentation/logger"
//	"encoding/json"
//	"fmt"
//	"github.com/gorilla/mux"
//	"net/http"
//)
//
//func GetThreadsHandler(w http.ResponseWriter, r *http.Request) {
//
//	var forum models.Forum
//
//	varMap := mux.Vars(r)
//	slug, found := varMap["slug"]
//	if !found {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println("not found")
//		return
//	}
//
//	row := models.DB.DatBase.QueryRow(`SELECT slug FROM public."forum" WHERE slug = $1`, slug)
//
//	err := row.Scan(&forum.Slug)
//	if err != nil && err.Error() != ErrorSqlNoRows {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//	if len(forum.Slug) == 0 {
//
//		myJSON := ErrorCantFindSlug + slug + `"}`
//
//		w.WriteHeader(http.StatusNotFound)
//		_, err = w.Write([]byte(myJSON))
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			logger.Error.Println(err.Error())
//			return
//		}
//		return
//	}
//
//	limit := r.URL.Query().Get("limit")
//	since := r.URL.Query().Get("since")
//	desc := r.URL.Query().Get("desc")
//	if desc == "true" {
//		desc = " DESC "
//	} else if desc == "false" {
//		desc = " ASC "
//	}
//
//	rows, err := models.DB.DatBase.Query(
//		fmt.Sprintf(`SELECT * FROM public."thread" WHERE created >= '%s' AND forum = '%s' ORDER BY created`, since, slug) +
//			desc +
//			fmt.Sprintf(`LIMIT %v`, limit))
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//
//	threads := make([]models.Branch, 0)
//	var thread models.Branch
//
//	for rows.Next() {
//		err = rows.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title, &thread.Message, &thread.Votes, &thread.Created)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			logger.Error.Println(err.Error())
//			return
//		}
//		threads = append(threads, thread)
//	}
//
//	data, err := json.Marshal(threads)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	_, err = w.Write(data)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//}
