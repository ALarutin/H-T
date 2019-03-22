package thread

import (
	"data_base/models"
	"data_base/presentation/logger"
	"net/http"
)

func GetThreadInfoHandler(w http.ResponseWriter, r *http.Request) {
	handler := models.Handler{
		Info:   "Handler for getting information about the discussion thread.",
		Name:   "thread_GetThreadInfo",
		Path:   "/{slug_or_id}/details",
		Method: http.MethodGet,
	}
	logger.Info.Printf("Get into handler\nname:    %s\ninfo:    %s\nmetod:   %s\n",
		handler.Name, handler.Info, handler.Method)
}
