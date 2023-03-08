// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/interfaces.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	domain "fiberinventory/internal/domain"
	models "fiberinventory/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserRepository) Delete(input *domain.UserInput) (*models.ModelUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", input)
	ret0, _ := ret[0].(*models.ModelUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepositoryMockRecorder) Delete(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepository)(nil).Delete), input)
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(email string) (*models.ModelUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*models.ModelUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), email)
}

// Login mocks base method.
func (m *MockUserRepository) Login(input *domain.UserInput) (*models.ModelUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", input)
	ret0, _ := ret[0].(*models.ModelUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserRepositoryMockRecorder) Login(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserRepository)(nil).Login), input)
}

// Register mocks base method.
func (m *MockUserRepository) Register(input *domain.UserInput) (*models.ModelUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", input)
	ret0, _ := ret[0].(*models.ModelUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserRepositoryMockRecorder) Register(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserRepository)(nil).Register), input)
}

// Result mocks base method.
func (m *MockUserRepository) Result(input *domain.UserInput) (*models.ModelUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result", input)
	ret0, _ := ret[0].(*models.ModelUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Result indicates an expected call of Result.
func (mr *MockUserRepositoryMockRecorder) Result(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockUserRepository)(nil).Result), input)
}

// Results mocks base method.
func (m *MockUserRepository) Results() (*[]models.ModelUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Results")
	ret0, _ := ret[0].(*[]models.ModelUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Results indicates an expected call of Results.
func (mr *MockUserRepositoryMockRecorder) Results() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Results", reflect.TypeOf((*MockUserRepository)(nil).Results))
}

// Update mocks base method.
func (m *MockUserRepository) Update(input *domain.UserInput) (*models.ModelUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(*models.ModelUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), input)
}

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCategoryRepository) Create(input *domain.CategoryInput) (*models.ModelCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*models.ModelCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCategoryRepositoryMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryRepository)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockCategoryRepository) Delete(input *domain.CategoryInput) (*models.ModelCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", input)
	ret0, _ := ret[0].(*models.ModelCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryRepositoryMockRecorder) Delete(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategoryRepository)(nil).Delete), input)
}

// Result mocks base method.
func (m *MockCategoryRepository) Result(input *domain.CategoryInput) (*models.ModelCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result", input)
	ret0, _ := ret[0].(*models.ModelCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Result indicates an expected call of Result.
func (mr *MockCategoryRepositoryMockRecorder) Result(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockCategoryRepository)(nil).Result), input)
}

// Results mocks base method.
func (m *MockCategoryRepository) Results() (*[]models.ModelCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Results")
	ret0, _ := ret[0].(*[]models.ModelCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Results indicates an expected call of Results.
func (mr *MockCategoryRepositoryMockRecorder) Results() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Results", reflect.TypeOf((*MockCategoryRepository)(nil).Results))
}

// Update mocks base method.
func (m *MockCategoryRepository) Update(input *domain.CategoryInput) (*models.ModelCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(*models.ModelCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCategoryRepositoryMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategoryRepository)(nil).Update), input)
}

// MockCustomerRepository is a mock of CustomerRepository interface.
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository.
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance.
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCustomerRepository) Create(input *domain.CustomerInput) (*models.ModelCustomer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*models.ModelCustomer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCustomerRepositoryMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomerRepository)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockCustomerRepository) Delete(input *domain.CustomerInput) (*models.ModelCustomer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", input)
	ret0, _ := ret[0].(*models.ModelCustomer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomerRepositoryMockRecorder) Delete(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomerRepository)(nil).Delete), input)
}

// Result mocks base method.
func (m *MockCustomerRepository) Result(input *domain.CustomerInput) (*models.ModelCustomer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result", input)
	ret0, _ := ret[0].(*models.ModelCustomer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Result indicates an expected call of Result.
func (mr *MockCustomerRepositoryMockRecorder) Result(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockCustomerRepository)(nil).Result), input)
}

// Results mocks base method.
func (m *MockCustomerRepository) Results() (*[]models.ModelCustomer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Results")
	ret0, _ := ret[0].(*[]models.ModelCustomer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Results indicates an expected call of Results.
func (mr *MockCustomerRepositoryMockRecorder) Results() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Results", reflect.TypeOf((*MockCustomerRepository)(nil).Results))
}

// Update mocks base method.
func (m *MockCustomerRepository) Update(input *domain.CustomerInput) (*models.ModelCustomer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(*models.ModelCustomer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCustomerRepositoryMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerRepository)(nil).Update), input)
}

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductRepository) Create(input *domain.ProductInput) (*models.ModelProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*models.ModelProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductRepositoryMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductRepository)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockProductRepository) Delete(input *domain.ProductInput) (*models.ModelProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", input)
	ret0, _ := ret[0].(*models.ModelProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockProductRepositoryMockRecorder) Delete(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductRepository)(nil).Delete), input)
}

// Result mocks base method.
func (m *MockProductRepository) Result(input *domain.ProductInput) (*models.ModelProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result", input)
	ret0, _ := ret[0].(*models.ModelProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Result indicates an expected call of Result.
func (mr *MockProductRepositoryMockRecorder) Result(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockProductRepository)(nil).Result), input)
}

// Results mocks base method.
func (m *MockProductRepository) Results() (*[]models.ModelProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Results")
	ret0, _ := ret[0].(*[]models.ModelProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Results indicates an expected call of Results.
func (mr *MockProductRepositoryMockRecorder) Results() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Results", reflect.TypeOf((*MockProductRepository)(nil).Results))
}

// Update mocks base method.
func (m *MockProductRepository) Update(input *domain.ProductInput) (*models.ModelProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(*models.ModelProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductRepositoryMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductRepository)(nil).Update), input)
}

// MockProductKeluarRepository is a mock of ProductKeluarRepository interface.
type MockProductKeluarRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductKeluarRepositoryMockRecorder
}

// MockProductKeluarRepositoryMockRecorder is the mock recorder for MockProductKeluarRepository.
type MockProductKeluarRepositoryMockRecorder struct {
	mock *MockProductKeluarRepository
}

// NewMockProductKeluarRepository creates a new mock instance.
func NewMockProductKeluarRepository(ctrl *gomock.Controller) *MockProductKeluarRepository {
	mock := &MockProductKeluarRepository{ctrl: ctrl}
	mock.recorder = &MockProductKeluarRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductKeluarRepository) EXPECT() *MockProductKeluarRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductKeluarRepository) Create(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*models.ModelProductKeluar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductKeluarRepositoryMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductKeluarRepository)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockProductKeluarRepository) Delete(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", input)
	ret0, _ := ret[0].(*models.ModelProductKeluar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockProductKeluarRepositoryMockRecorder) Delete(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductKeluarRepository)(nil).Delete), input)
}

// Result mocks base method.
func (m *MockProductKeluarRepository) Result(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result", input)
	ret0, _ := ret[0].(*models.ModelProductKeluar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Result indicates an expected call of Result.
func (mr *MockProductKeluarRepositoryMockRecorder) Result(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockProductKeluarRepository)(nil).Result), input)
}

// Results mocks base method.
func (m *MockProductKeluarRepository) Results() (*[]models.ModelProductKeluar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Results")
	ret0, _ := ret[0].(*[]models.ModelProductKeluar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Results indicates an expected call of Results.
func (mr *MockProductKeluarRepositoryMockRecorder) Results() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Results", reflect.TypeOf((*MockProductKeluarRepository)(nil).Results))
}

// Update mocks base method.
func (m *MockProductKeluarRepository) Update(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(*models.ModelProductKeluar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductKeluarRepositoryMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductKeluarRepository)(nil).Update), input)
}

// MockProductMasukRepository is a mock of ProductMasukRepository interface.
type MockProductMasukRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductMasukRepositoryMockRecorder
}

// MockProductMasukRepositoryMockRecorder is the mock recorder for MockProductMasukRepository.
type MockProductMasukRepositoryMockRecorder struct {
	mock *MockProductMasukRepository
}

// NewMockProductMasukRepository creates a new mock instance.
func NewMockProductMasukRepository(ctrl *gomock.Controller) *MockProductMasukRepository {
	mock := &MockProductMasukRepository{ctrl: ctrl}
	mock.recorder = &MockProductMasukRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductMasukRepository) EXPECT() *MockProductMasukRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductMasukRepository) Create(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*models.ModelProductMasuk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductMasukRepositoryMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductMasukRepository)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockProductMasukRepository) Delete(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", input)
	ret0, _ := ret[0].(*models.ModelProductMasuk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockProductMasukRepositoryMockRecorder) Delete(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductMasukRepository)(nil).Delete), input)
}

// Result mocks base method.
func (m *MockProductMasukRepository) Result(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result", input)
	ret0, _ := ret[0].(*models.ModelProductMasuk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Result indicates an expected call of Result.
func (mr *MockProductMasukRepositoryMockRecorder) Result(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockProductMasukRepository)(nil).Result), input)
}

// Results mocks base method.
func (m *MockProductMasukRepository) Results() (*[]models.ModelProductMasuk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Results")
	ret0, _ := ret[0].(*[]models.ModelProductMasuk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Results indicates an expected call of Results.
func (mr *MockProductMasukRepositoryMockRecorder) Results() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Results", reflect.TypeOf((*MockProductMasukRepository)(nil).Results))
}

// Update mocks base method.
func (m *MockProductMasukRepository) Update(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(*models.ModelProductMasuk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductMasukRepositoryMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductMasukRepository)(nil).Update), input)
}

// MockSaleRepository is a mock of SaleRepository interface.
type MockSaleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSaleRepositoryMockRecorder
}

// MockSaleRepositoryMockRecorder is the mock recorder for MockSaleRepository.
type MockSaleRepositoryMockRecorder struct {
	mock *MockSaleRepository
}

// NewMockSaleRepository creates a new mock instance.
func NewMockSaleRepository(ctrl *gomock.Controller) *MockSaleRepository {
	mock := &MockSaleRepository{ctrl: ctrl}
	mock.recorder = &MockSaleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaleRepository) EXPECT() *MockSaleRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSaleRepository) Create(input *domain.SaleInput) (*models.ModelSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*models.ModelSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSaleRepositoryMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSaleRepository)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockSaleRepository) Delete(input *domain.SaleInput) (*models.ModelSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", input)
	ret0, _ := ret[0].(*models.ModelSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSaleRepositoryMockRecorder) Delete(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSaleRepository)(nil).Delete), input)
}

// Result mocks base method.
func (m *MockSaleRepository) Result(input *domain.SaleInput) (*models.ModelSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result", input)
	ret0, _ := ret[0].(*models.ModelSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Result indicates an expected call of Result.
func (mr *MockSaleRepositoryMockRecorder) Result(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockSaleRepository)(nil).Result), input)
}

// Results mocks base method.
func (m *MockSaleRepository) Results() (*[]models.ModelSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Results")
	ret0, _ := ret[0].(*[]models.ModelSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Results indicates an expected call of Results.
func (mr *MockSaleRepositoryMockRecorder) Results() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Results", reflect.TypeOf((*MockSaleRepository)(nil).Results))
}

// Update mocks base method.
func (m *MockSaleRepository) Update(input *domain.SaleInput) (*models.ModelSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(*models.ModelSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSaleRepositoryMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSaleRepository)(nil).Update), input)
}

// MockSupplierRepository is a mock of SupplierRepository interface.
type MockSupplierRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSupplierRepositoryMockRecorder
}

// MockSupplierRepositoryMockRecorder is the mock recorder for MockSupplierRepository.
type MockSupplierRepositoryMockRecorder struct {
	mock *MockSupplierRepository
}

// NewMockSupplierRepository creates a new mock instance.
func NewMockSupplierRepository(ctrl *gomock.Controller) *MockSupplierRepository {
	mock := &MockSupplierRepository{ctrl: ctrl}
	mock.recorder = &MockSupplierRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSupplierRepository) EXPECT() *MockSupplierRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSupplierRepository) Create(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*models.ModelSupplier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSupplierRepositoryMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSupplierRepository)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockSupplierRepository) Delete(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", input)
	ret0, _ := ret[0].(*models.ModelSupplier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSupplierRepositoryMockRecorder) Delete(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSupplierRepository)(nil).Delete), input)
}

// Result mocks base method.
func (m *MockSupplierRepository) Result(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result", input)
	ret0, _ := ret[0].(*models.ModelSupplier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Result indicates an expected call of Result.
func (mr *MockSupplierRepositoryMockRecorder) Result(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockSupplierRepository)(nil).Result), input)
}

// Results mocks base method.
func (m *MockSupplierRepository) Results() (*[]models.ModelSupplier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Results")
	ret0, _ := ret[0].(*[]models.ModelSupplier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Results indicates an expected call of Results.
func (mr *MockSupplierRepositoryMockRecorder) Results() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Results", reflect.TypeOf((*MockSupplierRepository)(nil).Results))
}

// Update mocks base method.
func (m *MockSupplierRepository) Update(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(*models.ModelSupplier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSupplierRepositoryMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSupplierRepository)(nil).Update), input)
}
