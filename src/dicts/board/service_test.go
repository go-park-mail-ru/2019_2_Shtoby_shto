package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/models"
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
	type args struct {
		id customType.StringUUID
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
			got, err := tt.s.FindBoardByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.FindBoardByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.FindBoardByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateBoard(t *testing.T) {
	type args struct {
		data []byte
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
			got, err := tt.s.CreateBoard(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateBoard() = %v, want %v", got, tt.want)
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
