package controller

import (
	"github.com/renjiniravath/pokemon-unravelled/models"
)

//ListPokemon controller
func ListPokemon(params *models.Pokemon, page int) (*[]models.Pokemon, int, error){
	return models.ListPokemon(params, page)
}

//GetPokemonDetails controller
func GetPokemonDetails(uniqueID int) (models.Pokemon, error){
	return models.GetPokemonDetails(uniqueID)
}

//GetGenerationAvailability controller
func GetGenerationAvailability(params models.Pokemon)([]models.Generation, error){
	return models.GetGenerationAvailability(params)
}