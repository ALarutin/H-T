package models

import (
	"data_base/presentation/logger"
	"github.com/lib/pq"
)

func (db *dbManager) CreatePost(authors []string, messages []string, parents []int, slug string, threadId int, len int) (ps []Post, err error) {

	//row := db.dataBase.QueryRow(
	//	`INSERT INTO public."post" (author, thread, forum, message, parent)
	//	VALUES ($1, (SELECT id FROM public."thread" WHERE slug = $2 OR id = $3), (SELECT forum FROM public."thread" WHERE slug = $2 OR id = $3), $4, $5)
	//	RETURNING id, created, thread, forum`,
	//	post.Author, slug, threadId, post.Message, post.Parent)
	logger.Info.Print("//////////////////////")
	rows, err := db.dataBase.Query(`SELECT * FROM func_create_posts($1::citext, $2::INT, $3::citext[], $4::text[], $5::INT[], $6::INT)`,
		slug, threadId, pq.Array(authors), pq.Array(messages), pq.Array(parents), len)
	if err != nil {
		return
	}
	defer rows.Close()
	logger.Info.Print("//////////////////////")
	logger.Error.Print(rows)
	var post Post
	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Author, &post.Thread, &post.Forum,
			&post.Message, &post.IsEdited, &post.Parent, &post.Created, &post.Path)
		if err != nil {
			return
		}
		ps = append(ps, post)
	}
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
