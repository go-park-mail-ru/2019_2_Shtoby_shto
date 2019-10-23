package user

import (
	. "2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
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
		want HandlerUserService
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

func Test_service_CreateUser(t *testing.T) {
	type fields struct {
		db database.IDataManager
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				db: tt.fields.db,
			}
			got, err := s.CreateUser(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetUserById(t *testing.T) {
	type fields struct {
		db database.IDataManager
	}
	type args struct {
		id StringUUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				db: tt.fields.db,
			}
			got, err := s.GetUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetUserByLogin(t *testing.T) {
	type fields struct {
		db database.IDataManager
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				db: tt.fields.db,
			}
			got, err := s.GetUserByLogin(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetUserByLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetUserByLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateUser(t *testing.T) {
	type fields struct {
		db database.IDataManager
	}
	type args struct {
		data []byte
		id   StringUUID
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
			s := &service{
				db: tt.fields.db,
			}
			if err := s.UpdateUser(tt.args.data, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
