package card

import (
	"2019_2_Shtoby_shto/file_service/file"
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/utils"
	//"bufio"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"log"
	"reflect"
	"testing"
)

func CreateCardServiceTest(t *testing.T) HandlerCardService {
	db, _, err := sqlmock.New()
	if !assert.Nilf(t, err,
		"error on trying to start") {
		t.Fatal()
	}

	gDB, err := gorm.Open("postgres", db)
	if err != nil {
		log.Fatal(err)
	}
	dm := database.NewDataManager(gDB)
	fm := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	cardService := CreateInstance(dm, fm)
	return cardService
}

func TestCreateInstance(t *testing.T) {
	type args struct {
		db         database.IDataManager
		fileLoader file.IFileLoaderManagerClient
	}
	tests := []struct {
		name string
		args args
		want HandlerCardService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateInstance(tt.args.db, tt.args.fileLoader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateCard(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can not create mock SQL: %s", err)
	}

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	fileManager := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	createCard := CreateInstance(dataManager, fileManager)

	createCardMock := NewMockHandlerCardService(gomock.NewController(t))
	assert.Nil(t, err, "can not generate uuid")

	createCardMock.EXPECT().CreateCard("").Return(&models.Card{}, nil)

	newCard := &models.Card{
		BaseInfo: dicts.BaseInfo{
			ID: "14",
		},
		Caption:     "",
		Text:        "",
		Priority:    0,
		FileID:      "",
		CardUserID:  "",
		CardGroupID: "",
		File:        "",
		Comments:    nil,
		Tags:        nil,
		Users:       nil,
	}

	data, _ := newCard.MarshalJSON()

	createCardMock.EXPECT().CreateCard(data).Return(newCard, nil)
	_, err = createCard.CreateCard(data)
	assert.EqualError(t, err, "Card body is not valid")
}

func Test_service_DeleteCard(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can not create mock SQL: %s", err)
	}

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	fileManager := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	deleteCard := CreateInstance(dataManager, fileManager)

	deleteCardMock := NewMockHandlerCardService(gomock.NewController(t))
	assert.Nil(t, err, "can not generate uuid")

	deleteCardMock.EXPECT().DeleteCard(nil).Return(nil)

	err = deleteCard.DeleteCard("14")
	assert.EqualError(t, err, "all expectations were already fulfilled, call to database transaction Begin was not expected")
}

func Test_service_DownloadFileToCard(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can not create mock SQL: %s", err)
	}

	uploadFileMock := NewMockHandlerCardService(gomock.NewController(t))
	assert.Nil(t, err, "can not generate uuid")

	uploadFileMock.EXPECT().DownloadFileToCard(nil, nil).Return(&models.Card{}, nil)

	newFile := &file.File{
		ID:                   "55",
		Data:                 nil,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	newCard := &models.Card{
		BaseInfo: dicts.BaseInfo{
			ID: "41",
		},
		Caption:     "",
		Text:        "",
		Priority:    0,
		FileID:      "",
		CardUserID:  "",
		CardGroupID: "",
		File:        "",
		Comments:    nil,
		Tags:        nil,
		Users:       nil,
	}

	uploadFileMock.EXPECT().DownloadFileToCard(newFile, newCard.ID).Return(newCard, nil)
}

func Test_service_FetchCards(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	fetchCardsMock := NewMockHandlerCardService(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	fileManager := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	fetchCards := CreateInstance(dataManager, fileManager)

	//uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "can not generate uuid")

	fetchCardsMock.EXPECT().FetchCards(nil, nil).Return([]models.Card{}, nil)

	newCard := models.Card{
		BaseInfo: dicts.BaseInfo{
			ID: "41",
		},
		Caption:     "",
		Text:        "",
		Priority:    0,
		FileID:      "",
		CardUserID:  "",
		CardGroupID: "",
		File:        "",
		Comments:    nil,
		Tags:        nil,
		Users:       nil,
	}

	fetchCardsMock.EXPECT().FetchCards(1, 41).Return([]models.Card{newCard}, nil)
	_, err = fetchCards.FetchCards(1, 41)
	assert.EqualError(t, err, "all expectations were already fulfilled, call to Query 'SELECT * FROM \"cards\"   LIMIT 1 OFFSET 41' with args [] was not expected; all expectations were already fulfilled, call to Query 'SELECT count(*) FROM \"cards\"   LIMIT 1 OFFSET 41' with args [] was not expected")

}

func Test_service_FetchCardsByCardGroupIDs(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	fetchCardsByCardGroupIDsMock := NewMockHandlerCardService(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	fileManager := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	fetchCardGroupByCardGroupIDs := CreateInstance(dataManager, fileManager)

	fetchCardsByCardGroupIDsMock.EXPECT().FetchCardsByCardGroupIDs(nil).Return([]models.Card{}, nil)

	newCard := models.Card{
		BaseInfo: dicts.BaseInfo{
			ID: "41",
		},
		Caption:     "",
		Text:        "",
		Priority:    0,
		FileID:      "",
		CardUserID:  "",
		CardGroupID: "",
		File:        "",
		Comments:    nil,
		Tags:        nil,
		Users:       nil,
	}

	fetchCardsByCardGroupIDsMock.EXPECT().FetchCardsByCardGroupIDs(21).Return([]models.Card{newCard}, nil)
	_, err = fetchCardGroupByCardGroupIDs.FetchCards(1, 41)
	assert.EqualError(t, err, "all expectations were already fulfilled, call to Query 'SELECT * FROM \"cards\"   LIMIT 1 OFFSET 41' with args [] was not expected; all expectations were already fulfilled, call to Query 'SELECT count(*) FROM \"cards\"   LIMIT 1 OFFSET 41' with args [] was not expected")
}

func Test_service_FetchCardsByIDs(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	fetchCardsByIDsMock := NewMockHandlerCardService(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	fileManager := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	fetchCardGroupByIDs := CreateInstance(dataManager, fileManager)

	fetchCardsByIDsMock.EXPECT().FetchCardsByIDs(nil).Return([]models.Card{}, nil)

	newCard := models.Card{
		BaseInfo: dicts.BaseInfo{
			ID: "41",
		},
		Caption:     "",
		Text:        "",
		Priority:    0,
		FileID:      "",
		CardUserID:  "",
		CardGroupID: "",
		File:        "",
		Comments:    nil,
		Tags:        nil,
		Users:       nil,
	}

	fetchCardsByIDsMock.EXPECT().FetchCardsByIDs(21).Return([]models.Card{newCard}, nil)
	_, err = fetchCardGroupByIDs.FetchCards(1, 41)
	assert.EqualError(t, err, "all expectations were already fulfilled, call to Query 'SELECT * FROM \"cards\"   LIMIT 1 OFFSET 41' with args [] was not expected; all expectations were already fulfilled, call to Query 'SELECT count(*) FROM \"cards\"   LIMIT 1 OFFSET 41' with args [] was not expected")
}

func Test_service_FillLookupFields(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	fillLookupFieldsMock := NewMockHandlerCardService(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	fileManager := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	fillLookipFields := CreateInstance(dataManager, fileManager)

	//uuid, err := utils.GenerateUUID()
	//assert.Nil(t, err, "can not generate uuid")

	//newCard:=models.Card{
	//	BaseInfo: dicts.BaseInfo{
	//		ID: customType.StringUUID(uuid.String()),
	//	},
	//	Caption:     "",
	//	Text:        "",
	//	Priority:    0,
	//	FileID:      "",
	//	CardUserID:  "",
	//	CardGroupID: "",
	//	File:        "",
	//	Comments:    nil,
	//	Tags:        nil,
	//	Users:       nil,
	//}

	fillLookupFieldsMock.EXPECT().FillLookupFields("", "").Return(nil)
	err = fillLookipFields.FillLookupFields(&models.Card{}, []models.Comment{})
	assert.NoError(t, err)
}

func Test_service_FindCardByID(t *testing.T) {
	type fields struct {
		HandlerImpl handle.HandlerImpl
		db          database.IDataManager
		fl          file.IFileLoaderManagerClient
	}
	type args struct {
		id customType.StringUUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Card
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				HandlerImpl: tt.fields.HandlerImpl,
				db:          tt.fields.db,
				fl:          tt.fields.fl,
			}
			got, err := s.FindCardByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindCardByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCardByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetCardFile(t *testing.T) {
	//db, _, err := sqlmock.New()
	//defer db.Close()
	//
	//if err!=nil{
	//	t.Fatalf("can not create mock SQL: %s", err)
	//}
	//
	//getCardFileMock := NewMockHandlerCardService(gomock.NewController(t))
	//
	//gDB, err := gorm.Open("postgres", db)
	//assert.Nil(t, err, "db error")
	//dataManager := database.NewDataManager(gDB)
	//fileManager := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	//getCardFile := CreateInstance(dataManager, fileManager)
	//
	//uuid, err := utils.GenerateUUID()
	//assert.Nil(t, err, "can not generate uuid")
	//
	//getCardFileMock.EXPECT().GetCardFile(nil).Return("", nil)
	//
	////TODO: do download file to card first
	//
	//
	//type fields struct {
	//	HandlerImpl handle.HandlerImpl
	//	db          database.IDataManager
	//	fl          file.IFileLoaderManagerClient
	//}
	//type args struct {
	//	cardID customType.StringUUID
	//}
	//tests := []struct {
	//	name    string
	//	fields  fields
	//	args    args
	//	want    []byte
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		s := service{
	//			HandlerImpl: tt.fields.HandlerImpl,
	//			db:          tt.fields.db,
	//			fl:          tt.fields.fl,
	//		}
	//		got, err := s.GetCardFile(tt.args.cardID)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("GetCardFile() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("GetCardFile() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func Test_service_UpdateCard(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can not create mock SQL: %s", err)
	}

	updateCardMock := NewMockHandlerCardService(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	fileManager := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	updateCard := CreateInstance(dataManager, fileManager)

	uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "can not generate uuid")

	updateCardMock.EXPECT().UpdateCard(&file.File{}, "14").Return(&models.Card{}, nil)

	newFile := &file.File{
		ID:                   "14",
		Data:                 nil,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	updatedCard := &models.Card{
		BaseInfo: dicts.BaseInfo{
			ID: customType.StringUUID(uuid.String()),
		},
		Caption:     "",
		Text:        "lala",
		Priority:    0,
		FileID:      "",
		CardUserID:  "",
		CardGroupID: "",
		File:        "",
		Comments:    nil,
		Tags:        nil,
		Users:       nil,
	}

	data, _ := updatedCard.MarshalJSON()

	updateCardMock.EXPECT().UpdateCard(&newFile, "14")
	_, err = updateCard.CreateCard(data)
	assert.EqualError(t, err, "Card body is not valid")

	updateCardMock.EXPECT().UpdateCard(&newFile, "14").Return(updatedCard, nil)
	_, err = updateCard.CreateCard(data)
	assert.EqualError(t, err, "Card body is not valid")
}
