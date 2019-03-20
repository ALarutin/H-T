package post

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func GetThreadInfoHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for getting information about the discussion thread.",
		Name:   "post_GetThreadInfo",
		Path:   "/{id}/details",
		Method: http.MethodGet,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
