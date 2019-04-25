package controllers
//
//import (
//	"data_base/models"
//	"data_base/presentation/logger"
//	"encoding/json"
//	"github.com/gorilla/mux"
//	"net/http"
//	"strconv"
//)
//
//func VoteThreadHandler(w http.ResponseWriter, r *http.Request) {
//
//	varMap := mux.Vars(r)
//	slug, found := varMap["slug_or_id"]
//	if !found {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println("not found")
//		return
//	}
//
//	i, err := strconv.Atoi(slug)
//	if err != nil{
//		i = 0
//		logger.Error.Println(err.Error())
//	}
//
//	row := models.DB.DatBase.QueryRow(`SELECT * FROM public."thread" WHERE slug = $1 OR id = $2`, slug, i)
//
//	var thread models.Branch
//
//	err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title, &thread.Message, &thread.Votes, &thread.Created)
//	if err != nil && err.Error() != ErrorSqlNoRows {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//	if len(thread.Slug) == 0 {
//
//		myJSON := ErrorCantFindThread + slug + `"}`
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
//	var vote models.Vote
//
//	vote.Nickname = r.PostFormValue("nickname")
//	voice := r.PostFormValue("voice")
//	i, err = strconv.Atoi(voice)
//	if err != nil{
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//	vote.Voice = i
//	vote.Thread = slug
//
//	row = models.DB.DatBase.QueryRow(`SELECT nickname FROM public."person" WHERE nickname = $1`, vote.Nickname)
//
//	var user models.Users
//
//	err = row.Scan(&user.Nickname)
//	if err != nil && err.Error() != ErrorSqlNoRows {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//	if len(user.Nickname) == 0 {
//
//		myJSON := ErrorCantFindUser + vote.Nickname + `"}`
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
//	_, err = models.DB.DatBase.Exec(
//		`INSERT INTO public."vote" (thread_slug, user_nickname, voice)
//				VALUES ($1, $2, $3)
//				ON CONFLICT ON CONSTRAINT vote_pk DO UPDATE
//				SET voice = $3 WHERE "vote".thread_slug = $1 AND "vote".user_nickname = $2`, vote.Thread, vote.Nickname, vote.Voice)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logger.Error.Println(err.Error())
//		return
//	}
//
//	thread.Votes = thread.Votes + vote.Voice
//
//	data, err := json.Marshal(thread)
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
