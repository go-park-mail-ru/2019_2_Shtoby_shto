package сardUsers

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/handle"
	"errors"
)

type HandlerCardUsersService interface {
	FindCardUsersByID(id customType.StringUUID) (*CardUsers, error)
	CreateCardUsers(userID, сardID string) (*CardUsers, error)
	UpdateCardUsers(userID, сardID string, id customType.StringUUID) (*CardUsers, error)
	DeleteCardUsers(id customType.StringUUID) error
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerCardUsersService {
	return &service{
		db: db,
	}
}

func (s service) FindCardUsersByID(id customType.StringUUID) (*CardUsers, error) {
	сardUsers := &CardUsers{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(сardUsers)
	return сardUsers, err
}

func (s service) CreateCardUsers(userID, сardID string) (*CardUsers, error) {
	сardUsers := &CardUsers{
		CardID: сardID,
		UserID: userID,
	}
	if !сardUsers.IsValid() {
		return nil, errors.New("Board body is not valid")
	}
	err := s.db.CreateRecord(сardUsers)
	return сardUsers, err
}

func (s service) UpdateCardUsers(userID, сardID string, id customType.StringUUID) (*CardUsers, error) {
	boardUsers := &CardUsers{
		UserID: userID,
		CardID: сardID,
	}
	if !boardUsers.IsValid() {
		return nil, errors.New("Board body is not valid")
	}
	err := s.db.UpdateRecord(boardUsers, id)
	return boardUsers, err
}

func (s service) DeleteCardUsers(id customType.StringUUID) error {
	boardUsers := CardUsers{}
	return s.db.DeleteRecord(boardUsers, id)
}
