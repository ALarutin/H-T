package post

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func ChangeMessageHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for changing the message.",
		Name:   "post_ChangeMessage",
		Path:   "/{id}/details",
		Method: http.MethodPost,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
