package models

import "data_base/presentation/logger"

func (db *dbManager) UpdatePost(message string, id int) (post Post, err error) {

	logger.Error.Print(message)

	row := db.dataBase.QueryRow(
		`UPDATE public."post"
		SET message = $1
		WHERE id = $2 RETURNING id, author, thread, forum, message, is_edited, parent, created`,
		message, id)
	err = row.Scan(&post.ID, &post.Author, &post.Thread, &post.Forum,
		&post.Message, &post.IsEdited, &post.Parent, &post.Created)
	return
}
