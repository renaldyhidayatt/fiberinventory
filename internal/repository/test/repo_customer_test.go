package test

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	mocks "fiberinventory/internal/repository/mocks"
	"fiberinventory/internal/service"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCustomerRepository_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockCustomerRepo := mocks.NewMockCustomerRepository(mockCtrl)

	customerInput := &domain.CreateCustomerRequest{Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}
	expectedModel := &models.ModelCustomer{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockCustomerRepo.EXPECT().Create(customerInput).Return(expectedModel, nil)

	customerService := &service.ServiceCustomer{Repository: mockCustomerRepo}
	createdModel, err := customerService.Create(customerInput)
	assert.Nil(t, err)
	assert.Equal(t, expectedModel, createdModel)
}

func TestCustomerRepository_Result(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(mockCtrl)

	myUuid := uuid.NewString()

	expectedModel := &models.ModelCustomer{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockRepo.EXPECT().Result(myUuid).Return(expectedModel, nil)

	categoryService := &service.ServiceCustomer{Repository: mockRepo}

	resultModel, err := categoryService.Result(myUuid)

	assert.Equal(t, expectedModel, resultModel)

	assert.Nil(t, err)
}

func TestCustomerRepository_Results(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(mockCtrl)

	myUuid := uuid.NewString()
	expectedModels := []models.ModelCustomer{
		{ID: myUuid, Name: "ikanb", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"},
		{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"},
	}

	mockRepo.EXPECT().Results().Return(&expectedModels, nil)

	customerService := &service.ServiceCustomer{Repository: mockRepo}
	results, err := customerService.Results()

	assert.Nil(t, err)
	assert.Equal(t, &expectedModels, results)
}

func TestCustomerRepository_Delete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(mockCtrl)

	myUuid := uuid.NewString()

	expectedModel := &models.ModelCustomer{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockRepo.EXPECT().Delete(myUuid).Return(expectedModel, nil)

	customerService := service.ServiceCustomer{Repository: mockRepo}
	deletedModel, err := customerService.Delete(myUuid)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, deletedModel)
}

func TestCustomerRepository_Update(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(mockCtrl)

	myUuid := uuid.NewString()

	customerInput := &domain.UpdateCustomerRequest{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	expectedModel := &models.ModelCustomer{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockRepo.EXPECT().Update(customerInput).Return(expectedModel, nil)

	customerService := service.ServiceCustomer{Repository: mockRepo}
	updatedModel, err := customerService.Update(customerInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, updatedModel)
}
