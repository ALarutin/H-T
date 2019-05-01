package controllers

const (
	messageCantFind = `message": "cant find `
	cantFindUser    = `user with nickname `
	cantFindThread  = `thread with slug or id `
	cantFindForum   = `forum with slug `
	cantFindParentOrUser  = `parent or author`
	cantFindPost    = `post with id `
	emailUsed       = ` has already taken by another user`
)

const (
	errorUniqueViolation       = `pq: unique_violation`
	errorPqNoDataFound            = `pq: no_data_found`
	errorForeignKeyViolation   = `pq: foreign_key_violation`
)

const (
	forumUserForeignKey    = `forum_user_fk`
	threadForumForeignKey  = `thread_forum_fk`
	threadAuthorForeignKey = `thread_author_fk`

	postAuthorForeignKey   = `post_author_fk`
	voteThreadForeignKey   = `vote_thread_fk`
	forumPrimaryKey        = `forum_pk`
	threadPrimaryKey       = `thread_pk`
)
