package main

import "github.com/sundargautam18/Go-with-go/routes"

func main() {
	r := routes.InitRoutes()
	r.Run(":8080") // Start server
}
