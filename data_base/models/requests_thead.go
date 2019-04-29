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
