package user

import (
	. "2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"bytes"
	"github.com/gyepisam/mcf/pbkdf2"
	"github.com/pkg/errors"
)

type HandlerUserService interface {
	CreateUser(data []byte) (*User, error)
	UpdateUser(data []byte, id StringUUID) error
	GetUserById(id StringUUID) (User, error)
	GetUserByLogin(data []byte) (*User, error)
	FetchUsers(limit, offset int) (users []User, err error)
}

type service struct {
	db  database.IDataManager
	cfg pbkdf2.Config
}

func CreateInstance(db database.IDataManager) HandlerUserService {
	cfg := pbkdf2.Config{
		Hash:       pbkdf2.SHA256,
		KeyLen:     pbkdf2.SHA256.Size(),
		SaltLen:    pbkdf2.DefaultSaltLen,
		Iterations: pbkdf2.DefaultIterations,
	}
	return &service{
		db:  db,
		cfg: cfg,
	}
}

func (s *service) CreateUser(data []byte) (*User, error) {
	user := &User{}
	if err := user.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if err := s.setPasswordPBKDF2(user); err != nil {
		return nil, err
	}
	err := s.db.CreateRecord(user)
	return user, err
}

func (s service) setPasswordPBKDF2(user *User) error {
	salt, err := s.cfg.Salt()
	if err != nil {
		return err
	}
	user.Salt = salt
	passCrypt, err := s.getCryptPass(user.Password, salt)
	if err != nil {
		return err
	}
	user.PasswordCrypt = passCrypt
	return nil
}

func (s service) getCryptPass(password string, salt []byte) ([]byte, error) {
	passCrypt, err := s.cfg.Key([]byte(password), salt)
	if err != nil {
		return nil, err
	}
	return passCrypt, nil
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
	where := []string{"login = ?"}
	whereArgs := []string{curUser.Login}
	err := s.db.FindDictByColumn(user, where, whereArgs)
	if err != nil {
		return nil, err
	}
	pass, err := s.getCryptPass(curUser.Password, user.Salt)
	if err != nil {
		return nil, errors.New("Crypt error")
	}
	if bytes.Compare(pass, user.PasswordCrypt) != 0 {
		return nil, errors.New("Ne tot password )0))")
	}
	return user, err
}

func (s *service) UpdateUser(data []byte, id StringUUID) error {
	user := &User{}
	if err := user.UnmarshalJSON(data); err != nil {
		return err
	}
	if !user.IsValid() {
		return errors.New("User not valid!")
	}
	if err := s.setPasswordPBKDF2(user); err != nil {
		return err
	}
	if err := s.db.UpdateRecord(user, id); err != nil {
		return err
	}
	return nil
}

func (s service) FetchUsers(limit, offset int) (users []User, err error) {
	_, err = s.db.FetchDict(&users, "users", limit, offset, nil, nil)
	return users, err
}
