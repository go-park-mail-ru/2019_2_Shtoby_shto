package database

import (
	"flag"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
)

type DataManager struct {
	Recreate bool
	db       *gorm.DB
	cache    *redis.Client
}

func (d *DataManager) Init(dialect, args string) error {
	// TODO:: add timeout for docker
	db, err := gorm.Open(dialect, args)
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
