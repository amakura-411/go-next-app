package main

import (
	"fmt"
	"my-app/infrastructure/database"
	"my-app/server"
)

// hello World!!をコンソールに表示する
func main() {
	fmt.Println("Hello World!!")

	// データベース接続
	db, err := database.InitConnection()

	if err != nil {
		fmt.Println("データベース接続エラー")
		// プログラム終了
		return
	}
	// defer database.CloseConnection(db)
	// サーバー起動
	server.StartServer(db)
}
