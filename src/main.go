package main

import (
	"FiberPlayground/src/routes"
	"log"
)

func main() {
	app := routes.SetupRoutes()

	log.Fatal(app.Listen(":3000"))
}
