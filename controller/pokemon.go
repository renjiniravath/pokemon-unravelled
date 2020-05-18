package controller

import (
	"github.com/renjiniravath/pokemon-unravelled/models"
)

//ListPokemon controller
func ListPokemon(params *models.Pokemon, page int) (*[]models.Pokemon, int, error) {
	return models.ListPokemon(params, page)
}

//GetPokemonDetails controller
func GetPokemonDetails(uniqueID int) (models.Pokemon, error) {
	return models.GetPokemonDetails(uniqueID)
}

//GetGenerationAvailability controller
func GetGenerationAvailability(pokemonID int) ([]models.Generation, error) {
	return models.GetGenerationAvailability(pokemonID)
}

//GetPokemonData controller
func GetPokemonData(pokemonID int, generationID int) ([]models.Pokemon, models.PokemonBasicData, models.PokemonBasicData, error) {
	pokemon, err := models.GetPokemonData(pokemonID, generationID)
	if err != nil {
		return []models.Pokemon{}, models.PokemonBasicData{}, models.PokemonBasicData{}, err
	}
	prevPokemon, nextPokemon, err := models.GetPrevAndNextPokemon(pokemonID, generationID)
	if err != nil {
		return []models.Pokemon{}, models.PokemonBasicData{}, models.PokemonBasicData{}, err
	}
	return pokemon, prevPokemon, nextPokemon, nil
}
