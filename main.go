package main

import (
	"GinFrameworkPractice/server"
	"fmt"
)

func main() {
	fmt.Println("GinFrameworkPractice")
	server := server.SetUpServer()
	server.Run(":8080")
}
