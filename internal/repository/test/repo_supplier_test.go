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

func TestSupplierRepository_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockRepo := mocks.NewMockSupplierRepository(mockCtrl)

	SupplierInput := &domain.CreateSupplierRequest{Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}
	expectedModel := &models.ModelSupplier{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockRepo.EXPECT().Create(SupplierInput).Return(expectedModel, nil)

	supplierService := &service.ServiceSupplier{Repository: mockRepo}
	createdModel, err := supplierService.Create(SupplierInput)
	assert.Nil(t, err)
	assert.Equal(t, expectedModel, createdModel)
}

func TestSupplierRepository_Result(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockRepo := mocks.NewMockSupplierRepository(mockCtrl)

	expectedModel := &models.ModelSupplier{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockRepo.EXPECT().Result(myUuid).Return(expectedModel, nil)

	supplierService := &service.ServiceSupplier{Repository: mockRepo}

	resultModel, err := supplierService.Result(myUuid)

	assert.Equal(t, expectedModel, resultModel)

	assert.Nil(t, err)
}

func TestSupplierRepository_Results(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSupplierRepository(mockCtrl)

	myUuid := uuid.NewString()

	expectedModels := []models.ModelSupplier{
		{ID: myUuid, Name: "ikanb", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"},
		{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"},
	}

	mockRepo.EXPECT().Results().Return(&expectedModels, nil)

	supplierService := &service.ServiceSupplier{Repository: mockRepo}
	results, err := supplierService.Results()

	assert.Nil(t, err)
	assert.Equal(t, &expectedModels, results)
}

func TestSupplierRepository_Delete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSupplierRepository(mockCtrl)

	myUuid := uuid.NewString()

	expectedModel := &models.ModelSupplier{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockRepo.EXPECT().Delete(myUuid).Return(expectedModel, nil)

	supplierService := service.ServiceSupplier{Repository: mockRepo}
	deletedModel, err := supplierService.Delete(myUuid)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, deletedModel)
}

func TestSupplierRepository_Update(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSupplierRepository(mockCtrl)

	myUuid := uuid.NewString()

	SupplierInput := &domain.UpdateSupplierRequest{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	expectedModel := &models.ModelSupplier{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockRepo.EXPECT().Update(SupplierInput).Return(expectedModel, nil)

	supplierService := service.ServiceSupplier{Repository: mockRepo}
	updatedModel, err := supplierService.Update(SupplierInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, updatedModel)
}
