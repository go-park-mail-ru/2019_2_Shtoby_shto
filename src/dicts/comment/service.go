package comment

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"github.com/microcosm-cc/bluemonday"
)

type HandlerCommentService interface {
	FindCommentByID(id customType.StringUUID) (*models.Comment, error)
	FetchCommentsByCardID(cardID string) (comments []models.Comment, err error)
	CreateComment(data []byte) (*models.Comment, error)
	UpdateComment(data []byte, id customType.StringUUID) (*models.Comment, error)
	DeleteComment(id customType.StringUUID) error
	FetchComments(limit, offset int) (comments []models.Comment, err error)
}

type service struct {
	handle.HandlerImpl
	db        database.IDataManager
	sanitizer *bluemonday.Policy
}

func CreateInstance(db database.IDataManager) HandlerCommentService {
	sanitizer := bluemonday.UGCPolicy()
	return &service{
		db:        db,
		sanitizer: sanitizer,
	}
}

func (s service) FindCommentByID(id customType.StringUUID) (*models.Comment, error) {
	comment := &models.Comment{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(comment)
	return comment, err
}

func (s service) CreateComment(data []byte) (*models.Comment, error) {
	comment := &models.Comment{}
	if err := comment.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	comment.Text = s.sanitizer.Sanitize(comment.Text)
	err := s.db.CreateRecord(comment)
	return comment, err
}

func (s service) UpdateComment(data []byte, id customType.StringUUID) (*models.Comment, error) {
	comment := &models.Comment{}
	if err := comment.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	comment.ID = id
	comment.Text = s.sanitizer.Sanitize(comment.Text)
	err := s.db.UpdateRecord(comment)
	return comment, err
}

func (s service) DeleteComment(id customType.StringUUID) error {
	comment := &models.Comment{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	return s.db.DeleteRecord(comment)
}

func (s service) FetchComments(limit, offset int) (comments []models.Comment, err error) {
	comment := &models.Comment{}
	_, err = s.db.FetchDict(&comments, comment, limit, offset)
	return comments, err
}

func (s service) FetchCommentsByCardID(cardID string) (comments []models.Comment, err error) {
	comment := &models.Comment{
		CardID: cardID,
	}
	_, err = s.db.FetchDict(&comments, comment, 10000, 0)
	return comments, err
}
