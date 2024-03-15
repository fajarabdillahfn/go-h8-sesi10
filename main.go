package main

import (
	"fmt"
	"sesi10/router"
)

func main() {
	r := router.StartApp()

	fmt.Println("service start at :3333")
	r.Run(":3333")
}
