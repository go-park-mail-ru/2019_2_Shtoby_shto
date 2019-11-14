package card

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"github.com/pkg/errors"
)

type HandlerCardService interface {
	FindCardByID(id customType.StringUUID) (*models.Card, error)
	FetchCardsByIDs(ids []string) (cards []models.Card, err error)
	FetchCardsByCardGroupIDs(cardGroupIDs []string) (cards []models.Card, err error)
	CreateCard(data []byte) (*models.Card, error)
	UpdateCard(data []byte, id customType.StringUUID) (*models.Card, error)
	DeleteCard(id customType.StringUUID) error
	FetchCards(limit, offset int) (cards []models.Card, err error)
	FillLookupFields(card *models.Card, comments []models.Comment) error
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerCardService {
	return &service{
		db: db,
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
