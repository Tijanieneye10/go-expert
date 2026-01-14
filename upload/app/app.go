package app

import (
	"log"
	"net/http"
)

type App struct {
	Log *log.Logger
}

func (app *App) Serve(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello World"}`))
	})
}
