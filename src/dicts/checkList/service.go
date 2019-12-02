package checkList

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"github.com/microcosm-cc/bluemonday"
)

type HandlerCheckListService interface {
	FindCheckListByID(id customType.StringUUID) (*models.CheckList, error)
	FetchCheckListsByCardID(cardID string) (checkLists []models.CheckList, err error)
	CreateCheckList(data []byte) (*models.CheckList, error)
	UpdateCheckList(data []byte, id customType.StringUUID) (*models.CheckList, error)
	DeleteCheckList(id customType.StringUUID) error
	FetchCheckLists(limit, offset int) (checkLists []models.CheckList, err error)
}

type service struct {
	handle.HandlerImpl
	db        database.IDataManager
	sanitizer *bluemonday.Policy
}

func CreateInstance(db database.IDataManager) HandlerCheckListService {
	sanitizer := bluemonday.UGCPolicy()
	return &service{
		db:        db,
		sanitizer: sanitizer,
	}
}

func (s service) FindCheckListByID(id customType.StringUUID) (*models.CheckList, error) {
	checkList := &models.CheckList{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(checkList)
	return checkList, err
}

func (s service) CreateCheckList(data []byte) (*models.CheckList, error) {
	checkList := &models.CheckList{}
	if err := checkList.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	checkList.Text = s.sanitizer.Sanitize(checkList.Text)
	err := s.db.CreateRecord(checkList)
	return checkList, err
}

func (s service) UpdateCheckList(data []byte, id customType.StringUUID) (*models.CheckList, error) {
	checkList := &models.CheckList{}
	if err := checkList.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	checkList.ID = id
	checkList.Text = s.sanitizer.Sanitize(checkList.Text)
	err := s.db.UpdateRecord(checkList)
	return checkList, err
}

func (s service) DeleteCheckList(id customType.StringUUID) error {
	checkList := &models.CheckList{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	return s.db.DeleteRecord(checkList)
}

func (s service) FetchCheckLists(limit, offset int) (checkLists []models.CheckList, err error) {
	checkList := &models.CheckList{}
	_, err = s.db.FetchDict(&checkLists, checkList, limit, offset)
	return checkLists, err
}

func (s service) FetchCheckListsByCardID(cardID string) (checkLists []models.CheckList, err error) {
	checkList := &models.CheckList{
		CardID: cardID,
	}
	_, err = s.db.FetchDict(&checkLists, checkList, 10000, 0)
	return checkLists, err
}
