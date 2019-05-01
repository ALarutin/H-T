package models

import "github.com/lib/pq"

func (db *dbManager) CreatePost(post Post, id int, forum string) (p Post, err error) {

	//row := db.dataBase.QueryRow(
	//	`INSERT INTO public."post" (author, thread, forum, message, parent)
	//	VALUES ($1, (SELECT id FROM public."thread" WHERE slug = $2 OR id = $3), (SELECT forum FROM public."thread" WHERE slug = $2 OR id = $3), $4, $5)
	//	RETURNING id, created, thread, forum`,
	//	post.Author, slug, threadId, post.Message, post.Parent)

	row := db.dataBase.QueryRow(`SELECT * FROM func_create_post($1::citext, $2::INT, $3::text, $4::INT, $5::citext)`,
		post.Author, id, post.Message,  post.Parent, forum)
	err = row.Scan(&p.ID, &p.Author, &p.Thread, &p.Forum,
		&p.Message, &p.IsEdited, &p.Parent, &p.Created, pq.Array(&p.Path))
	return
}

func (db *dbManager) GetThread(slug string, threadId int) (thread Thread, err error) {
	row := db.dataBase.QueryRow(`SELECT * FROM func_get_thread($1::citext, $2::INT)`, slug, threadId)
	err = row.Scan(&thread.IsNew, &thread.ID, &thread.Slug, &thread.Author, &thread.Forum,
		&thread.Title, &thread.Message, &thread.Votes, &thread.Created)
	return
}

func (db *dbManager) UpdateThread(message string, title string, slug string, threadId int) (thread Thread, err error) {
	row := db.dataBase.QueryRow(`SELECT * FROM func_update_thread($1::text, $2::text, $3::citext, $4::INT)`,
		message, title, slug, threadId)
	err = row.Scan(&thread.IsNew, &thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title,
		&thread.Message, &thread.Votes, &thread.Created)
	return
}

func (db *dbManager) CreateOrUpdateVote(vote Vote, slug string, threadId int) (thread Thread, err error) {
	row := db.dataBase.QueryRow(`SELECT * FROM func_create_or_update_vote($1::citext, $2::citext, $3::INT, $4::INT)`,
		vote.Nickname, slug, threadId, vote.Voice)
	err = row.Scan(&thread.IsNew, &thread.ID, &thread.Slug, &thread.Author, &thread.Forum, &thread.Title,
		&thread.Message, &thread.Votes, &thread.Created)
	return
}
