package main

import (
	"github.com/labstack/echo/v4"
	"sinkzjs.org/m/v2/db"
	"sinkzjs.org/m/v2/routes"
)

func main() {
	db.LoadData()
	e := echo.New()
	routes.Routes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
