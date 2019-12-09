// Code generated by MockGen. DO NOT EDIT.
// Source: data_manager.go

// Package database is a generated GoMock package.
package database

import (
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
	reflect "reflect"
)

// MockIDataManager is a mock of IDataManager interface
type MockIDataManager struct {
	ctrl     *gomock.Controller
	recorder *MockIDataManagerMockRecorder
}

// MockIDataManagerMockRecorder is the mock recorder for MockIDataManager
type MockIDataManagerMockRecorder struct {
	mock *MockIDataManager
}

// NewMockIDataManager creates a new mock instance
func NewMockIDataManager(ctrl *gomock.Controller) *MockIDataManager {
	mock := &MockIDataManager{ctrl: ctrl}
	mock.recorder = &MockIDataManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDataManager) EXPECT() *MockIDataManagerMockRecorder {
	return m.recorder
}

// Db mocks base method
func (m *MockIDataManager) Db() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Db")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Db indicates an expected call of Db
func (mr *MockIDataManagerMockRecorder) Db() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Db", reflect.TypeOf((*MockIDataManager)(nil).Db))
}

// SetDb mocks base method
func (m *MockIDataManager) SetDb(db *gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDb", db)
}

// SetDb indicates an expected call of SetDb
func (mr *MockIDataManagerMockRecorder) SetDb(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDb", reflect.TypeOf((*MockIDataManager)(nil).SetDb), db)
}

// CloseConnection mocks base method
func (m *MockIDataManager) CloseConnection() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseConnection")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseConnection indicates an expected call of CloseConnection
func (mr *MockIDataManagerMockRecorder) CloseConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseConnection", reflect.TypeOf((*MockIDataManager)(nil).CloseConnection))
}

// ExecuteQuery mocks base method
func (m *MockIDataManager) ExecuteQuery(sql string, args ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecuteQuery", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteQuery indicates an expected call of ExecuteQuery
func (mr *MockIDataManagerMockRecorder) ExecuteQuery(sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteQuery", reflect.TypeOf((*MockIDataManager)(nil).ExecuteQuery), varargs...)
}

// FindDictById mocks base method
func (m *MockIDataManager) FindDictById(p interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDictById", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindDictById indicates an expected call of FindDictById
func (mr *MockIDataManagerMockRecorder) FindDictById(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDictById", reflect.TypeOf((*MockIDataManager)(nil).FindDictById), p)
}

// FindDictByColumn mocks base method
func (m *MockIDataManager) FindDictByColumn(p interface{}) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDictByColumn", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDictByColumn indicates an expected call of FindDictByColumn
func (mr *MockIDataManagerMockRecorder) FindDictByColumn(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDictByColumn", reflect.TypeOf((*MockIDataManager)(nil).FindDictByColumn), p)
}

// CreateRecord mocks base method
func (m *MockIDataManager) CreateRecord(p interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRecord indicates an expected call of CreateRecord
func (mr *MockIDataManagerMockRecorder) CreateRecord(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockIDataManager)(nil).CreateRecord), p)
}

// UpdateRecord mocks base method
func (m *MockIDataManager) UpdateRecord(p interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecord", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRecord indicates an expected call of UpdateRecord
func (mr *MockIDataManagerMockRecorder) UpdateRecord(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecord", reflect.TypeOf((*MockIDataManager)(nil).UpdateRecord), p)
}

// DeleteRecord mocks base method
func (m *MockIDataManager) DeleteRecord(p interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecord", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecord indicates an expected call of DeleteRecord
func (mr *MockIDataManagerMockRecorder) DeleteRecord(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecord", reflect.TypeOf((*MockIDataManager)(nil).DeleteRecord), p)
}

// FetchDict mocks base method
func (m *MockIDataManager) FetchDict(data, findRequest interface{}, limit, offset int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDict", data, findRequest, limit, offset)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchDict indicates an expected call of FetchDict
func (mr *MockIDataManagerMockRecorder) FetchDict(data, findRequest, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDict", reflect.TypeOf((*MockIDataManager)(nil).FetchDict), data, findRequest, limit, offset)
}

// FetchDictBySlice mocks base method
func (m *MockIDataManager) FetchDictBySlice(data interface{}, table string, limit, offset int, where, whereArg []string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDictBySlice", data, table, limit, offset, where, whereArg)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchDictBySlice indicates an expected call of FetchDictBySlice
func (mr *MockIDataManagerMockRecorder) FetchDictBySlice(data, table, limit, offset, where, whereArg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDictBySlice", reflect.TypeOf((*MockIDataManager)(nil).FetchDictBySlice), data, table, limit, offset, where, whereArg)
}