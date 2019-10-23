package user

import (
	. "2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"github.com/pkg/errors"
	"strings"
)

type HandlerUserService interface {
	CreateUser(data []byte) (*User, error)
	UpdateUser(data []byte, id StringUUID) error
	GetUserById(id StringUUID) (User, error)
	GetUserByLogin(data []byte) (*User, error)
}

type service struct {
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerUserService {
	return &service{
		db: db,
	}
}

func (s *service) CreateUser(data []byte) (*User, error) {
	user := &User{}
	if err := user.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	err := s.db.CreateRecord(user)
	return user, err
}

func (s *service) GetUserById(id StringUUID) (User, error) {
	user := User{}
	user.ID = id
	err := s.db.FindDictById(&user)
	return user, err
}

func (s *service) GetUserByLogin(data []byte) (*User, error) {
	curUser := User{}
	if err := curUser.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	user := &User{}
	err := s.db.FindDictByColumn(user, "login", curUser.Login)
	if err != nil {
		return nil, err
	}
	if strings.Compare(user.Password, curUser.Password) != 0 {
		return nil, errors.New("Ne tot password )0))")
	}
	return user, err
}

func (s *service) UpdateUser(data []byte, id StringUUID) error {
	user := User{}
	if err := user.UnmarshalJSON(data); err != nil {
		return err
	}
	if !user.IsValid() {
		return errors.New("User not valid!")
	}
	if err := s.db.UpdateRecord(&user, id); err != nil {
		return err
	}
	return nil
}
