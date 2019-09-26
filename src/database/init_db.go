package database

import (
	"2019_2_Shtoby_shto/src/dicts"
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"reflect"
)

type IDataManager interface {
	DbConnect(dialect, args string)
	FindDictById(sql string, p interface{}) error
}

type DataManager struct {
	Recreate bool
	db       *gorm.DB
}

func (d *DataManager) DbConnect(dialect, args string) error {
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

func (d DataManager) ExecuteQueryUser(sql string, args ...string) error {
	return d.db.Exec(sql, args).Error
}

func (d DataManager) FindDictById(sql string, p interface{}) error {
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	res := d.db.Table(obj.GetTableName()).Where("id = ?", obj.GetId()).First(p)
	if res.RecordNotFound() || res.Error != nil {
		log.Fatal(res)
		return res.Error
	}
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
