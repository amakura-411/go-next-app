package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo) error {
	fmt.Println("InitRouting!!")

	// hello
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// user
	gu := e.Group("/user")
	gu.GET("", func(c echo.Context) error {
		fmt.Println("GET /user")
		return c.String(http.StatusOK, "GET /user")
	})
	gu.POST("", func(c echo.Context) error {
		fmt.Println("POST /user")
		return c.String(http.StatusOK, "POST /user")
	})
	gu.PUT("", func(c echo.Context) error {
		fmt.Println("PUT /user")
		return c.String(http.StatusOK, "PUT /user")
	})
	gu.DELETE("", func(c echo.Context) error {
		fmt.Println("DELETE /user")
		return c.String(http.StatusOK, "DELETE /user")
	})

	return nil
}
