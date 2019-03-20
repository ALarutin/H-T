package forum

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func GetThreadsHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for getting a list of forum discussion branches.",
		Name:   "forum_GetThreads",
		Path:   "/{slug}/threads",
		Method: http.MethodGet,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
