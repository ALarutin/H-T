package thread

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func UpdateBranchHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for updating the branch.",
		Name:   "thread_UpdateBranch",
		Path:   "/{slug_or_id}/details",
		Method: http.MethodPost,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
