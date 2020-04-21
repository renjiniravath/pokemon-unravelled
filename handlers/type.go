package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/renjiniravath/pokemon-unravelled/controller"
	"github.com/renjiniravath/pokemon-unravelled/core/logger"
)

//ListTypes lists all types
func ListTypes(c echo.Context) error {
	logger.Info.Info("Preparing to list types")
	result := new(struct {
		Data  interface{} `json:"data"`
		Total int         `json:"total"`
	})
	types, total, err := controller.ListTypes()
	if err != nil {
		logger.Error.Error("Error while list of types", err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Error while list of types")
	}
	result.Data = types
	result.Total = total
	logger.Success.Info("Types listed successfully")
	return c.JSON(http.StatusOK, result)
}
