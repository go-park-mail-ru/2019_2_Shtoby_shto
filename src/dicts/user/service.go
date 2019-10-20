package user

import (
	. "2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"github.com/pkg/errors"
)

type HandlerUserService interface {
	CreateUser(user User) error
	UpdateUser(user User, id StringUUID) error
	GetUserById(id StringUUID) (User, error)
	GetUserByLogin(login string) (User, error)
}

type service struct {
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerUserService {
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
	err := s.db.FindDictByColumn(&user, "login", login)
	return user, err
}

func (s *service) UpdateUser(user User, id StringUUID) error {
	if !user.IsValid() {
		return errors.New("User not valid!")
	}
	if err := s.db.UpdateRecord(&user, id); err != nil {
		return err
	}
	return nil
}
