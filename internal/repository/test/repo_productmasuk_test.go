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

func TestProductMasukRepository_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockRepo := mocks.NewMockProductMasukRepository(mockCtrl)
	ProductMasukInput := &domain.CreateProductMasukRequest{Name: "TestProductMasuk", Qty: "5", ProductID: myUuid, SupplierID: myUuid}

	expectedModel := &models.ModelProductMasuk{ID: myUuid, Name: "TestProductMasuk", Qty: "5", ProductID: myUuid, SupplierID: myUuid}

	mockRepo.EXPECT().Create(ProductMasukInput).Return(expectedModel, nil)

	productMasukService := &service.ServiceProductMasuk{Repository: mockRepo}
	createdModel, err := productMasukService.Create(ProductMasukInput)
	assert.Nil(t, err)
	assert.Equal(t, expectedModel, createdModel)
}

func TestProductMasukRepository_Result(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockProductMasukRepository(mockCtrl)
	myUuid := uuid.NewString()

	expectedModel := &models.ModelProductMasuk{
		ID: myUuid, Name: "TestProductMasuk", Qty: "5", ProductID: myUuid, SupplierID: myUuid,
	}

	mockRepo.EXPECT().Result(myUuid).Return(expectedModel, nil)

	productMasukService := &service.ServiceProductMasuk{Repository: mockRepo}

	resultModel, err := productMasukService.Result(myUuid)

	assert.Equal(t, expectedModel, resultModel)

	assert.Nil(t, err)
}

func TestProductMasukRepository_Results(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockProductMasukRepository(mockCtrl)
	myUuid := uuid.NewString()

	expectedModels := []models.ModelProductMasuk{
		{ID: myUuid, Name: "TestProductMasuk", Qty: "5", ProductID: myUuid, SupplierID: myUuid},
		{ID: myUuid, Name: "TestProductMasuk", Qty: "5", ProductID: myUuid, SupplierID: myUuid},
	}

	mockRepo.EXPECT().Results().Return(&expectedModels, nil)

	productMasukService := &service.ServiceProductMasuk{Repository: mockRepo}

	result, err := productMasukService.Results()

	assert.Nil(t, err)
	assert.Equal(t, &expectedModels, result)

}

func TestProductMasukRepository_Delete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockProductMasukRepository(mockCtrl)
	myUuid := uuid.NewString()

	expectedModel := &models.ModelProductMasuk{ID: myUuid, Name: "TestProductMasuk", Qty: "5", ProductID: myUuid, SupplierID: myUuid}

	mockRepo.EXPECT().Delete(myUuid).Return(expectedModel, nil)

	productMasukService := service.ServiceProductMasuk{Repository: mockRepo}
	deletedModel, err := productMasukService.Delete(myUuid)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, deletedModel)
}

func TestProductMasukRepository_Update(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockProductMasukRepository(mockCtrl)
	myUuid := uuid.NewString()

	ProductMasukInput := &domain.UpdateProductMasukRequest{ID: myUuid, Name: "TestProductMasuk", Qty: "5", ProductID: myUuid, SupplierID: myUuid}

	expectedModel := &models.ModelProductMasuk{ID: myUuid, Name: "TestProductMasuk", Qty: "5", ProductID: myUuid, SupplierID: myUuid}

	mockRepo.EXPECT().Update(ProductMasukInput).Return(expectedModel, nil)

	productMasukService := service.ServiceProductMasuk{Repository: mockRepo}
	updatedModel, err := productMasukService.Update(ProductMasukInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, updatedModel)
}
