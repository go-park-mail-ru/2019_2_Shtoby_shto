package user

import (
	. "2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"reflect"
	"regexp"
	"testing"
)

type Suite struct {
	suite.Suite
	DB      *gorm.DB
	mock    sqlmock.Sqlmock
	handler HandlerUserService
	service *service
	model   *User
}

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

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, mock, err := sqlmock.New()
	if err != nil {
		require.NoError(s.T(), err)
	}
	s.mock = mock

	s.DB, err = gorm.Open("postgres", db)
	if err != nil {
		require.NoError(s.T(), err)
	}

	s.DB.LogMode(true)

	dm := database.NewDataManager(s.DB)

	s.handler = CreateInstance(dm)
}

func Test_service_CreateUser(t *testing.T) {
	s := Suite{}
	s.SetupSuite()
	id := uuid.NewV4()
	user := &User{
		BaseInfo: dicts.BaseInfo{
			ID: StringUUID(id.String()),
		},
		Login:    "sanya23232",
		Password: "1111",
	}
	userData, err := user.MarshalJSON()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("id","name") 
       VALUES ($1,$2) RETURNING "users"."id"`)).
		WithArgs(id, user.Login).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(id.String()))

	createdUser, err := s.handler.CreateUser(userData)
	if err != nil {
		require.NoError(s.T(), err)
	}
	require.Equal(t, createdUser.Password, user.Password)

	//type fields struct {
	//	db database.IDataManager
	//}
	//type args struct {
	//	data []byte
	//}
	//tests := []struct {
	//	name    string
	//	fields  fields
	//	args    args
	//	want    *User
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		s := &service{
	//			db: tt.fields.db,
	//		}
	//		got, err := s.CreateUser(tt.args.data)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("service.CreateUser() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func Test_service_GetUserById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	get := NewMockHandlerUserService(mockCtrl)
	get.EXPECT().GetUserById("b4f7a3e5-5d29-4f28-8aa8-05d9cec60bd4").Return(nil)

	s := Suite{}
	s.SetupSuite()
	id := "b4f7a3e5-5d29-4f28-8aa8-05d9cec60bd4"
	tests := []struct {
		name    string
		id      StringUUID
		want    User
		wantErr bool
	}{
		{
			name: "1",
			id:   StringUUID(id),
			want: User{
				BaseInfo: dicts.BaseInfo{
					ID: StringUUID(id),
				},
				Password: "",
				Login:    "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.handler.GetUserById(tt.id)
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
