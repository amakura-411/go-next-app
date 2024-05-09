package router

import (
	"fmt"
	userEntity "my-app/domain/entity/user"
	userRepository "my-app/infrastructure/database/user"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func InitRouting(e *echo.Echo, db *gorm.DB) error {
	fmt.Println("InitRouting!!")

	// ========================================
	// USERS
	userGroup := e.Group("/users")
	userRepository := userRepository.NewUserRepository(db)

	userGroup.GET("", func(c echo.Context) error {
		users, err := userRepository.GetAllUsers()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	})

	userGroup.GET("/:id", func(c echo.Context) error {
		user, err := userRepository.GetUserByID(c.Param("id"))
		if err != nil {
			fmt.Println("error:userRepository.GetUserByID", err)
			return err
		}
		return c.JSON(http.StatusOK, user)
	})

	userGroup.POST("", func(c echo.Context) error {

		// JSONで受け取ったデータを格納するための構造体
		var newUser userEntity.User

		// リクエストボディからデータを取得
		if err := c.Bind(&newUser); err != nil {
			return err
		}

		// ここでentity.Userを作成して、それを引数に渡す
		user, err := userEntity.NewUser(newUser.Username, newUser.Password, &newUser.Icon, &newUser.Created_at)
		if err != nil {
			return err
		}

		// uuidの生成
		existFlag := true
		for i := 0; i < 5; i++ {
			id := userEntity.GenerateUserID()
			fmt.Println("i:", i, id)
			existFlag = userRepository.CheckUserID(id)
			fmt.Println("existFlag:", existFlag)
			if existFlag == true {
				user.User_id = id
				break
			}
		}
		if user.User_id == "" {
			return fmt.Errorf("failed to generate user_id")
		}

		err = userRepository.CreateUser(user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	})

	// userGroup.PUT("/:id", func(c echo.Context) error {

	// 	var updateUser userEntity.User

	// 	if err := c.Bind(&updateUser); err != nil {
	// 		return err
	// 	}
	// 	user, err := userEntity.UpdateUser(updateUser.Username, updateUser.Icon)
	// 	if err != nil {
	// 		fmt.Println("error:entity.UpdateUser", err)
	// 		return err
	// 	}

	// 	err = userRepository.UpdateUser(c.Param("id"), &user)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return c.JSON(http.StatusOK, user)
	// })

	return nil
}
