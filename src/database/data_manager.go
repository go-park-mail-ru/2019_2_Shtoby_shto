package database

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
	"errors"
	"fmt"
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

func (d DataManager) ExecuteQuery(sql string, args ...string) error {
	return d.db.Exec(sql, args).Error
}

func (d DataManager) FindDictById(p interface{}) error {
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	if !obj.GetId().IsUUID() {
		return errors.New("Not valid ID!")
	}
	res := d.db.Table(obj.GetTableName()).Where("id = ?", obj.GetId()).First(p)
	if res.RecordNotFound() || res.Error != nil {
		log.Println(res)
		return res.Error
	}
	return nil
}

func (d DataManager) FindDictByLogin(p interface{}, where, whereArg string) error {
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	res := d.db.Table(obj.GetTableName()).Where(fmt.Sprintf("%s = ?", where), whereArg).First(p)
	if res.RecordNotFound() || res.Error != nil {
		log.Println(res)
		return res.Error
	}
	return nil
}

func (d DataManager) CreateRecord(p interface{}) error {
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	res := d.db.Table(obj.GetTableName()).Create(p)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d DataManager) UpdateRecord(p interface{}, id customType.StringUUID) error {
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	obj.SetId(id)
	res := d.db.Table(obj.GetTableName()).Save(p)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d DataManager) DeleteRecord(p interface{}, id customType.StringUUID) error {
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	res := d.db.Table(obj.GetTableName()).Delete(p, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
