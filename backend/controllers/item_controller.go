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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        err.Error(),
			"request_status": http.StatusBadRequest,
			"data":           nil,
		})
	}

	items := make([]types.Item, 0, len(itemsMap))

	for _, item := range itemsMap {
		items = append(items, item)
	}

	response := map[string]interface{}{
		"message":        "items retrieved successfully",
		"request_status": http.StatusOK,
		"data":           items,
	}
	return c.JSON(http.StatusOK, response)
}

func GetItemById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		response := map[string]interface{}{
			"message":        fmt.Sprintf("invalid format for parameter [Id]: %s", c.Param("id")),
			"request_status": http.StatusBadRequest,
			"data":           nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	item, err := db.GetItemById(id)

	if err != nil {
		response := map[string]interface{}{
			"message":        fmt.Sprintf("item not found with ID: %d", id),
			"request_status": http.StatusNotFound,
			"data":           nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := map[string]interface{}{
		"message":        "item retrieved successfully",
		"request_status": http.StatusOK,
		"data":           item,
	}

	return c.JSON(http.StatusOK, response)
}

func AddItemHandler(c echo.Context) error {
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

func GetRarities(c echo.Context) error {
	rarities, err := db.GetRarities()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        err.Error(),
			"request_status": http.StatusBadRequest,
			"data":           nil,
		})
	}

	response := map[string]interface{}{
		"message":        "rarities retrieved successfully",
		"request_status": http.StatusOK,
		"data":           rarities,
	}
	return c.JSON(http.StatusOK, response)
}

func GetItemTypes(c echo.Context) error {
	itemTypes, err := db.GetItemTypes()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        err.Error(),
			"request_status": http.StatusBadRequest,
			"data":           nil,
		})
	}

	response := map[string]interface{}{
		"message":        "item types retrieved successfully",
		"request_status": http.StatusOK,
		"data":           itemTypes,
	}
	return c.JSON(http.StatusOK, response)
}

func GetImages(c echo.Context) error {
	images, err := db.GetImages()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        err.Error(),
			"request_status": http.StatusBadRequest,
			"data":           nil,
		})
	}

	response := map[string]interface{}{
		"message":        "images retrieved successfully",
		"request_status": http.StatusOK,
		"data":           images,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAttributes(c echo.Context) error {
	attributes, err := db.GetAttributes()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        err.Error(),
			"request_status": http.StatusBadRequest,
			"data":           nil,
		})
	}

	response := map[string]interface{}{
		"message":        "attributes retrieved successfully",
		"request_status": http.StatusOK,
		"data":           attributes,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAttributeGroupings(c echo.Context) error {
	attributeGroupings, err := db.GetAttributeGroupings()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        err.Error(),
			"request_status": http.StatusBadRequest,
			"data":           nil,
		})
	}

	response := map[string]interface{}{
		"message":        "attribute groupings retrieved successfully",
		"request_status": http.StatusOK,
		"data":           attributeGroupings,
	}
	return c.JSON(http.StatusOK, response)
}