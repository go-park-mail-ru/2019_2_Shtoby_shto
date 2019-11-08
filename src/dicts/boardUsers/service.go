package boardUsers

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/handle"
	"errors"
)

type HandlerBoardUsersService interface {
	FindBoardUsersByIDs(userID, boardID customType.StringUUID) (*BoardUsers, error)
	CreateBoardUsers(boardUsersID, userID, boardID customType.StringUUID) (*BoardUsers, error)
	UpdateBoardUsers(userID, boardID customType.StringUUID, id customType.StringUUID) (*BoardUsers, error)
	DeleteBoardUsers(id customType.StringUUID) error
	FetchBoardUsersByUserID(userID customType.StringUUID) (boardUsers []BoardUsers, err error)
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

func (s service) CreateBoardUsers(boardUsersID, userID, boardID customType.StringUUID) (*BoardUsers, error) {
	boardUsers := &BoardUsers{
		BaseInfo: dicts.BaseInfo{
			ID: boardUsersID,
		},
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
	//if !boardUsers.IsValid() {
	//	return nil, errors.New("Board body is not valid")
	//}
	err := s.db.UpdateRecord(boardUsers, id)
	return boardUsers, err
}

func (s service) DeleteBoardUsers(id customType.StringUUID) error {
	boardUsers := BoardUsers{}
	return s.db.DeleteRecord(boardUsers, id)
}

func (s service) FetchBoardUsersByUserID(userID customType.StringUUID) (boardUsers []BoardUsers, err error) {
	where := []string{"user_id = ?"}
	whereArgs := []string{userID.String()}
	_, err = s.db.FetchDict(&boardUsers, "board_users", 10000, 0, where, whereArgs)
	return boardUsers, err
}
