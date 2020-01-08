package models

import (
	"fmt"
	"strings"

	"github.com/renjiniravath/pokemon-unravelled/config"
	"github.com/renjiniravath/pokemon-unravelled/container"
)

//Pokemon stores details of a pokemon
type Pokemon struct {
	UniqueID        int     `json:"uniqueId" db:"uniqueId,omitempty"`
	ID              string  `json:"id" db:"id,omitempty"`
	Name            string  `json:"name" db:"name,omitempty"`
	PrimaryTypeID   *int    `json:"primaryTypeId" db:"primary_type_id,omitempty"`
	SecondaryTypeID *int    `json:"secondaryTypeId" db:"secondary_type_id,omitempty"`
	PrimaryType     *string `json:"primaryType" db:"primary_type_name,omitempty"`
	SecondaryType   *string `json:"secondaryType" db:"secondary_type_name,omitempty"`
	GenerationID    int     `json:"generationId" db:"generation_id, omitempty"`
	// GenerationName string `json:"generationName" db:"generation_name,omitempty"`
	FormID         *int    `json:"formId" db:"form_id,omitempty"`
	FormName       *string `json:"formName" db:"form_name, omitempty"`
	Attack         int     `json:"attack" db:"attack,omitempty"`
	Defense        int     `json:"defense" db:"defense,omitempty"`
	Speed          int     `json:"speed" db:"speed,omitempty"`
	SpecialAttack  int     `json:"specialAttack" db:"special_attack,omitempty"`
	SpecialDefense int     `json:"specialDefense" db:"special_defense,omitempty"`
	HP             int     `json:"hp" db:"hp,omitempty"`
}

//ListPokemon lists pokemon according to search parameters
func ListPokemon(params *Pokemon, page int) (*[]Pokemon, int, error) {
	query := "SELECT p.id, p.name, p_t.id AS uniqueId, p_t.primary_type_id, p_t.secondary_type_id, p_t.generation_id,p_p_t.name AS primary_type_name, p_t.attack, p_t.defense, p_t.speed, p_t.special_attack, p_t.special_defense, p_t.hp, s_p_t.name AS secondary_type_name, p_t.form_id, p_f.name AS form_name FROM pokemon_stats AS p_t JOIN pokemon AS p ON p.id = p_t.pokemon_id LEFT JOIN type AS p_p_t ON p_p_t.id = p_t.primary_type_id LEFT JOIN type AS s_p_t ON s_p_t.id = p_t.secondary_type_id LEFT JOIN pokemon_form AS p_f ON p_f.id = p_t.form_id %s LIMIT %d, %d"
	whereConditions := []string{}
	var whereParams []interface{}
	if params.Name != "" {
		whereConditions = append(whereConditions, "(p.name LIKE ?)")
		key := "%" + params.Name + "%"
		whereParams = append(whereParams, key)
	}
	if params.ID != "" {
		whereConditions = append(whereConditions, "(p.id LIKE ?)")
		key := params.ID + "%"
		whereParams = append(whereParams, key)
	}
	if *params.PrimaryTypeID != 0 {
		whereConditions = append(whereConditions, "(p_t.primary_type_id = ? OR p_t.secondary_type_id = ?)")
		whereParams = append(whereParams, *params.PrimaryTypeID, *params.PrimaryTypeID)
	}
	if *params.SecondaryTypeID != 0 {
		whereConditions = append(whereConditions, "(p_t.primary_type_id = ? OR p_t.secondary_type_id = ?)")
		whereParams = append(whereParams, *params.SecondaryTypeID, *params.SecondaryTypeID)
	}
	if *params.FormID != 0 {
		whereConditions = append(whereConditions, "(p_t.form_id = ?)")
		whereParams = append(whereParams, *params.FormID)
	}
	whereConditions = append(whereConditions, "(p_t.generation_id = ?)")
	whereParams = append(whereParams, params.GenerationID)
	whereClause := ""
	if len(whereConditions) > 0 {
		whereClause = "WHERE" + strings.Join(whereConditions, "AND ")
	}
	start := 0
	if page > 1 {
		start = config.Current.PokemonPerPage * (page - 1)
	}
	finalQuery := fmt.Sprintf(query, whereClause, start, config.Current.PokemonPerPage)
	db := container.GetDbReader()
	list := []Pokemon{}
	err := db.Select(&list, finalQuery, whereParams...)
	if err != nil {
		return nil, 0, err
	}
	var count int
	query = "SELECT COUNT(*) FROM pokemon_stats AS p_t JOIN pokemon AS p ON p.id = p_t.pokemon_id %s"
	finalQuery = fmt.Sprintf(query, whereClause)
	err = db.Get(&count, finalQuery, whereParams...)
	if len(list) == 0 {
		return nil, 0, err
	}
	return &list, count, err
}

//GetPokemonDetails gets the details of a pokemon
func GetPokemonDetails(uniqueID int) (Pokemon, error) {
	query := "SELECT p.id, p.name, p_t.id AS uniqueId, p_t.primary_type_id, p_t.secondary_type_id, p_t.generation_id,p_p_t.name AS primary_type_name, p_t.attack, p_t.defense, p_t.speed, p_t.special_attack, p_t.special_defense, p_t.hp, s_p_t.name AS secondary_type_name, p_t.form_id, p_f.name AS form_name FROM pokemon_stats AS p_t JOIN pokemon AS p ON p.id = p_t.pokemon_id LEFT JOIN type AS p_p_t ON p_p_t.id = p_t.primary_type_id LEFT JOIN type AS s_p_t ON s_p_t.id = p_t.secondary_type_id LEFT JOIN pokemon_form AS p_f ON p_f.id = p_t.form_id WHERE p_t.id=?"
	db := container.GetDbReader()
	var pokemon Pokemon
	row := db.QueryRowx(query, uniqueID)
	err := row.StructScan(&pokemon)
	return pokemon, err
}

//GetGenerationAvailability returns the list of generations applicable to a pokemon
func GetGenerationAvailability(params Pokemon) ([]Generation, error) {
	query := "SELECT p_t.generation_id AS id, generation.name FROM pokemon_stats AS p_t JOIN generation ON generation.id = p_t.generation_id %s"
	whereConditions := []string{}
	var whereParams []interface{}
	if params.ID != "" {
		whereConditions = append(whereConditions, "(p_t.pokemon_id = ?)")
		key := params.ID
		whereParams = append(whereParams, key)
	}
	if *params.FormID != 0 {
		whereConditions = append(whereConditions, "(p_t.form_id = ?)")
		whereParams = append(whereParams, *params.FormID)
	}
	whereClause := ""
	whereClause = "WHERE" + strings.Join(whereConditions, "AND")
	finalQuery := fmt.Sprintf(query, whereClause)
	db := container.GetDbReader()
	var generationList []Generation
	err := db.Select(&generationList, finalQuery, whereParams...)
	if err != nil {
		return nil, err
	}
	return generationList, nil
}
