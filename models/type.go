
package models

import (
	"github.com/renjiniravath/pokemon-unravelled/container"
)

//Type stores id and name of a game generation
type Type struct {
	ID int `json:"id" db:"id,omitempty"`
	Name string `json:"name" db:"name,omitempty"`
}

//ListTypes lists all generations
func ListTypes()([]Type, int, error){
	query := "SELECT t.id, t.name FROM type AS t"
	db := container.GetDbReader()
	types := []Type{}
	err := db.Select(&types, query)
	if err != nil {
		return nil, 0, err
	}
	return types, len(types), nil
}

