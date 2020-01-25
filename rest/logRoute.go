package rest

import (
	"github.com/gorilla/mux"
	"github.com/hopperteam/hopper-monitor/storage"
	"github.com/hopperteam/hopper-monitor/types"
	"net/http"
)

type logHandler struct {
	config *storage.ConfigProvider
}

func (handler *logHandler) handleGetLog(w http.ResponseWriter, r *http.Request) {
	filter := types.LogFilter{}
	writeJson(handler.config.LogStorage.GetLogEntries(&filter), w, r)
}


func getLogRouter(cfg *storage.ConfigProvider) *mux.Router {
	r := mux.NewRouter()
	l := logHandler{cfg}

	r.Methods("GET").HandlerFunc(l.handleGetLog)

	return r
}


