package models

import "github.com/lib/pq"

func (db *dbManager) UpdatePost(message string, id int) (post Post, err error) {

	//row := db.dataBase.QueryRow(
	//	`UPDATE public."post"
	//	SET message = $1
	//	WHERE id = $2 RETURNING id, author, thread, forum, message, is_edited, parent, created`,
	//	message, id)
	row := db.dataBase.QueryRow(
		`SELECT * FROM func_update_post($1::text, $2::INT)`,
		message, id)
	err = row.Scan(&post.ID, &post.Author, &post.Thread, &post.Forum,
		&post.Message, &post.IsEdited, &post.Parent, &post.Created, pq.Array(&post.Path))
	return
}
