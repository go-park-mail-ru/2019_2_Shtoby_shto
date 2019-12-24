package сardUsers

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"errors"
)

type HandlerCardUsersService interface {
	FindCardUsersByID(id customType.StringUUID) (*models.CardUsers, error)
	FindCardUsersByUserID(userData []byte) ([]models.CardUsers, error)
	FindCardUsersByIDs(userID, cardID customType.StringUUID) (int, error)
	CreateCardUsers(userID, сardID customType.StringUUID) (*models.CardUsers, error)
	UpdateCardUsers(userID, сardID customType.StringUUID, id customType.StringUUID) (*models.CardUsers, error)
	DeleteCardUsers(id customType.StringUUID) error
	DeleteCardUsersByIDs(userID, сardID customType.StringUUID) error
	FetchCardUsersByCardID(cardID customType.StringUUID) (cardUsers []models.CardUsers, err error)
	FetchUserIDsByCardID(cardID customType.StringUUID) (users map[string]struct{}, err error)
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerCardUsersService {
	return &service{
		db: db,
	}
}

func (s *service) FindCardUsersByUserID(userData []byte) (cardUsers []models.CardUsers, err error) {
	userIDs := models.CardsUserRequest{}
	if err = userIDs.UnmarshalJSON(userData); err != nil {
		return nil, err
	}
	if len(userIDs.Users) == 0 {
		return nil, errors.New("User id is empty! ")
	}
	where := []string{"user_id in (?)"}
	whereArgs := userIDs.Users
	_, err = s.db.FetchDictBySlice(&cardUsers, "card_users", 10000, 0, where, whereArgs)
	return cardUsers, err
}

func (s *service) FindCardUsersByID(id customType.StringUUID) (*models.CardUsers, error) {
	сardUsers := &models.CardUsers{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(сardUsers)
	return сardUsers, err
}

func (s *service) FindCardUsersByIDs(userID, cardID customType.StringUUID) (int, error) {
	cardUsers := &models.CardUsers{
		CardID: cardID,
		UserID: userID,
	}
	count, err := s.db.FindDictByColumn(cardUsers)
	if count == 0 {
		return count, nil
	}
	return count, err
}

func (s *service) CreateCardUsers(userID, сardID customType.StringUUID) (*models.CardUsers, error) {
	сardUsers := &models.CardUsers{
		CardID: сardID,
		UserID: userID,
	}
	if !сardUsers.IsValid() {
		return nil, errors.New("Board body is not valid")
	}
	err := s.db.CreateRecord(сardUsers)
	return сardUsers, err
}

func (s *service) UpdateCardUsers(userID, сardID customType.StringUUID, id customType.StringUUID) (*models.CardUsers, error) {
	boardUsers := &models.CardUsers{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
		UserID: userID,
		CardID: сardID,
	}
	//if !boardUsers.IsValid() {
	//	return nil, errors.New("Board body is not valid")
	//}
	err := s.db.UpdateRecord(boardUsers)
	return boardUsers, err
}

func (s *service) DeleteCardUsers(id customType.StringUUID) error {
	cardUsers := &models.CardUsers{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	return s.db.DeleteRecord(cardUsers)
}

func (s *service) FetchCardUsers(limit, offset int) (cardUsers []models.CardUsers, err error) {
	cardUser := &models.CardUsers{}
	_, err = s.db.FetchDict(&cardUsers, cardUser, limit, offset)
	return cardUsers, err
}

func (s *service) FetchCardUsersByCardID(cardID customType.StringUUID) (cardUsers []models.CardUsers, err error) {
	cardUser := &models.CardUsers{
		CardID: cardID,
	}
	_, err = s.db.FetchDict(&cardUsers, cardUser, 10000, 0)
	return cardUsers, err
}

func (s *service) DeleteCardUsersByIDs(userID, сardID customType.StringUUID) error {
	cardUsers := &models.CardUsers{
		UserID: userID,
		CardID: сardID,
	}
	err := s.db.DeleteRecord(cardUsers)
	if err != nil {

	}
	return nil
}

func (s *service) FetchUserIDsByCardID(cardID customType.StringUUID) (userIDs map[string]struct{}, err error) {
	userIDs = make(map[string]struct{}, 0)
	cUsers := make([]models.CardUsers, 0)

	where := []string{"card_id in (?)"}
	whereArgs := []string{cardID.String()}
	_, err = s.db.FetchDictBySlice(&cUsers, "card_users", 10000, 0, where, whereArgs)
	for _, cardUser := range cUsers {
		userIDs[cardUser.UserID.String()] = struct{}{}
	}
	return userIDs, nil
}
