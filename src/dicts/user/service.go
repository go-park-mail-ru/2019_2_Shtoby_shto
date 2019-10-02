package user

import (
	. "2019_2_Shtoby_shto/src/custom_type"
	"2019_2_Shtoby_shto/src/database"
)

type HandlerUserService interface {
	CreateUser(user User) error
	UpdateUser(user User, id StringUUID) error
	GetUserById(id StringUUID) (User, error)
	GetUserByLogin(login string) (User, error)
}

type service struct {
	db *database.DataManager
}

func CreateInstance(db *database.DataManager) HandlerUserService {
	return &service{
		db: db,
	}
}

func (s *service) CreateUser(user User) error {
	//err := s.db.ExecuteQuery("insert into users(id, login, password) values($1, $2, $3)", user.ID.String(), user.Login, user.Password)
	return s.db.CreateRecord(&user)
}

func (s *service) GetUserById(id StringUUID) (User, error) {
	user := User{}
	user.ID = id
	err := s.db.FindDictById(&user)
	return user, err
}

func (s *service) GetUserByLogin(login string) (User, error) {
	user := User{}
	err := s.db.FindDictByLogin(&user, "login", login)
	return user, err
}

func (s *service) UpdateUser(user User, id StringUUID) error {
	if err := s.db.UpdateRecord(&user, id); err != nil {
		return err
	}
	return nil
}
