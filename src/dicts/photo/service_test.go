package photo

import (
	"2019_2_Shtoby_shto/file_service/file"
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/utils"
	"bufio"
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"reflect"
	"testing"
	"time"
)

func TestCreateInstance(t *testing.T) {
	type args struct {
		db         database.IDataManager
		cfg        *config.Config
		fileLoader file.IFileLoaderManagerClient
	}
	tests := []struct {
		name string
		args args
		want HandlerPhotoService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateInstance(tt.args.db, tt.args.cfg, tt.args.fileLoader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DownloadPhoto(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	downloadPhotoServiceMock := NewMockHandlerPhotoService(gomock.NewController(t))
	newdatabase := database.NewMockIDataManager(gomock.NewController(t))
	newIFileLoader := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})

	downloadPhotoService := CreateInstance(newdatabase, &config.Config{}, newIFileLoader)

	//gDB, err := gorm.Open("postgres", db)
	//assert.Nil(t, err, "db error")

	uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "generate uuid error")

	downloadPhotoServiceMock.EXPECT().DownloadPhoto(uuid.String()).Return(nil, nil)
	newPhoto := &models.Photo{
		BaseInfo: dicts.BaseInfo{
			ID: customType.StringUUID(uuid.String()),
		},
		TimeLoad: time.Time{},
		Path:     "/",
	}

	data, _ := newPhoto.MarshalJSON()

	downloadPhotoServiceMock.EXPECT().DownloadPhoto(data)
	buf := bytes.NewReader(data)
	r := bufio.NewReader(buf)
	_, err = downloadPhotoService.DownloadPhoto(r)
	assert.EqualError(t, err, "mkdir : The system cannot find the path specified.")
}

func Test_service_GetPhotoByUser(t *testing.T) {
	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	GetPhotoByUserMock := NewMockHandlerPhotoService(gomock.NewController(t))
	newdatabase := database.NewMockIDataManager(gomock.NewController(t))
	newIFileLoader := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})

	GetPhotoByUser := CreateInstance(newdatabase, &config.Config{}, newIFileLoader)

	//gDB, err := gorm.Open("postgres", db)
	//assert.Nil(t, err, "db error")

	uuid, err := utils.GenerateUUID()
	assert.Nil(t, err, "generate uuid error")

	GetPhotoByUserMock.EXPECT().GetPhotoByUser(uuid.String()).Return(nil, nil)

	newPhoto := &models.Photo{
		BaseInfo: dicts.BaseInfo{
			ID: customType.StringUUID(uuid.String()),
		},
		TimeLoad: time.Time{},
		Path:     "/",
	}

	data, _ := newPhoto.MarshalJSON()

	GetPhotoByUserMock.EXPECT().DownloadPhoto(data)
	buf := bytes.NewReader(data)
	r := bufio.NewReader(buf)
	_, err = GetPhotoByUser.DownloadPhoto(r)
	assert.EqualError(t, err, "mkdir : The system cannot find the path specified.")

	type fields struct {
		HandlerImpl handle.HandlerImpl
		cfg         *config.Config
		db          database.IDataManager
		fl          file.IFileLoaderManagerClient
	}
	type args struct {
		photoID customType.StringUUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				HandlerImpl: tt.fields.HandlerImpl,
				cfg:         tt.fields.cfg,
				db:          tt.fields.db,
				fl:          tt.fields.fl,
			}
			got, err := s.GetPhotoByUser(tt.args.photoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPhotoByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPhotoByUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
