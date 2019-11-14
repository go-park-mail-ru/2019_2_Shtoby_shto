package boardUsers

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"errors"
)

type HandlerBoardUsersService interface {
	CreateBoardUsers(boardUsersID, userID, boardID customType.StringUUID) (*models.BoardUsers, error)
	FindBoardUsersByIDs(userID, boardID customType.StringUUID) (int, error)
	UpdateBoardUsers(userID, boardID customType.StringUUID, id customType.StringUUID) (*models.BoardUsers, error)
	DeleteBoardUsers(id customType.StringUUID) error
	FetchBoardUsersByUserID(userID customType.StringUUID) (boardUsers []models.BoardUsers, err error)
	FetchBoardUsersByBoardID(boardID customType.StringUUID) (boardUsers []models.BoardUsers, err error)
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

func (s service) FindBoardUsersByIDs(userID, boardID customType.StringUUID) (int, error) {
	boardUsers := &models.BoardUsers{
		BoardID: boardID,
		UserID:  userID,
	}
	count, err := s.db.FindDictByColumn(boardUsers)
	if count == 0 {
		return count, nil
	}
	return count, err
}

func (s service) CreateBoardUsers(boardUsersID, userID, boardID customType.StringUUID) (*models.BoardUsers, error) {
	boardUsers := &models.BoardUsers{
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

func (s service) UpdateBoardUsers(userID, boardID customType.StringUUID, id customType.StringUUID) (*models.BoardUsers, error) {
	boardUsers := &models.BoardUsers{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
		UserID:  userID,
		BoardID: boardID,
	}
	//if !boardUsers.IsValid() {
	//	return nil, errors.New("Board body is not valid")
	//}
	err := s.db.UpdateRecord(boardUsers)
	return boardUsers, err
}

func (s service) DeleteBoardUsers(id customType.StringUUID) error {
	boardUsers := &models.BoardUsers{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	return s.db.DeleteRecord(boardUsers)
}

func (s service) FetchBoardUsersByUserID(userID customType.StringUUID) (boardUsers []models.BoardUsers, err error) {
	where := []string{"user_id = ?"}
	whereArgs := []string{userID.String()}
	_, err = s.db.FetchDict(&boardUsers, "board_users", 10000, 0, where, whereArgs)
	return boardUsers, err
}

func (s service) FetchBoardUsersByBoardID(boardID customType.StringUUID) (boardUsers []models.BoardUsers, err error) {
	where := []string{"board_id = ?"}
	whereArgs := []string{boardID.String()}
	_, err = s.db.FetchDict(&boardUsers, "board_users", 10000, 0, where, whereArgs)
	return boardUsers, err
}
