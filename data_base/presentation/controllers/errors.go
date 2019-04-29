package controllers

const (
	messageCantFind = `message": "cant find `
	cantFindUser    = `user with nickname `
	cantFindThread  = `thread with slug or id `
	cantFindForum   = `forum with slug `
	cantFindParent  = `parent with id `
	emailUsed       = ` has already taken by another user`
)

const (
	errorUniqueViolation       = `23`
	errorNameNotNullConstraint = `not-null constraint`
	errorNameNotNullViolation  = `not_null_violation`
	errorSqlNoRows             = `sql: no rows in result set`
)

const (
	forumUserForeignKey     = `forum_user_fk`
	threadForumForeignKey   = `thread_forum_fk`
	threadAuthorForeignKey  = `thread_author_fk`
	postParentForeignKey = `post_parent_fk`
	postAuthorForeignKey = `post_author_fk`
	voteThreadForeignKey = `vote_thread_fk`
	forumPrimaryKey         = `forum_pk`
	threadPrimaryKey        = `thread_pk`
)
