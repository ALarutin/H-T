package user

import (
	"data_base/core/bootstrap"
	"data_base/logger"
	"net/http"
)

func CreatNewUserHandler(w http.ResponseWriter, r *http.Request) {
	handler := bootstrap.Handler{
		Info:   "Handler for creating new user.",
		Name:   "user_CreatNewUser",
		Path:   "{nickname}/create",
		Method: http.MethodPost,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
