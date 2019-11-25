package tag

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"github.com/microcosm-cc/bluemonday"
)

type HandlerTagService interface {
	FindTagByID(id customType.StringUUID) (*models.Tag, error)
	FetchTagsByIDs(tagIDs []string) (tags []models.Tag, err error)
	CreateTag(data []byte) (*models.Tag, error)
	UpdateTag(data []byte, id customType.StringUUID) (*models.Tag, error)
	DeleteTag(id customType.StringUUID) error
	FetchTags(limit, offset int) (tags []models.Tag, err error)
}

type service struct {
	handle.HandlerImpl
	db        database.IDataManager
	sanitizer *bluemonday.Policy
}

func CreateInstance(db database.IDataManager) HandlerTagService {
	sanitizer := bluemonday.UGCPolicy()
	return &service{
		db:        db,
		sanitizer: sanitizer,
	}
}

func (s service) FindTagByID(id customType.StringUUID) (*models.Tag, error) {
	tag := &models.Tag{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(tag)
	return tag, err
}

func (s service) CreateTag(data []byte) (*models.Tag, error) {
	tag := &models.Tag{}
	if err := tag.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	tag.Text = s.sanitizer.Sanitize(tag.Text)
	err := s.db.CreateRecord(tag)
	return tag, err
}

func (s service) UpdateTag(data []byte, id customType.StringUUID) (*models.Tag, error) {
	tag := &models.Tag{}
	if err := tag.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	tag.ID = id
	tag.Text = s.sanitizer.Sanitize(tag.Text)
	err := s.db.UpdateRecord(tag)
	return tag, err
}

func (s service) DeleteTag(id customType.StringUUID) error {
	tag := &models.Tag{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	return s.db.DeleteRecord(tag)
}

func (s service) FetchTags(limit, offset int) (tags []models.Tag, err error) {
	tag := &models.Tag{}
	_, err = s.db.FetchDict(&tags, tag, limit, offset)
	return tags, err
}

func (s service) FetchTagsByIDs(tagIDs []string) (tags []models.Tag, err error) {
	where := []string{"id in(?)"}
	whereArgs := tagIDs
	_, err = s.db.FetchDictBySlice(&tags, "tags", 10000, 0, where, whereArgs)
	return tags, err
}
