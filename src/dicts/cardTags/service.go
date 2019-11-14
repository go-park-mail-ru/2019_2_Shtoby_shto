package cardTags

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"errors"
)

type HandlerCardTagsService interface {
	FindCardTagsByID(id customType.StringUUID) (*models.CardTags, error)
	FindCardTagsByCardID(cardID string) (cardTags []models.CardTags, err error)
	CreateCardTags(tagID, сardID customType.StringUUID) (*models.CardTags, error)
	UpdateCardTags(tagID, сardID customType.StringUUID, id customType.StringUUID) (*models.CardTags, error)
	DeleteCardTags(id customType.StringUUID) error
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerCardTagsService {
	return &service{
		db: db,
	}
}

func (s service) FindCardTagsByCardID(cardID string) (cardTags []models.CardTags, err error) {
	where := []string{"card_id in (?)"}
	whereArgs := []string{cardID}
	_, err = s.db.FetchDict(&cardTags, "card_tags", 10000, 0, where, whereArgs)
	return cardTags, err
}

func (s service) FindCardTagsByID(id customType.StringUUID) (*models.CardTags, error) {
	сardTags := &models.CardTags{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(сardTags)
	return сardTags, err
}

func (s service) CreateCardTags(tagID, сardID customType.StringUUID) (*models.CardTags, error) {
	сardTags := &models.CardTags{
		CardID: сardID,
		TagID:  tagID,
	}
	if !сardTags.IsValid() {
		return nil, errors.New("Board body is not valid")
	}
	err := s.db.CreateRecord(сardTags)
	return сardTags, err
}

func (s service) UpdateCardTags(tagID, сardID customType.StringUUID, id customType.StringUUID) (*models.CardTags, error) {
	cardTags := &models.CardTags{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
		TagID:  tagID,
		CardID: сardID,
	}
	//if !cardTags.IsValid() {
	//	return nil, errors.New("Board body is not valid")
	//}
	err := s.db.UpdateRecord(cardTags)
	return cardTags, err
}

func (s service) DeleteCardTags(id customType.StringUUID) error {
	cardTags := &models.CardTags{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	return s.db.DeleteRecord(cardTags)
}

func (s service) FetchCardTags(limit, offset int) (cardTags []models.CardTags, err error) {
	_, err = s.db.FetchDict(&cardTags, "card_tags", limit, offset, nil, nil)
	return cardTags, err
}
