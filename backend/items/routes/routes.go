package routes

import (
	"github.com/labstack/echo/v4"
	"sinkzjs.org/m/v2/items/controller"
)

func Routes(controller controller.ItemsController, e *echo.Echo) {
	e.Group("i")

	e.GET("/items", controller.GetItems).Name = "get-items"
	e.GET("/items/:id", controller.GetItemById).Name = "get-item-by-id"
	e.POST("/items", controller.AddItem).Name = "add-item"
	e.PUT("/items/:id", controller.UpdateItem).Name = "update-item"

	e.GET("/items/rarities", controller.GetRarities).Name = "get-rarities"
}
