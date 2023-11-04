package main

import (
	"fmt"
	"go-next-app/backend/server"
)

// hello World!!をコンソールに表示する
func main() {
	fmt.Println("Hello World!!")
	server.StartServer()

}