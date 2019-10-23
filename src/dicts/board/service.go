package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/handle"
	"bytes"
	"io"
)

type HandlerBoardService interface {
	FindBoardByID(id customType.StringUUID) (*Board, error)
	CreateBoard(request io.ReadCloser) (*Board, error)
	UpdateBoard(request io.ReadCloser, id customType.StringUUID) (*Board, error)
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

func (s service) CreateBoard(request io.ReadCloser) (*Board, error) {
	board := &Board{}
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(request); err != nil {
		return nil, err
	}
	if err := board.UnmarshalJSON(buf.Bytes()); err != nil {
		return nil, err
	}
	err := s.db.CreateRecord(board)
	return board, err
}

func (s service) UpdateBoard(request io.ReadCloser, id customType.StringUUID) (*Board, error) {
	board := &Board{}
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(request); err != nil {
		return nil, err
	}
	if err := board.UnmarshalJSON(buf.Bytes()); err != nil {
		return nil, err
	}
	err := s.db.UpdateRecord(board, id)
	return board, err
}

func (s service) DeleteBoard(id customType.StringUUID) error {
	board := Board{}
	return s.db.DeleteRecord(board, id)
}
