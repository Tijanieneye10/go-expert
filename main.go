package main

import (
	"encoding/json"
	"net/http"
)

type DefaultHandler struct{}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Hello World",
		"data": map[string]interface{}{
			"foo": "bar",
		},
	})

	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(map[string]string{"greet": "hello world"})
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	handler := &DefaultHandler{}

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
