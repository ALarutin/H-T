package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func VoteThreadHandler(w http.ResponseWriter, r *http.Request) {

	varMap := mux.Vars(r)
	slug, found := varMap["slug_or_id"]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println("not found")
		return
	}

	id, err := strconv.Atoi(slug)
	if err != nil {
		id = 0
		logger.Error.Println(err.Error())
	} else{
		slug = ""
	}
 	logger.Error.Print(slug)

	//row := models.DB.DatBase.QueryRow(`SELECT * FROM public."thread" WHERE slug = $1 OR id = $2`, slug, i)
	//
	//var thread models.Branch
	//
	//err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title, &thread.Message, &thread.Votes, &thread.Created)
	//if err != nil && err.Error() != ErrorSqlNoRows {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	logger.Error.Println(err.Error())
	//	return
	//}
	//if len(thread.Slug) == 0 {
	//
	//	myJSON := ErrorCantFindThread + slug + `"}`
	//
	//	w.WriteHeader(http.StatusNotFound)
	//	_, err = w.Write([]byte(myJSON))
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		logger.Error.Println(err.Error())
	//		return
	//	}
	//	return
	//}
	//

	//vote.Voice = i
	//vote.Thread = slug
	//
	//row = models.DB.DatBase.QueryRow(`SELECT nickname FROM public."person" WHERE nickname = $1`, vote.Nickname)
	//
	//var user models.Users
	//
	//err = row.Scan(&user.Nickname)
	//if err != nil && err.Error() != ErrorSqlNoRows {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	logger.Error.Println(err.Error())
	//	return
	//}
	//if len(user.Nickname) == 0 {
	//
	//	myJSON := ErrorCantFindUser + vote.Nickname + `"}`
	//
	//	w.WriteHeader(http.StatusNotFound)
	//	_, err = w.Write([]byte(myJSON))
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		logger.Error.Println(err.Error())
	//		return
	//	}
	//	return
	//}

	thread, err := models.GetInstance().GetThread(slug, id)
	if err != nil  {
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

	var vote models.Vote
	nickname := r.PostFormValue("nickname")
	voice := r.PostFormValue("voice")
	i, err := strconv.Atoi(voice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	vote.ThreadSlug = thread.Slug
	vote.Nickname = nickname
	vote.Voice = i

	err = models.GetInstance().CreateOrUpdateVote(vote)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	thread, err = models.GetInstance().GetThread(slug, id)
	if err != nil  {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	data, err := json.Marshal(thread)
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
