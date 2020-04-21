package controller

import "github.com/renjiniravath/pokemon-unravelled/models"

//ListTypes controller
func ListTypes() ([]models.Type, int, error) {
	return models.ListTypes()
}
