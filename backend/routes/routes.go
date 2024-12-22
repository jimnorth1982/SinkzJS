package routes

import (
	"github.com/labstack/echo/v4"
	"sinkzjs.org/m/v2/controller"
)

func Routes(controller controller.Controller, e *echo.Echo) {
	e.GET("/items", controller.GetAllItems)
	e.GET("/items/:id", controller.GetItemById)
	e.POST("/items", controller.AddItem)
}
