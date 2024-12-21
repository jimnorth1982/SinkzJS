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
	items, err := db.GetItems()

	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ItemsResponse{
			Message:    err.Error(),
			HttpStatus: http.StatusBadRequest,
			Items:      nil,
		})
	}

	return c.JSON(http.StatusOK, types.ItemsResponse{
		Message:    "items retrieved successfully",
		HttpStatus: http.StatusOK,
		Items:      items,
	})
}

func GetItemById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, types.ItemsResponse{
			Message:    fmt.Sprintf("invalid format for parameter [Id]: %s", c.Param("id")),
			HttpStatus: http.StatusBadRequest,
			Items:      nil,
		})
	}
	item, err := db.GetItemById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, types.ItemsResponse{
			Message:    fmt.Sprintf("item not found with ID: %d", id),
			HttpStatus: http.StatusNotFound,
			Items:      nil,
		})
	}

	return c.JSON(http.StatusOK, types.ItemsResponse{
		Message:    "item retrieved successfully",
		HttpStatus: http.StatusOK,
		Items:      []types.Item{*item},
	})
}

func AddItem(c echo.Context) error {
	var item = new(types.Item)

	if err := c.Bind(item); err != nil {
		data := types.ItemsResponse{
			Message:    err.Error(),
			Items:      nil,
			HttpStatus: http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	validate := validator.New()

	if err := validate.Struct(item); err != nil {
		return c.JSON(http.StatusBadRequest, types.ItemsResponse{
			Message:    err.Error(),
			Items:      nil,
			HttpStatus: http.StatusBadRequest,
		})
	}

	added_item, err := db.AddItem(*item)

	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ItemsResponse{
			Message:    err.Error(),
			Items:      nil,
			HttpStatus: http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, types.ItemsResponse{
		Message:    "item added successfully",
		Items:      []types.Item{added_item},
		HttpStatus: http.StatusCreated,
	})
}
