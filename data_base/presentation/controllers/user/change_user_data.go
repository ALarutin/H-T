package user

import (
	"data_base/models"
	"data_base/presentation/logger"
	"net/http"
)

func ChangUserDataHandler(w http.ResponseWriter, r *http.Request) {
	handler := models.Handler{
		Info:   "Handler for changing user data.",
		Name:   "user_ChangUserData",
		Path:   "{nickname}/profile",
		Method: http.MethodPost,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
