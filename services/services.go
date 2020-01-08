package services

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/renjiniravath/pokemon-unravelled/config"
	"github.com/renjiniravath/pokemon-unravelled/container"
)

//Load all services
func Load() error {
	mysqlMaster, err := getMySqlConnection(
		config.Current.MysqlMasterUsername,
		config.Current.MysqlMasterPassword,
		config.Current.MysqlMasterHost,
		config.Current.MysqlMasterPort,
		config.Current.MysqlMasterDB)

	if err != nil {
		return err
	}
	container.SetDbWriter(mysqlMaster)

	mysqlSlave, err := getMySqlConnection(
		config.Current.MysqlSlaveUsername,
		config.Current.MysqlSlavePassword,
		config.Current.MysqlSlaveHost,
		config.Current.MysqlSlavePort,
		config.Current.MysqlSlaveDB)

	if err != nil {
		return err
	}
	container.SetDbReader(mysqlSlave)

	return nil
}

func getMySqlConnection(username string, password string, host string, port int, dbname string) (*sqlx.DB, error) {
	const mysqlConnectionString = "%s:%s@tcp(%s:%d)/%s"
	connectionString := fmt.Sprintf(mysqlConnectionString, username, password, host, port, dbname)
	db, err := sqlx.Connect("mysql", connectionString)

	return db, err

}
