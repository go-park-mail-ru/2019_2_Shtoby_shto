package photo

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"bufio"
	"reflect"
	"testing"
)

func TestCreateInstance(t *testing.T) {
	type args struct {
		db database.IDataManager
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
			if got := CreateInstance(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DownloadPhoto(t *testing.T) {
	type args struct {
		photo *bufio.Reader
	}
	tests := []struct {
		name    string
		s       service
		args    args
		want    customType.StringUUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DownloadPhoto(tt.args.photo)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.DownloadPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.DownloadPhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetPhotoByUser(t *testing.T) {
	type args struct {
		photoID customType.StringUUID
	}
	tests := []struct {
		name    string
		s       service
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetPhotoByUser(tt.args.photoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetPhotoByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetPhotoByUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
