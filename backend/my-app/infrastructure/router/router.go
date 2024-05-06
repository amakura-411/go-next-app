package router

import (
	"fmt"
	"my-app/infrastructure/database"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func InitRouting(e *echo.Echo, db *gorm.DB) error {
	fmt.Println("InitRouting!!")
	userRepository := database.NewUserRepository(db)

	userGroup := e.Group("/users")
	userGroup.GET("", func(c echo.Context) error {
		users, err := userRepository.GetAllUsers()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	})

	return nil
}
