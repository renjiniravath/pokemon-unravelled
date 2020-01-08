package config

import (
	"github.com/caarlos0/env/v6"
)

//Current variable stores the environment variables
var Current Config

//Config struct defines the envs
type Config struct {
	MysqlMasterHost     string `env:"POKEMON_UNRAVELLED_API_MYSQL_MASTER_HOST,required"`
	MysqlMasterPort     int    `env:"POKEMON_UNRAVELLED_API_MYSQL_MASTER_PORT,required"`
	MysqlMasterDB       string `env:"POKEMON_UNRAVELLED_API_MYSQL_MASTER_DB,required"`
	MysqlMasterUsername string `env:"POKEMON_UNRAVELLED_API_MYSQL_MASTER_USERNAME,required"`
	MysqlMasterPassword string `env:"POKEMON_UNRAVELLED_API_MYSQL_MASTER_PASSWORD,required"`

	MysqlSlaveHost     string `env:"POKEMON_UNRAVELLED_API_MYSQL_SLAVE_HOST,required"`
	MysqlSlavePort     int    `env:"POKEMON_UNRAVELLED_API_MYSQL_SLAVE_PORT,required"`
	MysqlSlaveDB       string `env:"POKEMON_UNRAVELLED_API_MYSQL_SLAVE_DB,required"`
	MysqlSlaveUsername string `env:"POKEMON_UNRAVELLED_API_MYSQL_SLAVE_USERNAME,required"`
	MysqlSlavePassword string `env:"POKEMON_UNRAVELLED_API_MYSQL_SLAVE_PASSWORD,required"`

	LogFile        string `env:"POKEMON_UNRAVELLED_API_LOG_FILEPATH,required"`
	PokemonPerPage int    `env:"POKEMON_UNRAVELLED_API_POKEMON_PER_PAGE,required"`
}

//Load loads the envs
func Load() error {

	err := env.Parse(&Current)
	if err != nil {
		return err
	}
	return nil
}
