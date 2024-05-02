package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting() (bool, error) {
	fmt.Println("InitRouting!!")
	e := echo.New()

	// ルーティングの設定
	e.GET("/hello", func(c echo.Context) error {
		fmt.Println("Hello World!!")
		return c.String(http.StatusOK, "Hello, World!")
	})

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	// Routes

	return true, nil
}
