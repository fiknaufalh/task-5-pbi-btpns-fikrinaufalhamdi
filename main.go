package main

import (
	// "fmt"
	"profile-picture-api/database"
	"profile-picture-api/router"
)

func init() {
	database.ConnectDatabase()
}

func main() {
	r := router.SetupRouter()
	r.Run(":5000")
}
