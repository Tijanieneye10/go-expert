package main

import (
	"go-expert/upload/app"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	app := &app.App{
		Log: log.New(os.Stdout, "", log.LstdFlags),
	}

	app.Serve(mux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
