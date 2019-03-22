package user

import (
	"data_base/models"
	"data_base/presentation/logger"
	"net/http"
)

func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	handler := models.Handler{
		Info:   "Handler for getting information about user.",
		Name:   "user_GetUserInfo",
		Path:   "/{nickname}/profile",
		Method: http.MethodGet,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
