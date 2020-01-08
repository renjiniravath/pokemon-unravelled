package models

import (
	"github.com/renjiniravath/pokemon-unravelled/container"
)

//Generation stores id and name of a game generation
type Generation struct {
	ID int `json:"id" db:"id,omitempty"`
	Name string `json:"name" db:"name,omitempty"`
}

//ListGenerations lists all generations
func ListGenerations()([]Generation, int, error){
	query := "SELECT g.id, g.name FROM generation AS g"
	db := container.GetDbReader()
	generations := []Generation{}
	err := db.Select(&generations, query)
	if err != nil {
		return nil, 0, err
	}
	return generations, len(generations), nil
}