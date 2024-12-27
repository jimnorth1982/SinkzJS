package routes

import (
	"github.com/labstack/echo/v4"
	"sinkzjs.org/m/v2/exiles/controller"
)

func Routes(controller controller.Controller, e *echo.Echo) {
	e.Group("e")
	e.GET("/exiles", controller.GetExiles).Path = "get-exiles"
	e.GET("/exiles/:id", controller.GetExile).Name = "get-exile"
	e.POST("/exiles", controller.CreateExile).Name = "create-exile"
	e.PUT("/exiles/:id", controller.UpdateExile).Name = "update-exile"
	e.DELETE("/exiles/:id", controller.DeleteExile).Name = "delete-exile"
}
