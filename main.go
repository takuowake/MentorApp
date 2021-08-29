package main

import (
	"MentorApp/server"
)

func main() {
r := server.GetRouter()
r.Run(":8080")
}