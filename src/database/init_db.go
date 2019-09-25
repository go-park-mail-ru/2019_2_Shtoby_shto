package database

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"log"
)

type DataManager struct {
	Recreate bool
	db       *sql.DB
}

func (d *DataManager) NewDataManager(dialect, args string) error {
	// TODO:: add timeout for docker
	db, err := sql.Open(dialect, args)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	d.db = db
	return nil
}

var autoMigration = flag.Bool("auto-migration", true, "GORM autoMigration")

type tabler interface {
	TableName() string
}

var tables = []tabler{}

func Tables() []tabler {
	return tables
}
