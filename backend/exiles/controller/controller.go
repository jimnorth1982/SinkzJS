package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"sinkzjs.org/m/v2/exiles/db"
	"sinkzjs.org/m/v2/exiles/types"
)

type Controller struct {
	Provider db.Provider
}

func NewController(provider db.Provider) *Controller {
	return &Controller{Provider: provider}
}

func (c Controller) GetExiles(ctx echo.Context) error {
	exiles, err := c.Provider.GetExiles()

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, GenError(err))
	}

	return ctx.JSON(http.StatusOK, types.ExilesResponse{
		Message:    "exiles retrieved successfully",
		HttpStatus: http.StatusOK,
		Exiles:     exiles,
	})
}

func (c Controller) GetExile(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, types.ExilesResponse{
			Message:    fmt.Sprintf("invalid format for parameter [Id]: %s", ctx.Param("id")),
			HttpStatus: http.StatusBadRequest,
			Exiles:     nil,
		})
	}

	exile, err := c.Provider.GetExile(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, GenError(err))
	}

	return ctx.JSON(http.StatusOK, types.ExilesResponse{
		Message:    "exile retrieved successfully",
		HttpStatus: http.StatusOK,
		Exiles:     []types.Exile{exile},
	})
}

func GenError(err error) types.ExilesResponse {
	return types.ExilesResponse{
		Message:    err.Error(),
		HttpStatus: http.StatusBadRequest,
		Exiles:     nil,
	}
}

func (c Controller) CreateExile(ctx echo.Context) error {
	var exile types.Exile
	if err := ctx.Bind(&exile); err != nil {
		return ctx.JSON(http.StatusBadRequest, GenError(err))
	}

	if c.Provider.ExileNameExistsInDb(exile.Name) {
		return ctx.JSON(http.StatusBadRequest, types.ExilesResponse{
			Message:    fmt.Sprintf("exile with name %s already exists", exile.Name),
			HttpStatus: http.StatusBadRequest,
			Exiles:     nil,
		})
	}

	added_exile, err := c.Provider.CreateExile(exile)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, GenError(err))
	}

	return ctx.JSON(http.StatusCreated, types.ExilesResponse{
		Message:    "exile created successfully",
		HttpStatus: http.StatusCreated,
		Exiles:     []types.Exile{added_exile},
	})
}

func (c Controller) UpdateExile(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, types.ExilesResponse{
			Message:    fmt.Sprintf("invalid format for parameter [Id]: %s", ctx.Param("id")),
			HttpStatus: http.StatusBadRequest,
			Exiles:     nil,
		})
	}

	var exile types.Exile
	if err := ctx.Bind(&exile); err != nil {
		return ctx.JSON(http.StatusBadRequest, GenError(err))
	}

	updated_exile, err := c.Provider.UpdateExile(id, exile)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, GenError(err))
	}

	return ctx.JSON(http.StatusOK, types.ExilesResponse{
		Message:    "exile updated successfully",
		HttpStatus: http.StatusOK,
		Exiles:     []types.Exile{updated_exile},
	})
}

func (c Controller) DeleteExile(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, types.ExilesResponse{
			Message:    fmt.Sprintf("invalid format for parameter [Id]: %s", ctx.Param("id")),
			HttpStatus: http.StatusBadRequest,
			Exiles:     nil,
		})
	}

	err = c.Provider.DeleteExile(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, GenError(err))
	}

	return ctx.JSON(http.StatusOK, types.ExilesResponse{
		Message:    "exile deleted successfully",
		HttpStatus: http.StatusOK,
		Exiles:     nil,
	})
}

func (c Controller) ExileNameExistsInDb(ctx echo.Context) error {
	name := ctx.Param("name")
	exists := c.Provider.ExileNameExistsInDb(name)

	return ctx.JSON(http.StatusOK, types.ExilesResponse{
		Message:    fmt.Sprintf("exile with name %s exists: %t", name, exists),
		HttpStatus: http.StatusOK,
		Exiles:     nil,
	})
}
