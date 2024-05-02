package server

import (
	"fmt"
	"go-next-app/backend/my-app/infrastructure/router"
	"log"

	"github.com/labstack/echo"
)

func StartServer() {
	fmt.Println("Start Server!!")

	e := echo.New()
	if e == nil {
		log.Fatal("echo.New() error")
	}

	err := router.InitRouting(e)
	if err != nil {
		log.Fatal(err)
	}

	if err := e.Start(":1323"); err != nil {
		log.Fatal(err)
	}

}
