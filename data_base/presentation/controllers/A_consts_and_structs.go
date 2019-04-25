package controllers

const (
	ErrorSqlNoRows           = `sql: no rows in result set`
	ErrorCantFindUser        = `message": "cant find user with nickname `
	ErrorCantFindThread      = `message": "cant find thread with slug `
	ErrorCantFindSlug        = `message": "cant find forum with slug `
	ErrorCantFindParent      = `message": "cant find parent with id `
	ErrorNameNoNullViolation = `not_null_violation`
)

const(
	ErrorUniqueViolation     = `23`
)

const (
	ForumUserForeignKey    = `forum_user_fk`
	ThreadForumForeignKey  = `thread_forum_fk`
	ThreadAuthorForeignKey = `thread_author_fk`
	ForumPrimaryKey        = `forum_pk`
	ThreadPrimaryKey       = `thread_pk`
)
