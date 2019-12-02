package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/utils"
	"errors"
)

//go:generate mockgen -source=$GOFILE -destination=service_mock.go -package=$GOPACKAGE

type HandlerBoardService interface {
	FindBoardByID(id customType.StringUUID) (*models.Board, error)
	FindBoardByURL(shortURL string) (*models.Board, error)
	CreateBoard(data []byte, boardUserID customType.StringUUID) (*models.Board, error)
	UpdateBoard(data []byte, id customType.StringUUID) (*models.Board, error)
	DeleteBoard(id customType.StringUUID) error
	FetchBoards(limit, offset int) (boards []models.Board, err error)
	FetchBoardsByIDs(boardsIDs []string) (boards []models.Board, err error)
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

func (s service) FindBoardByID(id customType.StringUUID) (*models.Board, error) {
	board := &models.Board{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(board)
	return board, err
}

func (s service) FindBoardByURL(shortURL string) (*models.Board, error) {
	board := &models.Board{
		ShortURL: shortURL,
	}
	count, err := s.db.FindDictByColumn(board)
	if count == 0 {
		return nil, errors.New("Short URL is not valid! ")
	}
	return board, err
}

func (s service) CreateBoard(data []byte, boardUserID customType.StringUUID) (*models.Board, error) {

	board := &models.Board{
		BoardUsersID: boardUserID,
		ShortURL:     utils.RandStringBytesRmndr(5),
	}
	if err := board.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if !board.IsValid() {
		return nil, errors.New("Board body is not valid ")
	}
	err := s.db.CreateRecord(board)
	return board, err
}

func (s service) UpdateBoard(data []byte, id customType.StringUUID) (*models.Board, error) {
	board := &models.Board{}
	if err := board.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	board.ID = id
	err := s.db.UpdateRecord(board)
	return board, err
}

func (s service) DeleteBoard(id customType.StringUUID) error {
	board := &models.Board{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	return s.db.DeleteRecord(board)
}

func (s service) FetchBoards(limit, offset int) (boards []models.Board, err error) {
	boardModel := &models.Board{}
	_, err = s.db.FetchDict(&boards, boardModel, limit, offset)
	return boards, err
}

func (s service) FetchBoardsByIDs(boardsIDs []string) (boards []models.Board, err error) {
	where := []string{"id in(?)"}
	_, err = s.db.FetchDictBySlice(&boards, "boards", 10000, 0, where, boardsIDs)
	return boards, err
}
