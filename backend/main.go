package main

import (
	"github.com/labstack/echo/v4"
	itemsController "sinkzjs.org/m/v2/items/controller"
	itemsDb "sinkzjs.org/m/v2/items/db"
	itemsRoutes "sinkzjs.org/m/v2/items/routes"
)

func main() {
	e := echo.New()
	provider := itemsDb.NewInMemoryProvider("items/db/data/item_data.json")
	itemsController := *itemsController.NewController(provider)
	itemsRoutes.Routes(itemsController, e)
	e.Logger.Fatal(e.Start(":8080"))
}
