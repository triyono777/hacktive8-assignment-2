package main

import (
	"assigment-2/database"
	"assigment-2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":9000")
}
