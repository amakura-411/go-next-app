package server

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

// 8080ポートにアクセスがあったら、serverHandlerを呼び出す
func serverHandler(w http.ResponseWriter, r *http.Request) {
	//Fprintfは、第一引数にio.Writerを受け取るので、ResponseWriterを渡す
	fmt.Fprintf(w, "Hello Golang World! from serverHandler ok")
}

func StartServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error{
		// serverHandler;
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

// func StopServer() {

// 	fmt.Println("Stop Server")
// }
