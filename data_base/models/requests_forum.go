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