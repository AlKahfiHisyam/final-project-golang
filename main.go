package main

import (
	"final-project-golang/database"
	"final-project-golang/routers"
)

func main() {
	database.StartDB()

	r := routers.StartApp()
	r.Run(":8080")
}
