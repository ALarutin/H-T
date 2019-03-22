package forum

import (
	"data_base/models"
	"data_base/presentation/logger"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	handler := models.Handler{
		Info:   "Handler for obtaining the users of this forum.",
		Name:   "forum_GetUsers",
		Path:   "/{slug}/users",
		Method: http.MethodGet,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
