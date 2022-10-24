package main

import (
	"final-project-gin/database"
	"final-project-gin/router"
	"fmt"
)

func main() {
	fmt.Println()
	database.StartDB()
	r := router.StartApp()
	r.Run(":9898")

}
