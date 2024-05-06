package server

import (
	"fmt"
	"log"
	"my-app/infrastructure/router"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {

	fmt.Println("Start Server!!")
	e := echo.New()
	if e == nil {
		log.Fatal("echo.New() error")
	}

	err := router.InitRouting(e, db)
	if err != nil {
		log.Fatal(err)
	}

	if err := e.Start(":1323"); err != nil {
		log.Fatal(err)
	}

}
