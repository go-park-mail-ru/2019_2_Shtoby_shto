package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/handle"
	"errors"
)

type HandlerBoardService interface {
	FindBoardByID(id customType.StringUUID) (*Board, error)
	CreateBoard(data []byte, boardUserID customType.StringUUID) (*Board, error)
	UpdateBoard(data []byte, id customType.StringUUID) (*Board, error)
	DeleteBoard(id customType.StringUUID) error
	FetchBoards(limit, offset int) (boards []Board, err error)
	FetchBoardsByIDs(boardsIDs []string) (boards []Board, err error)
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerBoardService {
	return &service{
		db: db,
	}
}

func (s service) FindBoardByID(id customType.StringUUID) (*Board, error) {
	board := &Board{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(board)
	return board, err
}

func (s service) CreateBoard(data []byte, boardUserID customType.StringUUID) (*Board, error) {
	board := &Board{
		BoardUsersID: boardUserID,
	}
	if err := board.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if !board.IsValid() {
		return nil, errors.New("Board body is not valid")
	}
	err := s.db.CreateRecord(board)
	return board, err
}

func (s service) UpdateBoard(data []byte, id customType.StringUUID) (*Board, error) {
	board := &Board{}
	if err := board.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	//if !board.IsValid() {
	//	return nil, errors.New("Board body is not valid")
	//}
	err := s.db.UpdateRecord(board, id)
	return board, err
}

func (s service) DeleteBoard(id customType.StringUUID) error {
	board := &Board{}
	return s.db.DeleteRecord(board, id)
}

func (s service) FetchBoards(limit, offset int) (boards []Board, err error) {
	_, err = s.db.FetchDict(&boards, "boards", limit, offset, nil, nil)
	return boards, err
}

func (s service) FetchBoardsByIDs(boardsIDs []string) (boards []Board, err error) {
	where := []string{"id in(?)"}
	_, err = s.db.FetchDict(&boards, "boards", 10000, 0, where, boardsIDs)
	return boards, err
}
