package server

import (
	"fmt"
	"net/http"
)

//8080ポートにアクセスがあったら、serverHandlerを呼び出す
func serverHandler(w http.ResponseWriter, r *http.Request) {
	//Fprintfは、第一引数にio.Writerを受け取るので、ResponseWriterを渡す
	fmt.Fprintf(w, "Hello World! from serverHandler ❤️")
}

func StartServer() {
	// localhost:8080にアクセスがあったら、serverHandlerを呼び出す
	http.HandleFunc("/", serverHandler)
	// 8080ポートでサーバーを立ち上げる
	http.ListenAndServe(":8080", nil)
}

// func StopServer() {

// 	fmt.Println("Stop Server")
// }
