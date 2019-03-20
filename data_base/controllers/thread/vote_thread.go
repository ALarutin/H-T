package thread

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func VoteThreadHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for voting the discussion thread.",
		Name:   "thread_VoteThread",
		Path:   "/{slug_or_id}/vote",
		Method: http.MethodPost,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
