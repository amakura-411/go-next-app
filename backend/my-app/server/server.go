package server

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error{
		// port:1323の画面に表示させる
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
	fmt.Fprint("port:1323 OPEN")
}
