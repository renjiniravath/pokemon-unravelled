package controller

import "github.com/renjiniravath/pokemon-unravelled/models"

//ListGenerations controller
func ListGenerations() ([]models.Generation, int, error) {
	return models.ListGenerations()
}
