package initDB

import (
	"2019_2_Shtoby_shto/src/dicts/board"
	"2019_2_Shtoby_shto/src/dicts/photo"
	"2019_2_Shtoby_shto/src/dicts/user"
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var autoMigration = flag.Bool("auto-migration", true, "GORM autoMigration")

type InitDBManager interface {
	DbConnect(dialect, args string) (*gorm.DB, error)
}
type InitDB struct {
	Recreate bool
	//db       *gorm.DB
}

type tabler interface {
	GetTableName() string
}

var tables = []tabler{
	&user.User{},
	&photo.Photo{},
	&board.Board{},
}

func Tables() []tabler {
	return tables
}

func Init() *InitDB {
	return &InitDB{}
}

func (d *InitDB) DbConnect(dialect, args string) (*gorm.DB, error) {
	// TODO:: add timeout for docker
	db, err := gorm.Open(dialect, args)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if *autoMigration {
		d.autoMigrate(db)
	}
	return db, nil
}

func (d *InitDB) autoMigrate(db *gorm.DB) {
	for _, value := range Tables() {
		if db.HasTable(value.GetTableName()) {
			db.AutoMigrate(value)
		}
	}
}
