package card

import (
	"2019_2_Shtoby_shto/file_service/file"
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"bufio"
	"bytes"
	"context"
	"github.com/pkg/errors"
	"time"
)

type HandlerCardService interface {
	FindCardByID(id customType.StringUUID) (*models.Card, error)
	FetchCardsByIDs(ids []string) (cards []models.Card, err error)
	FetchCardsByCardGroupIDs(cardGroupIDs []string) (cards []models.Card, err error)
	CreateCard(data []byte) (*models.Card, error)
	UpdateCard(data []byte, id customType.StringUUID) (*models.Card, error)
	DownloadFileToCard(file *bufio.Reader, cardID customType.StringUUID) (*models.Card, error)
	GetCardFile(cardID customType.StringUUID) ([]byte, error)
	DeleteCard(id customType.StringUUID) error
	FetchCards(limit, offset int) (cards []models.Card, err error)
	FillLookupFields(card *models.Card, comments []models.Comment) error
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
	fl file.IFileLoaderManagerClient
}

func CreateInstance(db database.IDataManager, fileLoader file.IFileLoaderManagerClient) HandlerCardService {
	return &service{
		db: db,
		fl: fileLoader,
	}
}

func (s service) FindCardByID(id customType.StringUUID) (*models.Card, error) {
	card := &models.Card{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(card)
	return card, err
}

func (s service) CreateCard(data []byte) (*models.Card, error) {
	card := &models.Card{}
	card.Deadline = time.Now();
	if err := card.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if !card.IsValid() {
		return nil, errors.New("Card body is not valid")
	}
	err := s.db.CreateRecord(card)
	return card, err
}

func (s service) UpdateCard(data []byte, id customType.StringUUID) (*models.Card, error) {
	card := &models.Card{}
	if err := card.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	card.ID = id
	//if !card.IsValid() {
	//	return nil, errors.New("Card body is not valid")
	//}
	err := s.db.UpdateRecord(card)
	return card, err
}

func (s service) DeleteCard(id customType.StringUUID) error {
	card := &models.Card{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	return s.db.DeleteRecord(card)
}

func (s service) FetchCards(limit, offset int) (cards []models.Card, err error) {
	cardModel := &models.Card{}
	_, err = s.db.FetchDict(&cards, cardModel, limit, offset)
	return cards, err
}

func (s service) FetchCardsByIDs(ids []string) (cards []models.Card, err error) {
	where := []string{"id in (?)"}
	whereArgs := ids
	_, err = s.db.FetchDictBySlice(&cards, "cards", 10000, 0, where, whereArgs)
	return cards, err
}

func (s service) FetchCardsByCardGroupIDs(cardGroupIDs []string) (cards []models.Card, err error) {
	where := []string{"card_group_id in(?)"}
	whereArgs := cardGroupIDs
	_, err = s.db.FetchDictBySlice(&cards, "cards", 10000, 0, where, whereArgs)
	return cards, err
}

func (s service) FillLookupFields(card *models.Card, comments []models.Comment) error {
	return nil
}

func (s service) DownloadFileToCard(reader *bufio.Reader, cardID customType.StringUUID) (*models.Card, error) {
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(reader); err != nil {
		return nil, err
	}
	newFile := &file.File{
		ID:   cardID.String(),
		Data: buf.Bytes(),
	}
	_, err := s.fl.DownloadFile(context.Background(), newFile)
	if err != nil {
		return nil, err
	}
	card := models.Card{
		File: cardID,
	}
	data, err := card.MarshalJSON()
	return s.UpdateCard(data, cardID)
}

func (s service) GetCardFile(cardID customType.StringUUID) ([]byte, error) {
	fileID := &file.FileID{
		ID: cardID.String(),
	}
	uploadFile, err := s.fl.UploadFile(context.Background(), fileID)
	if err != nil {
		return nil, err
	}
	return uploadFile.Data, nil
}
