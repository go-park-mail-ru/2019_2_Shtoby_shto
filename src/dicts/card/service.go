package card

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/handle"
)

type HandlerCardService interface {
	FindCardByID(id customType.StringUUID) (*Card, error)
	CreateCard(data []byte) (*Card, error)
	UpdateCard(data []byte, id customType.StringUUID) (*Card, error)
	DeleteCard(id customType.StringUUID) error
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

func (s service) FindCardByID(id customType.StringUUID) (*Card, error) {
	card := &Card{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(card)
	return card, err
}

func (s service) CreateCard(data []byte) (*Card, error) {
	card := &Card{}
	if err := card.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	err := s.db.CreateRecord(card)
	return card, err
}

func (s service) UpdateCard(data []byte, id customType.StringUUID) (*Card, error) {
	card := &Card{}
	if err := card.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	err := s.db.UpdateRecord(card, id)
	return card, err
}

func (s service) DeleteCard(id customType.StringUUID) error {
	card := Card{}
	return s.db.DeleteRecord(&card, id)
}
