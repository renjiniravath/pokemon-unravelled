package models

import (
	"fmt"
	"strings"

	"github.com/renjiniravath/pokemon-unravelled/config"
	"github.com/renjiniravath/pokemon-unravelled/container"
)

//PokemonBasicData stores the most basic data of a pokemon
type PokemonBasicData struct {
	ID   string `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`
}

//Pokemon stores details of a pokemon
type Pokemon struct {
	UniqueID        int     `json:"uniqueId,omitempty" db:"uniqueId"`
	ID              string  `json:"id,omitempty" db:"id"`
	Name            string  `json:"name,omitempty" db:"name"`
	PrimaryTypeID   *int    `json:"primaryTypeId,omitempty" db:"primary_type_id"`
	SecondaryTypeID *int    `json:"secondaryTypeId,omitempty" db:"secondary_type_id"`
	PrimaryType     *string `json:"primaryType,omitempty" db:"primary_type_name"`
	SecondaryType   *string `json:"secondaryType,omitempty" db:"secondary_type_name"`
	GenerationID    int     `json:"generationId,omitempty" db:"generation_id, omitempty"`
	// GenerationName string `json:"generationName,omitempty" db:"generation_name"`
	FormID         *int    `json:"formId,omitempty" db:"form_id"`
	FormName       *string `json:"formName,omitempty" db:"form_name, omitempty"`
	Attack         int     `json:"attack,omitempty" db:"attack"`
	Defense        int     `json:"defense,omitempty" db:"defense"`
	Speed          int     `json:"speed,omitempty" db:"speed"`
	SpecialAttack  int     `json:"specialAttack,omitempty" db:"special_attack"`
	SpecialDefense int     `json:"specialDefense,omitempty" db:"special_defense"`
	HP             int     `json:"hp,omitempty" db:"hp"`
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

//GetPokemonData gets the data of all forms of a pokemon from a generation
func GetPokemonData(pokemonID int, generationID int) ([]Pokemon, error) {
	query := "SELECT p.id, p.name, p_t.id AS uniqueId, p_t.primary_type_id, p_t.secondary_type_id, p_t.generation_id,p_p_t.name AS primary_type_name, p_t.attack, p_t.defense, p_t.speed, p_t.special_attack, p_t.special_defense, p_t.hp, s_p_t.name AS secondary_type_name, COALESCE(p_t.form_id,0) AS form_id, TRIM(CONCAT(COALESCE(SPLIT_STR(p_f.name, '-', 1),''),' ',p.name,' ',COALESCE(SPLIT_STR(p_f.name, '-', 2),''))) AS form_name FROM pokemon_stats AS p_t JOIN pokemon AS p ON p.id = p_t.pokemon_id LEFT JOIN type AS p_p_t ON p_p_t.id = p_t.primary_type_id LEFT JOIN type AS s_p_t ON s_p_t.id = p_t.secondary_type_id LEFT JOIN pokemon_form AS p_f ON p_f.id = p_t.form_id WHERE p.id=? AND p_t.generation_id=? ORDER BY p_f.id"
	db := container.GetDbReader()
	list := []Pokemon{}
	err := db.Select(&list, query, pokemonID, generationID)
	if err != nil {
		return nil, err
	}
	return list, err
}

//GetPrevAndNextPokemon gets the data of the pokemons coming before and after a particular pokemon
func GetPrevAndNextPokemon(pokemonID int, generationID int) (PokemonBasicData, PokemonBasicData, error) {
	var prevID, nextID int
	var count int
	db := container.GetDbReader()
	query := "SELECT COUNT(*) FROM pokemon JOIN pokemon_stats ON pokemon.id = pokemon_stats.pokemon_id WHERE pokemon_stats.generation_id=? AND pokemon_stats.form_id IS NULL"
	err := db.Get(&count, query, generationID)
	if err != nil {
		return PokemonBasicData{}, PokemonBasicData{}, err
	}
	prevID = pokemonID - 1
	nextID = pokemonID + 1
	if pokemonID == 1 {
		prevID = count
	}
	if pokemonID == count {
		nextID = 1
	}
	query = "SELECT p.id, p.name FROM pokemon_stats AS p_s JOIN pokemon AS p ON p.id = p_s.pokemon_id WHERE p.id=? AND p_s.generation_id=? AND p_s.form_id IS NULL"
	var prevPokemonData, nextPokemonData PokemonBasicData
	row := db.QueryRowx(query, prevID, generationID)
	err = row.StructScan(&prevPokemonData)
	if err != nil {
		return PokemonBasicData{}, PokemonBasicData{}, err
	}
	row = db.QueryRowx(query, nextID, generationID)
	err = row.StructScan(&nextPokemonData)
	if err != nil {
		return PokemonBasicData{}, PokemonBasicData{}, err
	}
	return prevPokemonData, nextPokemonData, nil
}

//GetGenerationAvailability returns the list of generations applicable to a pokemon
func GetGenerationAvailability(pokemonID int) ([]Generation, error) {
	query := "SELECT p_s.generation_id AS id, generation.name FROM pokemon_stats AS p_s JOIN generation ON generation.id = p_s.generation_id WHERE p_s.pokemon_id=? AND p_s.form_id IS NULL"
	db := container.GetDbReader()
	var generationList []Generation
	err := db.Select(&generationList, query, pokemonID)
	if err != nil {
		return nil, err
	}
	return generationList, nil
}
