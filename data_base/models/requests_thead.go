package models

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

func (db *dbManager) GetThread (slug string, threadId int) (thread Thread, err error){
	row := db.dataBase.QueryRow(
		`SELECT * FROM public."thread" WHERE slug = $1 OR id = $2`,
		slug, threadId)
	err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum,
		&thread.Title, &thread.Message, &thread.Votes, &thread.Created,)
	return
}