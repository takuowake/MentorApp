package main

import (
	"MentorApp/server"
)

func main() {
	router := server.GetRouter()

	router.Run(":8080")
}
