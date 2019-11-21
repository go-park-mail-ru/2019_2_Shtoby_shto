package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func CreateBoardServiceTest(t *testing.T) HandlerBoardService {
	db, _, err := sqlmock.New()
	if !assert.Nilf(t,
		err,
		"Error on trying to start up a stub database connection") {
		t.Fatal()
	}

	gDB, err := gorm.Open("postgres", db)
	if err != nil {
		log.Fatal(err)
	}
	dm := database.NewDataManager(gDB)
	boardService := CreateInstance(dm)
	return boardService
}

func TestCreateInstance(t *testing.T) {
	type args struct {
		db database.IDataManager
	}
	tests := []struct {
		name string
		args args
		want HandlerBoardService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateInstance(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_FindBoardByID(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}

	boardServiceMock := NewMockHandlerBoardService(gomock.NewController(t))
	//dataManagerMock := database.NewMockIDataManager(gomock.NewController(t))

	gDB, err := gorm.Open("postgres", db)
	assert.Nil(t, err, "db error")
	dataManager := database.NewDataManager(gDB)
	boardService := CreateInstance(dataManager)

	uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "generate uuid error")
	//findThisID := []customType.StringUUID {
	//	//"5f8b4635-ed70-4d5c-be62-2a59eb79b7a8",
	//	"9cebe93e-d43b-4cc0-8c2a-40ac4992dc7d",
	//}
	boardServiceMock.EXPECT().FindBoardByID(uuid.String()).Return(nil, nil)
	newBoard := &models.Board{
		BaseInfo: dicts.BaseInfo{
			ID: customType.StringUUID(uuid.String()),
		},
		Name: "New test board",
	}
	data, _ := newBoard.MarshalJSON()

	boardServiceMock.EXPECT().CreateBoard(data, "")
	_, err = boardService.CreateBoard(data, "")
	assert.EqualError(t, err, "all expectations were already fulfilled, call to database transaction Begin was not expected")

	searchBoard := &models.Board{}
	boardServiceMock.EXPECT().FindBoardByID(uuid.String()).Return(searchBoard, nil)

	_, err = boardService.CreateBoard(data, "")
	assert.EqualError(t, err, "all expectations were already fulfilled, call to database transaction Begin was not expected")

	//type args struct {
	//	id customType.StringUUID
	//}
	//tests := []struct {
	//	name    string
	//	s       service
	//	args    args
	//	want    *models.Board
	//	wantErr bool
	//}{
	//	{
	//
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, err := tt.s.FindBoardByID(tt.args.id)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("service.FindBoardByID() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("service.FindBoardByID() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func Test_service_CreateBoard(t *testing.T) {
	boardService := CreateBoardServiceTest(t)

	tests := []struct {
		name     string
		board    *models.Board
		wantName string
		wantErr  bool
	}{
		{
			name: "1",
			board: &models.Board{
				Name: "new board 1",
			},
			wantName: "new board 1",
			wantErr:  false,
		},
		{
			name: "2",
			board: &models.Board{
				BaseInfo: dicts.BaseInfo{
					ID: "9cebe93e-d43b-4cc0-8c2a-40ac4992dc7d",
				},
				Name: "new board 1",
			},
			wantName: "new board 1",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.board.ID.IsEmpty() {
				uuid, err := utils.GenerateUUID()
				if err != nil {
					t.Error(err)
				}
				tt.board.ID = customType.StringUUID(uuid.String())
			}
			data, err := tt.board.MarshalJSON()
			if err != nil {
				t.Errorf("Marshal error")
			}
			got, err := boardService.CreateBoard(data, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.wantName) {
				t.Errorf("service.CreateBoard() = %v, want %v", got.Name, tt.wantName)
			}
		})
	}
}

func Test_service_UpdateBoard(t *testing.T) {
	type args struct {
		data []byte
		id   customType.StringUUID
	}
	tests := []struct {
		name    string
		s       service
		args    args
		want    *models.Board
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateBoard(tt.args.data, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.UpdateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteBoard(t *testing.T) {
	type args struct {
		id customType.StringUUID
	}
	tests := []struct {
		name    string
		s       service
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeleteBoard(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteBoard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
