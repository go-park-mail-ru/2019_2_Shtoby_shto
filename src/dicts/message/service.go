package message

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"github.com/pkg/errors"
)

type HandlerMessageService interface {
	FindMessageByID(id customType.StringUUID) (*models.Message, error)
	FetchMessagesByIDs(ids []string) (messages []models.Message, err error)
	FetchMessagesByMessageGroupIDs(messageGroupIDs []string) (messages []models.Message, err error)
	CreateMessage(data []byte) (*models.Message, error)
	UpdateMessage(data []byte, id customType.StringUUID) (*models.Message, error)
	DeleteMessage(id customType.StringUUID) error
	FetchMessages(limit, offset int) (messages []models.Message, err error)
	FillLookupFields(message *models.Message, comments []models.Comment) error
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerMessageService {
	return &service{
		db: db,
	}
}

func (s service) FindMessageByID(id customType.StringUUID) (*models.Message, error) {
	message := &models.Message{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(message)
	return message, err
}

func (s service) CreateMessage(data []byte) (*models.Message, error) {
	message := &models.Message{}
	if err := message.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if !message.IsValid() {
		return nil, errors.New("Message body is not valid")
	}
	err := s.db.CreateRecord(message)
	return message, err
}

func (s service) UpdateMessage(data []byte, id customType.StringUUID) (*models.Message, error) {
	message := &models.Message{}
	if err := message.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	message.ID = id
	//if !message.IsValid() {
	//	return nil, errors.New("Message body is not valid")
	//}
	err := s.db.UpdateRecord(message)
	return message, err
}

func (s service) DeleteMessage(id customType.StringUUID) error {
	message := &models.Message{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	return s.db.DeleteRecord(message)
}

func (s service) FetchMessages(limit, offset int) (messages []models.Message, err error) {
	messageModel := &models.Message{}
	_, err = s.db.FetchDict(&messages, messageModel, limit, offset)
	return messages, err
}

func (s service) FetchMessagesByIDs(ids []string) (messages []models.Message, err error) {
	where := []string{"id in (?)"}
	whereArgs := ids
	_, err = s.db.FetchDictBySlice(&messages, "messages", 10000, 0, where, whereArgs)
	return messages, err
}

func (s service) FetchMessagesByMessageGroupIDs(messageGroupIDs []string) (messages []models.Message, err error) {
	where := []string{"message_group_id in(?)"}
	whereArgs := messageGroupIDs
	_, err = s.db.FetchDictBySlice(&messages, "messages", 10000, 0, where, whereArgs)
	return messages, err
}

func (s service) FillLookupFields(message *models.Message, comments []models.Comment) error {
	return nil
}
