package database

import (
	"2019_2_Shtoby_shto/src/dicts"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
)

//type IDataManager interface {
//	Db() *gorm.DB
//	DbConnect(dialect, args string)
//	FindDictById(sql string, p interface{}) error
//}

type DataManager struct {
	db *gorm.DB
}

func NewDataManager() *DataManager {
	return &DataManager{}
}

func (d DataManager) Db() *gorm.DB {
	return d.db
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
