package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/renjiniravath/pokemon-unravelled/config"
	"github.com/renjiniravath/pokemon-unravelled/controller"
	"github.com/renjiniravath/pokemon-unravelled/core/logger"
	"github.com/renjiniravath/pokemon-unravelled/models"
)

//ListPokemon gives a list of pokemon
func ListPokemon(c echo.Context) error {
	logger.Info.Info("Preparing to list pokemon")
	request := new(struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Type1        int    `json:"type1"`
		Type2        int    `json:"type2"`
		GenerationID int    `json:"generationId"`
		FormID       int    `json:"formId"`
		Page         int    `json:"page"`
	})
	result := new(struct {
		Data      interface{} `json:"data"`
		Total     int         `json:"total"`
		NoOfPages float64     `json:"noOfPages"`
	})
	if err := c.Bind(request); err != nil {
		logger.Error.Error("Wrong format inserted")
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong format inserted")
	}
	if request.Page == 0 {
		request.Page = 1
	}
	pokemon := models.Pokemon{
		ID:              request.ID,
		Name:            request.Name,
		PrimaryTypeID:   &request.Type1,
		SecondaryTypeID: &request.Type2,
		FormID:          &request.FormID,
		GenerationID:    request.GenerationID,
	}
	pokemons, total, err := controller.ListPokemon(&pokemon, request.Page)
	if err != nil {
		logger.Error.Error("Error while getting pokemon list", err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Error while getting pokemon list")
	}
	if pokemons == nil {
		result.Data = []models.Pokemon{}
		result.Total = 0
		return c.JSON(http.StatusOK, result)
	}
	result.Data = *pokemons
	result.Total = total
	result.NoOfPages = math.Ceil(float64(total) / float64(config.Current.PokemonPerPage))
	logger.Success.Info("Pokemons successfully listed")
	return c.JSON(http.StatusOK, result)
}

//GetPokemonDetails gets all the details of a pokemon
func GetPokemonDetails(c echo.Context) error {
	logger.Info.Info("Getting details of pokemon")
	uniqueID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Error("Wrong format inserted")
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong format inserted")
	}
	result := new(struct {
		Data interface{} `json:"data"`
	})
	pokemon, err := controller.GetPokemonDetails(uniqueID)
	if err != nil {
		logger.Error.Error("Error while getting details of pokemon table id ", uniqueID)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Error while getting details of pokemon table id ", uniqueID)
	}
	result.Data = pokemon
	logger.Success.Info("Pokemon details successfully shown")
	return c.JSON(http.StatusOK, result)
}

//GetPokemonData gets the data of all forms of a particular pokemon in a particular generation
func GetPokemonData(c echo.Context) error {
	logger.Info.Info("Getting the data of all forms of a pokemon")
	pokemonID, err := strconv.Atoi(c.Param("pokemonId"))
	if err != nil {
		logger.Error.Error("Wrong format for pokemonID inserted ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong format for pokemonID inserted")
	}
	generationID, err := strconv.Atoi(c.Param("generationId"))
	if err != nil {
		logger.Error.Error("Wrong format for generationID inserted")
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong format for generationID inserted")
	}
	result := new(struct {
		Data        interface{} `json:"data"`
		PrevPokemon interface{} `json:"prevPokemon"`
		NextPokemon interface{} `json:"nextPokemon"`
	})
	pokemonData, prevPokemon, nextPokemon, err := controller.GetPokemonData(pokemonID, generationID)
	if err != nil {
		logger.Error.Error("Error while getting data of pokemon id ", pokemonID, " generation id ", generationID, err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Error while getting data of pokemon id ", pokemonID, " generation id ", generationID)
	}
	result.Data = pokemonData
	result.PrevPokemon = prevPokemon
	result.NextPokemon = nextPokemon
	logger.Success.Info("Pokemon data successfully returned")
	return c.JSON(http.StatusOK, result)
}

//GetGenerationAvailability returns the list of generations applicable to a pokemon
func GetGenerationAvailability(c echo.Context) error {
	logger.Info.Info("Getting list of generations applicable to a pokemon")
	pokemonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Error("PokemonID inserted wrong ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "PokemonID inserted wrong")
	}
	result := new(struct {
		Data interface{} `json:"data"`
	})
	generationsList, err := controller.GetGenerationAvailability(pokemonID)
	if err != nil {
		logger.Error.Error("Error while getting generations applicable to pokemon", err.Error())
		return echo.NewHTTPError(http.StatusNotAcceptable, "Error while getting generations applicable to pokemon")
	}
	result.Data = generationsList
	logger.Success.Info("Applicable generations list fetched successfully")
	return c.JSON(http.StatusOK, result)
}
