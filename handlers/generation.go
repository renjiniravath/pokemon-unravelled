package handlers

import (
	"github.com/labstack/echo"
	"github.com/renjiniravath/pokemon-unravelled/core/logger"
	"github.com/renjiniravath/pokemon-unravelled/models"
	"net/http"
)

//ListGenerations lists all generations
func ListGenerations(c echo.Context) error {
	logger.Info.Info("Preparing to list generations")
	result := new(struct {
		Data  interface{} `json:"data"`
		Total int         `json:"total"`
	})
	generations, total, err := models.ListGenerations()
	if err != nil {
		logger.Error.Error("Error while getting list of generations", err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Error while getting list of generations")
	}
	result.Data = generations
	result.Total = total
	logger.Success.Info("Generations listed successfully")
	return c.JSON(http.StatusOK, result)
}
