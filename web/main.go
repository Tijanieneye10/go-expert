package web

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	output := fmt.Sprintf("%s", "<h1>Hello World</h1>")
	_, err := w.Write([]byte("<h1>Hello World</h1>"))
	if err != nil {
		return
	}
}
