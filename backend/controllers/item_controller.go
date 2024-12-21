package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"sinkzjs.org/m/v2/db"
	"sinkzjs.org/m/v2/types"
)

func GetAllItems(c echo.Context) error {
	itemsMap, err := db.GetItems()

	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ItemsResponse{
			Message:    err.Error(),
			HttpStatus: http.StatusBadRequest,
			Items:      nil,
		})
	}

	items := make([]types.Item, 0, len(itemsMap))

	for _, item := range itemsMap {
		items = append(items, item)
	}

	response := types.ItemsResponse{
		Message:    "items retrieved successfully",
		HttpStatus: http.StatusOK,
		Items:      items,
	}
	return c.JSON(http.StatusOK, response)
}

func GetItemById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		response := types.ItemsResponse{
			Message:    fmt.Sprintf("invalid format for parameter [Id]: %s", c.Param("id")),
			HttpStatus: http.StatusBadRequest,
			Items:      nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	item, err := db.GetItemById(id)

	if err != nil {
		response := types.ItemsResponse{
			Message:    fmt.Sprintf("item not found with ID: %d", id),
			HttpStatus: http.StatusNotFound,
			Items:      nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := types.ItemsResponse{
		Message:    "item retrieved successfully",
		HttpStatus: http.StatusOK,
		Items:      []types.Item{*item},
	}

	return c.JSON(http.StatusOK, response)
}

func AddItem(c echo.Context) error {
	var item = new(types.Item)

	if err := c.Bind(item); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	validate := validator.New()

	if err := validate.Struct(item); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	added_item, err := db.AddItem(*item)

	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	response := map[string]interface{}{
		"message": "item added successfully",
		"data":    added_item,
	}

	return c.JSON(http.StatusOK, response)
}
