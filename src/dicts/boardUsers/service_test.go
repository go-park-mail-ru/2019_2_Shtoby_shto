package boardUsers

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"
)

func CreateBoardUsersServiceTest(t *testing.T) HandlerBoardUsersService {
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
	boardUsersService := CreateInstance(dm)
	return boardUsersService
}

func Test_service_CreateBoardUsers(t *testing.T) {
	boardUsersService := CreateBoardUsersServiceTest(t)

	type fields struct {
		HandlerImpl handle.HandlerImpl
		db          database.IDataManager
	}
	type args struct {
		boardUsersID customType.StringUUID
		userID       customType.StringUUID
		boardID      customType.StringUUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.BoardUsers
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				HandlerImpl: handle.HandlerImpl{},
				db:          nil,
			},
			args: args{
				boardUsersID: "50",
				userID:       "42",
				boardID:      "41",
			},
			want: &models.BoardUsers{
				BaseInfo: dicts.BaseInfo{
					ID: "50",
				},
				BoardID: "41",
				UserID:  "42",
			},
			wantErr: true,
		}, {
			name: "2",
			fields: fields{
				HandlerImpl: handle.HandlerImpl{},
				db:          nil,
			},
			args: args{
				boardUsersID: "51",
				userID:       "43",
				boardID:      "42",
			},
			want: &models.BoardUsers{
				BaseInfo: dicts.BaseInfo{
					ID: "51",
				},
				BoardID: "42",
				UserID:  "43",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//s := service{
			//	HandlerImpl: tt.fields.HandlerImpl,
			//	db:          tt.fields.db,
			//}
			got, err := boardUsersService.CreateBoardUsers(tt.args.boardUsersID, tt.args.userID, tt.args.boardID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBoardUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBoardUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteBoardUsers(t *testing.T) {
	type fields struct {
		HandlerImpl handle.HandlerImpl
		db          database.IDataManager
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
			}
			if err := s.DeleteBoardUsers(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBoardUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteBoardUsersByIDs(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	DeleteBoardUsersByIDMock := NewMockHandlerBoardUsersService(gomock.NewController(t))

	//gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	//dataManager := database.NewDataManager(gDB)
	//deleteBoardUsers := CreateInstance(dataManager)

	uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "can not generate uuid")

	//TODO: create board with users and destroy it

	DeleteBoardUsersByIDMock.EXPECT().DeleteBoardUsersByIDs(uuid.String(), uuid.String()).Return(nil)

	DeleteBoardUsersByIDMock.EXPECT().DeleteBoardUsersByIDs(42, 41).Return(nil)
	//err = deleteBoardUsers.DeleteBoardUsersByIDs(customType.StringUUID(42),customType.StringUUID(41))
	//assert.EqualError(t, err, "nil")

	DeleteBoardUsersByIDMock.EXPECT().DeleteBoardUsersByIDs(customType.StringUUID(41), customType.StringUUID(42)).Return(nil)

	//err = deleteBoardUsers.DeleteBoardUsersByIDs(customType.StringUUID(42),customType.StringUUID(41))
	//assert.EqualError(t, err, "nil")

}

func Test_service_FetchBoardUsersByBoardID(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	fetchBoardUsersByBoardIDMock := NewMockHandlerBoardUsersService(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	fetchBoardUsers := CreateInstance(dataManager)

	uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "can not generate uuid")

	fetchBoardUsersByBoardIDMock.EXPECT().FetchBoardUsersByBoardID(uuid.String()).Return([]models.BoardUsers{}, nil)

	fetchBoardUsersByBoardIDMock.EXPECT().FetchBoardUsersByBoardID(41).Return([]models.BoardUsers{}, nil)
	_, err = fetchBoardUsers.FetchBoardUsersByUserID(customType.StringUUID(41))
	assert.EqualError(t, err, "all expectations were already fulfilled, call to Query 'SELECT * FROM \"board_users\"  WHERE (\"board_users\".\"user_id\" = $1) LIMIT 10000 OFFSET 0' with args [{Name: Ordinal:1 Value:)}] was not expected; all expectations were already fulfilled, call to Query 'SELECT count(*) FROM \"board_users\"  WHERE (\"board_users\".\"user_id\" = $1) LIMIT 10000 OFFSET 0' with args [{Name: Ordinal:1 Value:)}] was not expected")

	fetchBoardUsersByBoardIDMock.EXPECT().FetchBoardUsersByBoardID(uuid.String()).Return([]models.BoardUsers{}, nil)

	_, err = fetchBoardUsers.FetchBoardUsersByUserID(customType.StringUUID(41))
	assert.EqualError(t, err, "all expectations were already fulfilled, call to Query 'SELECT * FROM \"board_users\"  WHERE (\"board_users\".\"user_id\" = $1) LIMIT 10000 OFFSET 0' with args [{Name: Ordinal:1 Value:)}] was not expected; all expectations were already fulfilled, call to Query 'SELECT count(*) FROM \"board_users\"  WHERE (\"board_users\".\"user_id\" = $1) LIMIT 10000 OFFSET 0' with args [{Name: Ordinal:1 Value:)}] was not expected")

}

func Test_service_FetchBoardUsersByUserID(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	fetchBoardUsersByUserIDMock := NewMockHandlerBoardUsersService(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	fetchBoardUsers := CreateInstance(dataManager)

	uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "can not generate uuid")

	fetchBoardUsersByUserIDMock.EXPECT().FetchBoardUsersByUserID(uuid.String()).Return([]models.BoardUsers{}, nil)

	fetchBoardUsersByUserIDMock.EXPECT().FetchBoardUsersByUserID(41).Return([]models.BoardUsers{}, nil)
	_, err = fetchBoardUsers.FetchBoardUsersByUserID(customType.StringUUID(41))
	assert.EqualError(t, err, "all expectations were already fulfilled, call to Query 'SELECT * FROM \"board_users\"  WHERE (\"board_users\".\"user_id\" = $1) LIMIT 10000 OFFSET 0' with args [{Name: Ordinal:1 Value:)}] was not expected; all expectations were already fulfilled, call to Query 'SELECT count(*) FROM \"board_users\"  WHERE (\"board_users\".\"user_id\" = $1) LIMIT 10000 OFFSET 0' with args [{Name: Ordinal:1 Value:)}] was not expected")

	fetchBoardUsersByUserIDMock.EXPECT().FetchBoardUsersByUserID(uuid.String()).Return([]models.BoardUsers{}, nil)

	_, err = fetchBoardUsers.FetchBoardUsersByUserID(customType.StringUUID(41))
	assert.EqualError(t, err, "all expectations were already fulfilled, call to Query 'SELECT * FROM \"board_users\"  WHERE (\"board_users\".\"user_id\" = $1) LIMIT 10000 OFFSET 0' with args [{Name: Ordinal:1 Value:)}] was not expected; all expectations were already fulfilled, call to Query 'SELECT count(*) FROM \"board_users\"  WHERE (\"board_users\".\"user_id\" = $1) LIMIT 10000 OFFSET 0' with args [{Name: Ordinal:1 Value:)}] was not expected")

}

func Test_service_FindBoardUsersByIDs(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	boardUsersMock := NewMockHandlerBoardUsersService(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	boardUsers := CreateInstance(dataManager)

	uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "can not generate uuid")

	boardUsersMock.EXPECT().FindBoardUsersByIDs(uuid.String(), 42).Return(0, nil)

	boardUsersMock.EXPECT().CreateBoardUsers(42, 41, "")
	_, err = boardUsers.CreateBoardUsers(customType.StringUUID(42), customType.StringUUID(41), "")
	assert.EqualError(t, err, "Board body is not valid")

	boardUsersMock.EXPECT().FindBoardUsersByIDs(42, 41).Return(42, nil)
	_, err = boardUsers.CreateBoardUsers(customType.StringUUID(42), customType.StringUUID(41), "")
	assert.EqualError(t, err, "Board body is not valid")

}

func Test_service_UpdateBoardUsers(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can not create mock SQL: %s", err)
	}

	boardUsersMock := NewMockHandlerBoardUsersService(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	updateBoardUsers := CreateInstance(dataManager)

	uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "can not generate uuid")

	boardUsersMock.EXPECT().UpdateBoardUsers(uuid.String(), uuid.String(), "").Return(&models.BoardUsers{}, nil)

	updatedBoard := &models.BoardUsers{
		BaseInfo: dicts.BaseInfo{
			ID: customType.StringUUID(uuid.String()),
		},
		BoardID: "41",
		UserID:  "42",
	}

	boardUsersMock.EXPECT().CreateBoardUsers(42, 41, "")
	_, err = updateBoardUsers.CreateBoardUsers(customType.StringUUID(42), customType.StringUUID(41), "")
	assert.EqualError(t, err, "Board body is not valid")

	boardUsersMock.EXPECT().UpdateBoardUsers(42, 41, "").Return(updatedBoard, nil)
	_, err = updateBoardUsers.CreateBoardUsers(customType.StringUUID(42), customType.StringUUID(41), "")
	assert.EqualError(t, err, "Board body is not valid")
}
