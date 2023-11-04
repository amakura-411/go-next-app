package server

import (
	"fmt"
	"net/http"
)

func serverHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! from serverHandler ❤️")
}

func StartServer() {
	http.HandleFunc("/", serverHandler)
	http.ListenAndServe(":8080", nil)
}

func StopServer() {

	fmt.Println("Stop Server")
}
