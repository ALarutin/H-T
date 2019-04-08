package controllers

const (
	ErrorSqlNoRows           = `sql: no rows in result set`
	ErrorCantFindUser        = `message": "cant find user with nickname `
	ErrorCantFindThread      = `message": "cant find thread with slug `
	ErrorCantFindSlug        = `message": "cant find forum with slug `
	ErrorHaveDuplicates      = `pq: duplicate key value violates unique constraint "forum_users_pk"`
	ErrorCantFindParent      = `message": "cant find parent with id `
	ErrorUniqueViolation     = `23`
	ErrorNameNoNullViolation = `not_null_violation`
)

const (
	ForumUserForeignKey    = `forum_user_fk`
	ThreadForumForeignKey  = `thread_forum_fk`
	ThreadAuthorForeignKey = `thread_author_fk`
	ForumPrimaryKey        = `forum_pk`
	ThreadPrimaryKey       = `thread_pk`
)
