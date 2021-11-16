package main

import (
	"presence-app-backend/configs"
	"presence-app-backend/routes"
)

func main() {
	configs.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
