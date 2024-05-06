package router

import (
	"fmt"
	"my-app/domain/entity"
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
		fmt.Println("/usersへのアクセス")
		users, err := userRepository.GetAllUsers()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	})

	userGroup.GET("/:id", func(c echo.Context) error {
		fmt.Println("/users/:idへのアクセス")
		user, err := userRepository.GetUserByID(c.Param("id"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	})

	userGroup.POST("", func(c echo.Context) error {
		fmt.Println("/usersへのPOSTアクセス")

		// JSONで受け取ったデータを格納するための構造体
		var newUser entity.User

		// リクエストボディからデータを取得
		if err := c.Bind(&newUser); err != nil {
			return err
		}

		// 一応、受け取ったデータを表示
		fmt.Println(newUser)

		// ここでentity.Userを作成して、それを引数に渡す
		user, err := entity.NewUser(newUser.Username, newUser.Password, &newUser.Icon, &newUser.Created_at)
		if err != nil {
			return err
		}
		fmt.Println(user)

		err = userRepository.CreateUser(user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	})

	return nil
}
