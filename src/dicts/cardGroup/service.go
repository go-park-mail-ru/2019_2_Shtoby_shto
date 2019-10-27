package cardGroup

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/handle"
	"github.com/pkg/errors"
)

type HandlerCardGroupService interface {
	FindCardGroupByID(id customType.StringUUID) (*CardGroup, error)
	CreateCardGroup(data []byte) (*CardGroup, error)
	UpdateCardGroup(data []byte, id customType.StringUUID) (*CardGroup, error)
	DeleteCardGroup(id customType.StringUUID) error
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerCardGroupService {
	return &service{
		db: db,
	}
}

func (s service) FindCardGroupByID(id customType.StringUUID) (*CardGroup, error) {
	card := &CardGroup{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(card)
	return card, err
}

func (s service) CreateCardGroup(data []byte) (*CardGroup, error) {
	card := &CardGroup{}
	if err := card.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if !card.IsValid() {
		return nil, errors.New("CardGroup body is not valid")
	}
	err := s.db.CreateRecord(card)
	return card, err
}

func (s service) UpdateCardGroup(data []byte, id customType.StringUUID) (*CardGroup, error) {
	card := &CardGroup{}
	if err := card.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if !card.IsValid() {
		return nil, errors.New("CardGroup body is not valid")
	}
	err := s.db.UpdateRecord(card, id)
	return card, err
}

func (s service) DeleteCardGroup(id customType.StringUUID) error {
	card := CardGroup{}
	return s.db.DeleteRecord(&card, id)
}
