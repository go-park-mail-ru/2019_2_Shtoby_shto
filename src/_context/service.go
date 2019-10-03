package _context

import (
	"2019_2_Shtoby_shto/src/database"
	"log"
)

type Context struct {
	Dm      database.DataManager
	Session database.Session
	Logger  log.Logger
}

func (c *Context) Clone() (res *Context) {
	res = &Context{Dm: c.Dm, Session: c.Session, Logger: c.Logger}
	return res
}

func (c *Context) DataManager() *database.DataManager {
	return &c.Dm
}
