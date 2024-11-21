package repo

import (
	context "context"
	reflect "reflect"

	models "github.com/cooldev900/go-rest-concurrency/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRepo) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRepoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepo)(nil).Close))
}

// CreateOrder mocks base method.
func (m *MockRepo) CreateOrder(item models.Item) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", item)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockRepoMockRecorder) CreateOrder(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockRepo)(nil).CreateOrder), item)
}

// GetAllProducts mocks base method.
func (m *MockRepo) GetAllProducts() []models.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts")
	ret0, _ := ret[0].([]models.Product)
	return ret0
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockRepoMockRecorder) GetAllProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockRepo)(nil).GetAllProducts))
}

// GetOrder mocks base method.
func (m *MockRepo) GetOrder(id string) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", id)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockRepoMockRecorder) GetOrder(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockRepo)(nil).GetOrder), id)
}

// GetOrderStats mocks base method.
func (m *MockRepo) GetOrderStats(ctx context.Context) (models.Statistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderStats", ctx)
	ret0, _ := ret[0].(models.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderStats indicates an expected call of GetOrderStats.
func (mr *MockRepoMockRecorder) GetOrderStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderStats", reflect.TypeOf((*MockRepo)(nil).GetOrderStats), ctx)
}

// RequestReversal mocks base method.
func (m *MockRepo) RequestReversal(orderId string) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestReversal", orderId)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestReversal indicates an expected call of RequestReversal.
func (mr *MockRepoMockRecorder) RequestReversal(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestReversal", reflect.TypeOf((*MockRepo)(nil).RequestReversal), orderId)
}
