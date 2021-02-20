// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_template is a generated GoMock package.
package mock_template

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/rknruben56/feedback-api/entity"
	reflect "reflect"
)

// MockReader is a mock of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockReader) Get(id entity.ID) (*entity.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockReaderMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReader)(nil).Get), id)
}

// List mocks base method
func (m *MockReader) List() ([]*entity.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockReaderMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReader)(nil).List))
}

// MockWriter is a mock of Writer interface
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockWriter) Create(e *entity.Template) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockWriterMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriter)(nil).Create), e)
}

// Update mocks base method
func (m *MockWriter) Update(e *entity.Template) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockWriterMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWriter)(nil).Update), e)
}

// Delete mocks base method
func (m *MockWriter) Delete(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockWriterMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWriter)(nil).Delete), id)
}

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRepository) Get(id entity.ID) (*entity.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), id)
}

// List mocks base method
func (m *MockRepository) List() ([]*entity.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *entity.Template) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// Update mocks base method
func (m *MockRepository) Update(e *entity.Template) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRepositoryMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), e)
}

// Delete mocks base method
func (m *MockRepository) Delete(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// GetTemplate mocks base method
func (m *MockUseCase) GetTemplate(id entity.ID) (*entity.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplate", id)
	ret0, _ := ret[0].(*entity.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplate indicates an expected call of GetTemplate
func (mr *MockUseCaseMockRecorder) GetTemplate(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplate", reflect.TypeOf((*MockUseCase)(nil).GetTemplate), id)
}

// ListTemplates mocks base method
func (m *MockUseCase) ListTemplates() ([]*entity.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTemplates")
	ret0, _ := ret[0].([]*entity.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTemplates indicates an expected call of ListTemplates
func (mr *MockUseCaseMockRecorder) ListTemplates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTemplates", reflect.TypeOf((*MockUseCase)(nil).ListTemplates))
}

// CreateTemplate mocks base method
func (m *MockUseCase) CreateTemplate(class, content string) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTemplate", class, content)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTemplate indicates an expected call of CreateTemplate
func (mr *MockUseCaseMockRecorder) CreateTemplate(class, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemplate", reflect.TypeOf((*MockUseCase)(nil).CreateTemplate), class, content)
}

// UpdateTemplate mocks base method
func (m *MockUseCase) UpdateTemplate(e *entity.Template) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTemplate", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTemplate indicates an expected call of UpdateTemplate
func (mr *MockUseCaseMockRecorder) UpdateTemplate(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTemplate", reflect.TypeOf((*MockUseCase)(nil).UpdateTemplate), e)
}

// DeleteTemplate mocks base method
func (m *MockUseCase) DeleteTemplate(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTemplate", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTemplate indicates an expected call of DeleteTemplate
func (mr *MockUseCaseMockRecorder) DeleteTemplate(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTemplate", reflect.TypeOf((*MockUseCase)(nil).DeleteTemplate), id)
}
