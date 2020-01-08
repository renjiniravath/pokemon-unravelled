package routes

import (
	"github.com/labstack/echo"
	"github.com/renjiniravath/pokemon-unravelled/handlers"
)

//Set sets all routes
func Set(e *echo.Echo) {
	e.GET("/pokemon", handlers.ListPokemon)
	e.GET("/pokemon/:id", handlers.GetPokemonDetails)
	e.GET("/pokemon/generations", handlers.GetGenerationAvailability)
	e.GET("/generations", handlers.ListGenerations)
	e.GET("/types", handlers.ListTypes)
}
