package router

import (
	"data_base/models"
	"data_base/presentation/controllers/forum"
	"data_base/presentation/controllers/post"
	"data_base/presentation/controllers/service"
	"data_base/presentation/controllers/thread"
	"data_base/presentation/controllers/user"
	"data_base/presentation/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter() (router *mux.Router) {

	router = mux.NewRouter()

	//router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){})

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//forum

	forumSubRouter := router.PathPrefix("/forum").Subrouter()

	_forum := []models.Route{
		{
			Info:    "Handler for creating forum.",
			Name:    "forum_CreatForum",
			Path:    "/create",
			Method:  http.MethodPost,
			Handler: forum.CreatForumHandler,
		},
		{
			Info:    "Handler for creating branch.",
			Name:    "forum_CreatBranch",
			Path:    "/{slug}/create",
			Method:  http.MethodPost,
			Handler: forum.CreatBranchHandler,
		},
		{
			Info:    "Handler for obtaining information about the forum.",
			Name:    "forum_GetForumInfo",
			Path:    "/{slug}/details",
			Method:  http.MethodGet,
			Handler: forum.GetForumInfoHandler,
		},
		{
			Info:    "Handler for getting a list of forum discussion branches.",
			Name:    "forum_GetThreads",
			Path:    "/{slug}/threads",
			Method:  http.MethodGet,
			Handler: forum.GetThreadsHandler,
		},
		{
			Info:    "Handler for obtaining the users of this forum.",
			Name:    "forum_GetUsers",
			Path:    "/{slug}/users",
			Method:  http.MethodGet,
			Handler: forum.GetUsersHandler,
		},
	}

	for _, r := range _forum {
		forumSubRouter.
			HandleFunc(r.Path, r.Handler).
			Methods(r.Method).
			Name(r.Name)
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//post

	postSubRouter := router.PathPrefix("/post").Subrouter()

	_post := []models.Route{
		{
			Info:    "Handler for changing the message.",
			Name:    "post_ChangeMessage",
			Path:    "/{id}/details",
			Method:  http.MethodPost,
			Handler: post.ChangeMessageHandler,
		},
		{
			Info:    "Handler for getting information about the discussion thread.",
			Name:    "post_GetThreadInfo",
			Path:    "/{id}/details",
			Method:  http.MethodGet,
			Handler: post.GetThreadInfoHandler,
		},
	}

	for _, r := range _post {
		postSubRouter.
			Methods(r.Method).
			Name(r.Name).
			Path(r.Path).
			HandlerFunc(r.Handler)
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//service

	serviceSubRouter := router.PathPrefix("/service").Subrouter()

	_service := []models.Route{
		{
			Info:    "Handler for clearing all data in the database.",
			Name:    "service_ClearDataBase",
			Path:    "/clear",
			Method:  http.MethodPost,
			Handler: service.ClearDataBaseHandler,
		},
		{
			Info:    "Handler for obtaining information about the database.",
			Name:    "service_GetDataBaseInfo",
			Path:    "/status",
			Method:  http.MethodGet,
			Handler: service.GetDataBaseInfoHandler,
		},
	}

	for _, r := range _service {
		serviceSubRouter.
			Methods(r.Method).
			Name(r.Name).
			Path(r.Path).
			HandlerFunc(r.Handler)
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//thread

	threadSubRouter := router.PathPrefix("/thread").Subrouter()

	_thread := []models.Route{
		{
			Info:    "Handler for creating new post.",
			Name:    "thread_CreatNewPost",
			Path:    "/{slug_or_id}/create",
			Method:  http.MethodPost,
			Handler: thread.CreatNewPostHandler,
		},
		{
			Info:    "Handler for updating the branch.",
			Name:    "thread_UpdateBranch",
			Path:    "/{slug_or_id}/details",
			Method:  http.MethodPost,
			Handler: thread.UpdateBranchHandler,
		},
		{
			Info:    "Handler for voting the discussion thread.",
			Name:    "thread_VoteThread",
			Path:    "/{slug_or_id}/vote",
			Method:  http.MethodPost,
			Handler: thread.VoteThreadHandler,
		},
		{
			Info:    "Handler for getting information about the discussion thread.",
			Name:    "thread_GetThreadInfo",
			Path:    "/{slug_or_id}/details",
			Method:  http.MethodGet,
			Handler: thread.GetThreadInfoHandler,
		},
		{
			Info:    "Handler for getting messages of this branch of the discussion.",
			Name:    "thread_GetBranchMessages",
			Path:    "/{slug_or_id}/posts",
			Method:  http.MethodGet,
			Handler: thread.GetBranchMessagesHandler,
		},
	}

	for _, r := range _thread {
		threadSubRouter.
			Methods(r.Method).
			Name(r.Name).
			Path(r.Path).
			HandlerFunc(r.Handler)
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//user

	userSubRouter := router.PathPrefix("/user").Subrouter()

	_user := []models.Route{
		{
			Info:    "Handler for creating new user.",
			Name:    "user_CreatNewUser",
			Path:    "{nickname}/create",
			Method:  http.MethodPost,
			Handler: user.CreatNewUserHandler,
		},
		{
			Info:    "Handler for changing user data.",
			Name:    "user_ChangUserData",
			Path:    "{nickname}/profile",
			Method:  http.MethodPost,
			Handler: user.ChangUserDataHandler,
		},
		{
			Info:    "Handler for getting information about user.",
			Name:    "user_GetUserInfo",
			Path:    "/{nickname}/profile",
			Method:  http.MethodGet,
			Handler: user.GetUserInfoHandler,
		},
	}

	for _, r := range _user {
		userSubRouter.
			Methods(r.Method).
			Name(r.Name).
			Path(r.Path).
			HandlerFunc(r.Handler)
	}

	router.Use(middlewares.MiddlewareLogger)
	router.Use(middlewares.MiddlewarePanic)
	return
}
