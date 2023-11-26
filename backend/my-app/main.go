package main


import (
	"fmt"
	"go-next-app/backend/my-app/server"
)




// hello World!!をコンソールに表示する
func main() {
	fmt.Println("Hello World!!")

	server.StartServer()
}

