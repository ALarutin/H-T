package service

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func GetDataBaseInfoHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for obtaining information about the database.",
		Name:   "service_GetDataBaseInfo",
		Path:   "/status",
		Method: http.MethodGet,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
