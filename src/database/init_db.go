package database

import (
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var autoMigration = flag.Bool("auto-migration", true, "GORM autoMigration")

type IDataManager interface {
	DbConnect(manager *DataManager, dialect, args string) error
}
type InitDB struct {
	Recreate bool
	//db       *gorm.DB
}

type tabler interface {
	GetTableName() string
}

var tables = []tabler{
	//&user.User{},
}

func Tables() []tabler {
	return tables
}

func Init() *InitDB {
	return &InitDB{}
}

func (d *InitDB) DbConnect(manager *DataManager, dialect, args string) error {
	// TODO:: add timeout for docker
	db, err := gorm.Open(dialect, args)
	//defer db.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	manager.db = db
	if *autoMigration {
		d.autoMigrate(manager)
	}
	return nil
}

func (d *InitDB) autoMigrate(dm *DataManager) {
	for _, value := range Tables() {
		if dm.db.HasTable(value.GetTableName()) {
			dm.db.AutoMigrate(value)
		}
	}
}
