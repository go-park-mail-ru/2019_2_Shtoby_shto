package boardUsers

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/handle"
	"errors"
)

type HandlerBoardUsersService interface {
	FindBoardUsersByID(id customType.StringUUID) (*BoardUsers, error)
	CreateBoardUsers(userID, boardID string) (*BoardUsers, error)
	UpdateBoardUsers(userID, boardID string, id customType.StringUUID) (*BoardUsers, error)
	DeleteBoardUsers(id customType.StringUUID) error
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerBoardUsersService {
	return &service{
		db: db,
	}
}

func (s service) FindBoardUsersByID(id customType.StringUUID) (*BoardUsers, error) {
	boardUsers := &BoardUsers{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(boardUsers)
	return boardUsers, err
}

func (s service) CreateBoardUsers(userID, boardID string) (*BoardUsers, error) {
	boardUsers := &BoardUsers{
		UserID:  userID,
		BoardID: boardID,
	}
	if !boardUsers.IsValid() {
		return nil, errors.New("Board body is not valid")
	}
	err := s.db.CreateRecord(boardUsers)
	return boardUsers, err
}

func (s service) UpdateBoardUsers(userID, boardID string, id customType.StringUUID) (*BoardUsers, error) {
	boardUsers := &BoardUsers{
		UserID:  userID,
		BoardID: boardID,
	}
	if !boardUsers.IsValid() {
		return nil, errors.New("Board body is not valid")
	}
	err := s.db.UpdateRecord(boardUsers, id)
	return boardUsers, err
}

func (s service) DeleteBoardUsers(id customType.StringUUID) error {
	boardUsers := BoardUsers{}
	return s.db.DeleteRecord(boardUsers, id)
}
