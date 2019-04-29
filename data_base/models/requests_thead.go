package models

import "data_base/presentation/logger"

func (db *dbManager) CreatePost(post Post, slug string, threadId int) (p Post, err error) {

	err = db.dataBase.QueryRow(
		`INSERT INTO public."post" (author, thread, forum, message, parent)
		VALUES ($1, (SELECT id FROM public."thread" WHERE slug = $2 OR id = $3), (SELECT forum FROM public."thread" WHERE slug = $2 OR id = $3), $4, $5) 
		RETURNING id, created, thread, forum`,
		post.Author, slug, threadId, post.Message, post.Parent).
		Scan(&p.ID, &p.Created, &p.Thread, &p.Forum)
	p.Author = post.Author
	p.Parent = post.Parent
	p.Message = post.Message
	return
}

func (db *dbManager) GetThread(slug string, threadId int) (thread Thread, err error) {
	row := db.dataBase.QueryRow(
		`SELECT * FROM public."thread" WHERE slug = $1 OR id = $2`,
		slug, threadId)
	err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum,
		&thread.Title, &thread.Message, &thread.Votes, &thread.Created)
	return
}

func (db *dbManager) UpdateThread(message string, title string, slug string, threadId int) (thread Thread, err error) {
	row := db.dataBase.QueryRow(
		`UPDATE public."thread"
				SET message = $1, title = $2
				WHERE slug = $3 OR id = $4 RETURNING *`,
				message, title, slug, threadId)
	err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum,
		&thread.Title, &thread.Message, &thread.Votes, &thread.Created)
	return
}

func (db *dbManager) CreateOrUpdateVote(vote Vote) (err error) {
	logger.Error.Print(vote.ThreadSlug)
	_, err = db.dataBase.Exec(
		`INSERT INTO public."vote" (thread_slug, user_nickname, voice)
			VALUES ($1, $2, $3)
			ON CONFLICT ON CONSTRAINT vote_pk DO UPDATE
			  SET voice = $3
			  WHERE vote.thread_slug = $1 AND vote.user_nickname = $2`,
		vote.ThreadSlug, vote.Nickname, vote.Voice)
	return
}
