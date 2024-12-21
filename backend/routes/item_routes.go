package routes

import (
	"github.com/labstack/echo/v4"
	"sinkzjs.org/m/v2/controllers"
)

func Routes(e *echo.Echo) {
	e.GET("/items", controllers.GetAllItems)
	e.GET("/items/:id", controllers.GetItemById)
	e.POST("/items", controllers.AddItem)
}
