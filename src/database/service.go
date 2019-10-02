package database

import (
	"github.com/jinzhu/gorm"
)

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
