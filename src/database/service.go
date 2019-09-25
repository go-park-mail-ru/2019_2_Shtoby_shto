package database

import (
	"github.com/jinzhu/gorm"
)

type DictHandler interface {
	IsValid() bool
	TableName() string
}

type Session struct {
	ID string
}

type Context struct {
	db *gorm.DB
	Session
}

func (c *Context) GetSessionID() string {
	return c.ID
}
