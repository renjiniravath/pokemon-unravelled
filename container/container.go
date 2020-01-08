package container

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var dbReaderOnce, dbWriterOnce, sesClientOnce sync.Once
var dbReader, dbWriter *sqlx.DB

func SetDbReader(dbCon *sqlx.DB) {
	dbReaderOnce.Do(func() {
		dbReader = dbCon
	})
}

func GetDbReader() *sqlx.DB {
	return dbReader
}

func SetDbWriter(dbCon *sqlx.DB) {
	dbWriterOnce.Do(func() {
		dbWriter = dbCon
	})
}
func GetDbWriter() *sqlx.DB {
	return dbWriter
}
