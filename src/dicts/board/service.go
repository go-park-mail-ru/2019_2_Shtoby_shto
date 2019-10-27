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
	CreateBoard(data []byte) (*Board, error)
	UpdateBoard(data []byte, id customType.StringUUID) (*Board, error)
	DeleteBoard(id customType.StringUUID) error
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

func (s service) CreateBoard(data []byte) (*Board, error) {
	board := &Board{}
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
	if !board.IsValid() {
		return nil, errors.New("Board body is not valid")
	}
	err := s.db.UpdateRecord(board, id)
	return board, err
}

func (s service) DeleteBoard(id customType.StringUUID) error {
	board := Board{}
	return s.db.DeleteRecord(board, id)
}
