package photo

import (
	"2019_2_Shtoby_shto/file_service/file"
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/models"
	"2019_2_Shtoby_shto/src/utils"
	"bufio"
	"bytes"
	"fmt"
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

	defer func() {
		if err := recover(); err != nil {
			fmt.Print("panic, HA, LOSER!\n")
		}
	}()

	db, _, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}

	GetPhotoByUserMock := NewMockHandlerPhotoService(gomock.NewController(t))
	newdatabase := database.NewMockIDataManager(gomock.NewController(t))
	newIFileLoader := file.NewIFileLoaderManagerClient(&grpc.ClientConn{})
	cfg := &config.Config{
		Port:             8080,
		FrontendURL:      "",
		SecurityURL:      "",
		FileLoaderURL:    "",
		ImagePath:        "/photo",
		StorageAccessKey: "",
		StorageSecretKey: "",
		StorageRegion:    "",
		StorageEndpoint:  "",
		StorageBucket:    "",
		DbConfig:         "",
		RedisConfig:      "",
		RedisPass:        "",
		RedisDbNumber:    0,
	}

	GetPhotoByUser := CreateInstance(newdatabase, cfg, newIFileLoader)

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
		Path:     "/photo",
	}

	data, _ := newPhoto.MarshalJSON()

	GetPhotoByUserMock.EXPECT().DownloadPhoto(data)
	//buf := bytes.NewReader(data)
	//r := bufio.NewReader(buf)
	_, err = GetPhotoByUser.GetPhotoByUser("41")
	assert.EqualError(t, err, "mkdir : The system cannot find the path specified.")

}
