package service

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func ClearDataBaseHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for clearing all data in the database.",
		Name:   "service_ClearDataBase",
		Path:   "/clear",
		Method: http.MethodPost,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
