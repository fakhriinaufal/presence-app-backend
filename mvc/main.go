package main

import (
	"presence-app-backend/mvc/configs"
	"presence-app-backend/mvc/routes"
)

func main() {
	configs.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
