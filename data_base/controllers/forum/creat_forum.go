package forum

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func CreatForumHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for creating forum.",
		Name:   "forum_CreatForum",
		Path:   "/create",
		Method: http.MethodPost,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
