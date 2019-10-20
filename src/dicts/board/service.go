package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/handle"
)

type HandlerBoardService interface {
	FindBoardByID(id customType.StringUUID) (*Board, error)
	CreateBoard(board *Board) error
	UpdateBoard(board *Board, id customType.StringUUID) error
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

func (s service) CreateBoard(board *Board) error {
	return s.db.CreateRecord(board)
}

func (s service) UpdateBoard(board *Board, id customType.StringUUID) error {
	return s.db.UpdateRecord(board, id)
}

func (s service) DeleteBoard(id customType.StringUUID) error {
	board := Board{}
	return s.db.DeleteRecord(board, id)
}
