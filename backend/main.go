package main

import (
	"github.com/labstack/echo/v4"
	itemsController "sinkzjs.org/m/v2/items/controller"
	itemsDb "sinkzjs.org/m/v2/items/db"
	itemsRoutes "sinkzjs.org/m/v2/items/routes"

	exilesController "sinkzjs.org/m/v2/exiles/controller"
	exilesDb "sinkzjs.org/m/v2/exiles/db"
	exilesRoutes "sinkzjs.org/m/v2/exiles/routes"
)

func main() {
	go itemsServer()
	go exilesServer()

	select {}
}

func itemsServer() {
	e := echo.New()
	provider := itemsDb.NewInMemoryProvider("items/db/data/item_data.json")
	itemsController := *itemsController.NewController(provider)
	itemsRoutes.Routes(itemsController, e)
	e.Logger.Fatal(e.Start(":8080"))
}

func exilesServer() {
	e := echo.New()
	exilesProvider := exilesDb.NewInMemoryProvider("exiles/db/data/exile_data.json")
	exilesController := *exilesController.NewController(exilesProvider)
	exilesRoutes.Routes(exilesController, e)
	e.Logger.Fatal(e.Start(":9090"))
}
