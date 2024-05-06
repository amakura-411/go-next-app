package main

import (
	"fmt"
	"go-next-app/backend/my-app/infrastructure/database"
	"go-next-app/backend/my-app/server"
)

// hello World!!をコンソールに表示する
func main() {
	fmt.Println("Hello World!!")

	server.StartServer()

	// db接続テスト
	_, err := database.InitConnection()
	if err != nil {
		fmt.Println("DB接続エラー")
		return
	}
	fmt.Println("DB接続成功")
}
