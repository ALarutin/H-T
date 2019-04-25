package models

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