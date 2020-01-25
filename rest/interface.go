package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hopperteam/hopper-monitor/storage"
	"net/http"
	"time"
)

func getRootRouter(cfg *storage.ConfigProvider) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/log", getLogRouter(cfg))

	return r
}

func writeJson(obj interface{}, w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ListenAndServe(cfg *storage.ConfigProvider) error {
	srv := &http.Server{
		Handler: getRootRouter(cfg),
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
