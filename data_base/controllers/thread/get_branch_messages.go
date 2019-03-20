package thread

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func GetBranchMessagesHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for getting messages of this branch of the discussion.",
		Name:   "thread_GetBranchMessages",
		Path:   "/{slug_or_id}/posts",
		Method: http.MethodGet,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}