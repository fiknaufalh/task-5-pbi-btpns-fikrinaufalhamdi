package main

import "profile-picture-api/router"


func main() {
	r := router.SetupRouter()
	r.Run(":5432")
}
