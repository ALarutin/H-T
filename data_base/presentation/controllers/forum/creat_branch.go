package forum

import (
	"data_base/models"
	"data_base/presentation/logger"
	"net/http"
)

func CreatBranchHandler(w http.ResponseWriter, r *http.Request) {

	handler := models.Handler{
		Info:   "Handler for creating branch.",
		Name:   "forum_CreatBranch",
		Path:   "/{slug}/create",
		Method: http.MethodPost,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
