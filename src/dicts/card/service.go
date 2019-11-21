package card

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/fileLoader"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/utils"
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
)

//go:generate mockgen -source=$GOFILE -destination=usecase_mock.go -package=$GOPACKAGE

type HandlerCardService interface {
	FindCardByID(id customType.StringUUID) (*models.Card, error)
	FetchCardsByIDs(ids []string) (cards []models.Card, err error)
	FetchCardsByCardGroupIDs(cardGroupIDs []string) (cards []models.Card, err error)
	CreateCard(data []byte) (*models.Card, error)
	UpdateCard(data []byte, id customType.StringUUID) (*models.Card, error)
	DeleteCard(id customType.StringUUID) error
	FetchCards(limit, offset int) (cards []models.Card, err error)
	FillLookupFields(card *models.Card, comments []models.Comment) error
	AttachFile(cardID customType.StringUUID, file *bufio.Reader) (*models.Card, error)
	GetCardFile(fileID customType.StringUUID) ([]byte, error)
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
	fl fileLoader.IFileLoaderManager
}

func CreateInstance(db database.IDataManager, manager fileLoader.IFileLoaderManager) HandlerCardService {
	return &service{
		db: db,
		fl: manager,
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

func (s service) AttachFile(cardID customType.StringUUID, file *bufio.Reader) (*models.Card, error) {
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}
	card := &models.Card{
		BaseInfo: dicts.BaseInfo{
			ID: cardID,
		},
		FileID: customType.StringUUID(uuid.String()),
	}
	cardData, err := card.MarshalJSON()
	if err != nil {
		return nil, err
	}
	newCard, err := s.UpdateCard(cardData, cardID)
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(file); err != nil {
		return nil, err
	}
	err = s.fl.DownloadFile(uuid.String(), buf.Bytes())
	if err != nil {
		return nil, err
	}
	return newCard, nil
}

func (s service) GetCardFile(fileID customType.StringUUID) ([]byte, error) {
	file, err := s.fl.UploadFile(fileID.String())
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}
