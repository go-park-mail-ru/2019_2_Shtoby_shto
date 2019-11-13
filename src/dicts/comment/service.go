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
	FetchCommentsByCardIDs(cardIDs []string) (comments []models.Comment, err error)
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
	comment.Text = s.sanitizer.Sanitize(comment.Text)
	err := s.db.UpdateRecord(comment, id)
	return comment, err
}

func (s service) DeleteComment(id customType.StringUUID) error {
	comment := models.Comment{}
	return s.db.DeleteRecord(&comment, id)
}

func (s service) FetchComments(limit, offset int) (comments []models.Comment, err error) {
	_, err = s.db.FetchDict(&comments, "comments", limit, offset, nil, nil)
	return comments, err
}

func (s service) FetchCommentsByCardIDs(cardIDs []string) (comments []models.Comment, err error) {
	where := []string{"card_id = ?"}
	whereArgs := cardIDs
	_, err = s.db.FetchDict(&comments, "comments", 10000, 0, where, whereArgs)
	return comments, err
}
