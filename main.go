package main

import (
	"fmt"
	"os"
	"sesi10/database"
	"sesi10/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()

	port := os.Getenv("PORT")
	fmt.Println("service start at :" + port)
	r.Run(":" + port)
}
