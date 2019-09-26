package user

import (
	"2019_2_Shtoby_shto/src/database"
)

type UserHandler interface {
	PutUser(user User) error
	GetUser(id string) (*User, error)
}

type service struct {
	db database.DataManager
}

func CreateInstance(db database.DataManager) UserHandler {
	return &service{
		db: db,
	}
}

func (s *service) PutUser(user User) error {
	err := s.db.ExecuteQueryUser("insert into users(id, login, password) values($1, $2, $3)", user.ID, user.Login, user.Password)
	return err
}

func (s *service) GetUser(id string) (*User, error) {
	user := &User{}
	err := s.db.FindDictById("select id, login, password from users where id = $1", user)
	return user, err
}
