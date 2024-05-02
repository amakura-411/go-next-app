package server

import (
	"fmt"
	"go-next-app/backend/my-app/infrastructure/router"
)

func StartServer() {
	fmt.Println("Start Server!!")
	//ルーティングの設定
	router.InitRouting()

}
