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
	type fields struct {
		HandlerImpl handle.HandlerImpl
		db          database.IDataManager
		fl          file.IFileLoaderManagerClient
	}
	type args struct {
		data []byte
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
			got, err := s.CreateCard(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteCard(t *testing.T) {
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
			if err := s.DeleteCard(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
	type fields struct {
		HandlerImpl handle.HandlerImpl
		db          database.IDataManager
		fl          file.IFileLoaderManagerClient
	}
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCards []models.Card
		wantErr   bool
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
			gotCards, err := s.FetchCards(tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCards, tt.wantCards) {
				t.Errorf("FetchCards() gotCards = %v, want %v", gotCards, tt.wantCards)
			}
		})
	}
}

func Test_service_FetchCardsByCardGroupIDs(t *testing.T) {
	type fields struct {
		HandlerImpl handle.HandlerImpl
		db          database.IDataManager
		fl          file.IFileLoaderManagerClient
	}
	type args struct {
		cardGroupIDs []string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCards []models.Card
		wantErr   bool
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
			gotCards, err := s.FetchCardsByCardGroupIDs(tt.args.cardGroupIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchCardsByCardGroupIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCards, tt.wantCards) {
				t.Errorf("FetchCardsByCardGroupIDs() gotCards = %v, want %v", gotCards, tt.wantCards)
			}
		})
	}
}

func Test_service_FetchCardsByIDs(t *testing.T) {
	type fields struct {
		HandlerImpl handle.HandlerImpl
		db          database.IDataManager
		fl          file.IFileLoaderManagerClient
	}
	type args struct {
		ids []string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCards []models.Card
		wantErr   bool
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
			gotCards, err := s.FetchCardsByIDs(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchCardsByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCards, tt.wantCards) {
				t.Errorf("FetchCardsByIDs() gotCards = %v, want %v", gotCards, tt.wantCards)
			}
		})
	}
}

func Test_service_FillLookupFields(t *testing.T) {
	type fields struct {
		HandlerImpl handle.HandlerImpl
		db          database.IDataManager
		fl          file.IFileLoaderManagerClient
	}
	type args struct {
		card     *models.Card
		comments []models.Comment
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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
			if err := s.FillLookupFields(tt.args.card, tt.args.comments); (err != nil) != tt.wantErr {
				t.Errorf("FillLookupFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
