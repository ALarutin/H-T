package thread

import (
	"data_base/models"
	"data_base/presentation/logger"
	"net/http"
)

func CreatNewPostHandler(w http.ResponseWriter, r *http.Request) {
	handler := models.Handler{
		Info:   "Handler for creating new post.",
		Name:   "thread_CreatNewPost",
		Path:   "/{slug_or_id}/create",
		Method: http.MethodPost,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
