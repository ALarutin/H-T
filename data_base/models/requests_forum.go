package models

import (
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

	row := db.dataBase.QueryRow(
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
			since, slug) + fmt.Sprintf(`%s `, desc) + fmt.Sprintf(`LIMIT %s`, limit))
	if err != nil {
		return
	}

	var thread Thread
	for rows.Next() {
		err = rows.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title, &thread.Message, &thread.Votes, &thread.Created)
		if err != nil {
			return
		}
		threads = append(threads, thread)
	}
	return
}

func (db *dbManager) GetUsers(slug string, since string, desc string, limit string) (users []User, err error) {

	//tempSelect := fmt.Sprintf(
	//	`SELECT * FROM public."person"
	//		WHERE nickname in (SELECT user_nickname
	//		FROM (SELECT * FROM public."forum_users" GROUP BY forum_slug, user_nickname) as m
	//		WHERE forum_slug = '%s')`,
	//	slug)
	tempSelect := fmt.Sprintf(
		`SELECT * FROM public."person" 
		WHERE nickname in (SELECT user_nickname
		FROM  public."forum_users" 
		WHERE forum_slug = '%s')`,
		slug)
	rows, err := db.dataBase.Query(
		fmt.Sprintf(
			`SELECT nickname, email, fullname, about 
			FROM (%s) as p WHERE id >= '%s' ORDER BY id`,
			tempSelect, since) + fmt.Sprintf(`%s `, desc) + fmt.Sprintf(`LIMIT %s`, limit))
	if err != nil {
		return
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.Nickname, &user.Email, &user.Fullname, &user.About)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}
