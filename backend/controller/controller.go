package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"sinkzjs.org/m/v2/db"
	"sinkzjs.org/m/v2/types"
)

type Controller struct {
	provider db.Provider
}

func NewController(provider db.Provider) *Controller {
	return &Controller{provider: provider}
}

func (c Controller) GetAllItems(ctx echo.Context) error {
	items, err := c.provider.GetItems()

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, types.ItemsResponse{
			Message:    err.Error(),
			HttpStatus: http.StatusBadRequest,
			Items:      nil,
		})
	}

	return ctx.JSON(http.StatusOK, types.ItemsResponse{
		Message:    "items retrieved successfully",
		HttpStatus: http.StatusOK,
		Items:      items,
	})
}

func (c Controller) GetItemById(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, types.ItemsResponse{
			Message:    fmt.Sprintf("invalid format for parameter [Id]: %s", ctx.Param("id")),
			HttpStatus: http.StatusBadRequest,
			Items:      nil,
		})
	}
	item, err := c.provider.GetItemById(id)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, types.ItemsResponse{
			Message:    fmt.Sprintf("item not found with ID: %d", id),
			HttpStatus: http.StatusNotFound,
			Items:      nil,
		})
	}

	return ctx.JSON(http.StatusOK, types.ItemsResponse{
		Message:    "item retrieved successfully",
		HttpStatus: http.StatusOK,
		Items:      []types.Item{item},
	})
}

func (c Controller) AddItem(ctx echo.Context) error {
	var item = new(types.Item)

	if err := ctx.Bind(item); err != nil {
		data := types.ItemsResponse{
			Message:    err.Error(),
			Items:      nil,
			HttpStatus: http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, data)
	}

	validate := validator.New()

	if err := validate.Struct(item); err != nil {
		return ctx.JSON(http.StatusBadRequest, types.ItemsResponse{
			Message:    err.Error(),
			Items:      nil,
			HttpStatus: http.StatusBadRequest,
		})
	}

	added_item, err := c.provider.AddItem(*item)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, types.ItemsResponse{
			Message:    err.Error(),
			Items:      nil,
			HttpStatus: http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusOK, types.ItemsResponse{
		Message:    "item added successfully",
		Items:      []types.Item{added_item},
		HttpStatus: http.StatusCreated,
	})
}
