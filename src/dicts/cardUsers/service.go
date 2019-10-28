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
	FindCardUsersByUserID(userData []byte) ([]CardUsers, error)
	CreateCardUsers(userID, сardID customType.StringUUID) (*CardUsers, error)
	UpdateCardUsers(userID, сardID customType.StringUUID, id customType.StringUUID) (*CardUsers, error)
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

func (s service) FindCardUsersByUserID(userData []byte) (cardUsers []CardUsers, err error) {
	userIDs := CardsUserRequest{}
	if err = userIDs.UnmarshalJSON(userData); err != nil {
		return nil, err
	}
	if len(userIDs.Users) == 0 {
		return nil, errors.New("User id is empty! ")
	}
	where := []string{"user_id in (?)"}
	whereArgs := userIDs.Users
	_, err = s.db.FetchDict(&cardUsers, "card_users", 10000, 0, where, whereArgs)
	return cardUsers, err
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

func (s service) CreateCardUsers(userID, сardID customType.StringUUID) (*CardUsers, error) {
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

func (s service) UpdateCardUsers(userID, сardID customType.StringUUID, id customType.StringUUID) (*CardUsers, error) {
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

func (s service) FetchCardUsers(limit, offset int) (cardUsers []CardUsers, err error) {
	_, err = s.db.FetchDict(&cardUsers, "card_users", limit, offset, nil, nil)
	return cardUsers, err
}