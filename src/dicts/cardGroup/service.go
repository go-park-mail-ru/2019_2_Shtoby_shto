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
	FetchCardGroupsByBoardIDs(boardID []string) (cardGroups []CardGroup, err error)
	CreateCardGroup(data []byte) (*CardGroup, error)
	UpdateCardGroup(data []byte, id customType.StringUUID) (*CardGroup, error)
	DeleteCardGroup(id customType.StringUUID) error
	FetchCardGroup(limit, offset int) (cardGroup []CardGroup, err error)
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
	//if !card.IsValid() {
	//	return nil, errors.New("CardGroup body is not valid")
	//}
	err := s.db.UpdateRecord(card, id)
	return card, err
}

func (s service) DeleteCardGroup(id customType.StringUUID) error {
	card := CardGroup{}
	return s.db.DeleteRecord(&card, id)
}

func (s service) FetchCardGroup(limit, offset int) (cardGroup []CardGroup, err error) {
	_, err = s.db.FetchDict(&cardGroup, "card_groups", limit, offset, nil, nil)
	return cardGroup, err
}

func (s service) FetchCardGroupsByBoardIDs(boardID []string) (cardGroups []CardGroup, err error) {
	where := []string{"board_id in(?)"}
	whereArgs := boardID
	_, err = s.db.FetchDict(&cardGroups, "card_groups", 10000, 0, where, whereArgs)
	return cardGroups, err
}
