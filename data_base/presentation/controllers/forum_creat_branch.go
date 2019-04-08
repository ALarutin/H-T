package controllers

import (
	"data_base/models"
	"data_base/presentation/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"net/http"
	"strings"
)

func CreatBranchHandler(w http.ResponseWriter, r *http.Request) {

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

	//row := models.DB.DatBase.QueryRow(`SELECT slug FROM public."forum" WHERE slug = $1`, slugUrl)
	//
	//err = row.Scan(&forum.Slug)
	//if err != nil && err.Error() != ErrorSqlNoRows {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	logger.Error.Println(err.Error())
	//	return
	//}
	//if len(forum.Slug) == 0 {
	//
	//	myJSON := ErrorCantFindSlug + slugUrl + `"}`
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

	nickname := r.PostFormValue("author")

	//row := models.DB.DatBase.QueryRow(`SELECT nickname FROM public."person" WHERE nickname = $1`, nickname)
	//
	//err = row.Scan(&user.Nickname)
	//if err != nil && err.Error() != ErrorSqlNoRows{
	//	w.WriteHeader(http.StatusInternalServerError)
	//	logger.Error.Println(err.Error())
	//	return
	//}
	//if len(user.Nickname) == 0 {
	//	myJSON := ErrorCantFindUser + nickname + `"}`
	//
	//	w.WriteHeader(http.StatusNotFound)
	//	_, err = w.Write([]byte(myJSON))
	//	if err != nil {
	//		logger.Error.Println(err.Error())
	//		return
	//	}
	//	return
	//}

	slugBody := r.PostFormValue("slug")
	if len(slugBody) == 0 {
		title := r.PostFormValue("title")
		slugBody = strings.Replace(strings.ToLower(title), " ", "_", -1)
	}
	//
	//row := models.DB.DatBase.QueryRow(`SELECT * FROM public."thread" WHERE slug = $1`, slugBody)
	//
	//err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title, &thread.Message, &thread.Votes, &thread.Created)
	//if err != nil && err.Error() != ErrorSqlNoRows {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	logger.Error.Println(err.Error())
	//	return
	//}
	//if thread.Slug == slugBody {
	//
	//	data, err := json.Marshal(thread)
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		logger.Error.Println(err.Error())
	//		return
	//	}
	//
	//	w.WriteHeader(http.StatusConflict)
	//	_, err = w.Write(data)
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		logger.Error.Println(err.Error())
	//		return
	//	}
	//	return
	//}

	var thread models.Branch

	thread.Author = r.PostFormValue("author")
	thread.Created = r.PostFormValue("created")
	thread.Forum = slugUrl
	thread.Message = r.PostFormValue("message")
	thread.Title = r.PostFormValue("title")
	thread.Slug = r.PostFormValue("slug")
	if len(thread.Slug) == 0 {
		thread.Slug = strings.Replace(strings.ToLower(thread.Title), " ", "_", -1)
	}

	eerr := models.DB.DatBase.QueryRow(
		`INSERT INTO public."thread" (author, created, forum, message, slug, title, votes) 
				VALUES ((SELECT nickname FROM public."person" WHERE nickname = $1), $2, (SELECT slug FROM public."forum" WHERE slug = $3),
				        $4, $5, $6, $7) RETURNING id`,
		thread.Author, thread.Created, thread.Forum, thread.Message, thread.Slug, thread.Title, thread.Votes).Scan(&thread.ID)

	logger.Error.Println(eerr.Error())

	if err, ok := eerr.(*pq.Error); ok && err.Code.Class() != ErrorUniqueViolation {
		logger.Error.Println(eerr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		logger.Error.Println(err.Constraint)
		logger.Error.Println(err.Code.Class())
		logger.Error.Println(err.Code.Name())
		return
	} else if ok && err.Code.Class() == ErrorUniqueViolation && len(err.Constraint) == 0 {

		myJSON := `{"` + ErrorCantFindSlug + slugUrl + ` or ` + ErrorCantFindUser + nickname + `"}`

		w.WriteHeader(http.StatusNotFound)
		_, _err := w.Write([]byte(myJSON))
		if _err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(_err.Error())
			return
		}
		return
	} else if ok && err.Code.Class() == ErrorUniqueViolation && err.Constraint == ThreadPrimaryKey {

		logger.Debug.Print(err.Error())

		row := models.DB.DatBase.QueryRow(`SELECT * FROM public."thread" WHERE slug = $1`, slugBody)

		_err := row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title, &thread.Message, &thread.Votes, &thread.Created)
		if _err != nil && _err.Error() != ErrorSqlNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error.Println(_err.Error())
			return
		}

		data, _err := json.Marshal(thread)
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

	data, err := json.Marshal(thread)
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
