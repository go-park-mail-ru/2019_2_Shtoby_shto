package database

import (
	"flag"
	"github.com/jinzhu/gorm"
	"log"
)

type InitDb struct {
	Recreate bool
	db       *gorm.DB
}

func NewDBService() (*gorm.DB, error) {
	// TODO:: add timeout for docker
	dbInfo := "postgres://postgres:Aebnm@postgres:5432/db_1?sslmode=disable"
	db, err := gorm.Open("postgres", dbInfo)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

var autoMigration = flag.Bool("auto-migration", true, "GORM autoMigration")

type tabler interface {
	TableName() string
}

var tables = []tabler{}

func Tables() []tabler {
	return tables
}
