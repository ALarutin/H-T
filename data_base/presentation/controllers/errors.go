package controllers

const (
	messageCantFind          = `message": "cant find `
	cantFindUser             = `user with nickname `
	cantFindThread           = `thread with slug `
	cantFindForum            = `forum with slug `
	cantFindParent           = `parent with id `
	emailUsed                = ` has already taken by another user`
	ErrorNameNoNullViolation = `not_null_violation`
)

const (
	errorUniqueViolation = `23`
)

const (
	errorSqlNoRows         = `sql: no rows in result set`
	forumUserForeignKey    = `forum_user_fk`
	threadForumForeignKey  = `thread_forum_fk`
	threadAuthorForeignKey = `thread_author_fk`
	forumPrimaryKey        = `forum_pk`
	threadPrimaryKey       = `thread_pk`
)
