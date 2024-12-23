package main

import (
	"github.com/labstack/echo/v4"
	"sinkzjs.org/m/v2/controller"
	"sinkzjs.org/m/v2/db"
	"sinkzjs.org/m/v2/routes"
)

func main() {
	e := echo.New()
	routes.Routes(*controller.NewController(db.NewInMemoryProvider("db/data/item_data.json")), e)
	e.Logger.Fatal(e.Start(":8080"))
}
