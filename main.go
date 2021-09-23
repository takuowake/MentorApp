package main

import (
	"MentorApp/server"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	router := server.GetRouter()

	router.Run(":8080")
}
