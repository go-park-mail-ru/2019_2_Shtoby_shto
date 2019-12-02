package initDB

import (
	"2019_2_Shtoby_shto/src/dicts/models"
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
	&models.User{},
	&models.Photo{},
	&models.Board{},
	&models.Card{},
	&models.Comment{},
	&models.Tag{},
	&models.BoardUsers{},
	&models.CardUsers{},
	&models.CardGroup{},
	&models.CardTags{},
	&models.Message{},
	&models.CheckList{},
}

func Tables() []tabler {
	return tables
}

func Init() InitDBManager {
	return &InitDB{}
}

func (d *InitDB) DbConnect(dialect, args string) (*gorm.DB, error) {
	// TODO:: add timeout for docker
	db, err := gorm.Open(dialect, args)
	db.LogMode(true)
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
		db.AutoMigrate(value)
	}
}
