package forum

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func GetForumInfoHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for obtaining information about the forum.",
		Name:   "forum_GetForumInfo",
		Path:   "/{slug}/details",
		Method: http.MethodGet,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
