package models

import (
	"data_base/presentation/logger"
	"errors"
	"fmt"
)

func (db *dbManager) CreateForum(forum Forum) (err error) {

	_, err = db.dataBase.Exec(
		`INSERT INTO public."forum" (author, slug, title,  posts, threads)
		VALUES ($1, $2, $3, $4, $5)`,
		forum.User, forum.Slug, forum.Title, forum.Posts, forum.Threads)
	return
}

func (db *dbManager) SelectForum(slug string) (forum Forum, err error) {

	row :=db.dataBase.QueryRow(
		`SELECT * FROM public."forum" 
		WHERE slug = $1`,
		slug)
	err = row.Scan(&forum.Slug, &forum.User, &forum.Title, &forum.Posts, &forum.Threads)
	return
}

func (db *dbManager) CreateThread(thread Thread) (err error) {
	_, err = db.dataBase.Exec(
		`INSERT INTO public."thread" (author, created, forum, message, slug, title, votes)
		  VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		  thread.Author, thread.Created, thread.Forum, thread.Message, thread.Slug, thread.Title, thread.Votes)
	return
}

func (db *dbManager) SelectThread(slug string) (thread Thread, err error) {

	row := db.dataBase.QueryRow(
		`SELECT * FROM public."thread" 
			WHERE slug = $1`,
			slug)
	err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title, &thread.Message, &thread.Votes, &thread.Created)
	return
}

func (db *dbManager) GetForum(slug string) (forum Forum, err error) {

	row := db.dataBase.QueryRow(
		`SELECT * FROM public."forum" 
			WHERE slug = $1`,
			slug)
	err = row.Scan(&forum.Slug, &forum.User, &forum.Title, &forum.Posts, &forum.Threads)
	return
}

func (db *dbManager) GetThreads(slug string, since string, desc string, limit string) (threads []Thread, err error) {

	rows, err := db.dataBase.Query(
		fmt.Sprintf(
			`SELECT * FROM public."thread" 
			WHERE created >= '%s' AND forum = '%s' ORDER BY created `,
			since, slug) +
		fmt.Sprintf(`%s `,
			desc) +
		fmt.Sprintf(
			`LIMIT %s`,
			limit))
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	if !rows.Next(){
		err = errors.New(ErrorSqlNoRows)
	}

	var thread Thread
	for rows.Next() {
		err = rows.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title, &thread.Message, &thread.Votes, &thread.Created)
		if err != nil {
			logger.Error.Println(err.Error())
			return
		}
		threads = append(threads, thread)
	}
	return
}