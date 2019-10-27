package boardUsers

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/handle"
	"errors"
)

type HandlerBoardUsersService interface {
	FindBoardUsersByIDs(userID, boardID customType.StringUUID) (*BoardUsers, error)
	CreateBoardUsers(userID, boardID customType.StringUUID) (*BoardUsers, error)
	UpdateBoardUsers(userID, boardID customType.StringUUID, id customType.StringUUID) (*BoardUsers, error)
	DeleteBoardUsers(id customType.StringUUID) error
	FetchBoardUsers(limit, offset int) (users []BoardUsers, err error)
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

func (s service) FindBoardUsersByIDs(userID, boardID customType.StringUUID) (*BoardUsers, error) {
	boardUsers := &BoardUsers{}
	where := []string{"user_id = ?", "board_id = ?"}
	whereArgs := []string{userID.String(), boardID.String()}
	err := s.db.FindDictByColumn(boardUsers, where, whereArgs)
	return boardUsers, err
}

func (s service) CreateBoardUsers(userID, boardID customType.StringUUID) (*BoardUsers, error) {
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

func (s service) UpdateBoardUsers(userID, boardID customType.StringUUID, id customType.StringUUID) (*BoardUsers, error) {
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

func (s service) FetchBoardUsers(limit, offset int) (boardUsers []BoardUsers, err error) {
	_, err = s.db.FetchDict(&boardUsers, "board_users", limit, offset, nil, nil)
	return boardUsers, err
}
