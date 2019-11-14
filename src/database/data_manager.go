package database

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/utils"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
	"strings"
)

type IDataManager interface {
	Db() *gorm.DB
	SetDb(db *gorm.DB)
	CloseConnection() error
	ExecuteQuery(sql string, args ...string) error
	FindDictById(p interface{}) error
	FindDictByColumn(p interface{}) (int, error)
	FindCountDictByColumn(p interface{}) (int, error)
	CreateRecord(p interface{}) error
	UpdateRecord(p interface{}, id customType.StringUUID) error
	DeleteRecord(p interface{}, id customType.StringUUID) error
	FetchDict(data interface{}, table string, limit, offset int, where, whereArg []string) (int, error)
}

type DataManager struct {
	db *gorm.DB
}

func NewDataManager(db *gorm.DB) IDataManager {
	return &DataManager{
		db: db,
	}
}

func (d DataManager) Db() *gorm.DB {
	return d.db
}

func (d DataManager) SetDb(db *gorm.DB) {
	d.db = db
}

func (d *DataManager) CloseConnection() error {
	return d.db.Close()
}

func (d DataManager) ExecuteQuery(sql string, args ...string) error {
	return d.db.Exec(sql, args).Error
}

func (d DataManager) FindDictById(p interface{}) error {
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	if !obj.GetId().IsUUID() {
		return errors.New("Not valid ID! ")
	}
	res := d.db.Table(obj.GetTableName()).Where("id = ?", obj.GetId()).First(p)
	if res.RecordNotFound() || res.Error != nil {
		log.Println(res)
		return res.Error
	}
	return nil
}

func (d DataManager) FindDictByColumn(p interface{}) (int, error) {
	count := 0
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	res := d.db.Table(obj.GetTableName()).Where(p).First(p)
	if res.Error != nil {
		log.Println(res)
		return 0, res.Error
	}
	if res.RecordNotFound() {
		return 0, nil
	}
	return count, nil
}

func (d DataManager) FindCountDictByColumn(p interface{}) (int, error) {
	count := 0
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	res := d.db.Table(obj.GetTableName()).Where(p).Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func (d DataManager) FetchDict(data interface{}, table string, limit, offset int, where, whereArg []string) (int, error) {
	fetchLen := 0
	whereResult := strings.Join(where, " and ")
	res := d.db.Table(table).Where(whereResult, whereArg).Limit(limit).Offset(offset).Find(data)
	if res.RecordNotFound() {
		return 0, nil
	} else if res.Error != nil {
		return 0, res.Error
	}
	return fetchLen, nil
}

func (d DataManager) CreateRecord(p interface{}) error {
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	id := obj.GetId().String()
	if obj.GetId().IsEmpty() {
		newID, err := utils.GenerateUUID()
		if err != nil {
			return err
		}
		id = newID.String()
	}
	obj.SetId(customType.StringUUID(id))
	res := d.db.Table(obj.GetTableName()).Create(p)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d DataManager) UpdateRecord(p interface{}, id customType.StringUUID) error {
	obj := reflect.ValueOf(p).Interface().(dicts.Dict)
	obj.SetId(id)
	res := d.db.Model(obj).UpdateColumns(p)
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
