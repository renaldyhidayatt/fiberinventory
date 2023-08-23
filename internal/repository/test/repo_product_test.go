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

func TestProductRepository_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	productInput := &domain.CreateProductRequest{Name: "TestCategory", Image: "ikan.png", Qty: "5"}

	expectedModel := &models.ModelProduct{Name: "TestCategory", Image: "ikan.png", Qty: "5", CategoryID: myUuid, ID: myUuid}

	mockProductRepository.EXPECT().Create(productInput).Return(expectedModel, nil)

	productService := &service.ServiceProduct{Repository: mockProductRepository}
	createdModel, err := productService.Create(productInput)
	assert.Nil(t, err)
	assert.Equal(t, expectedModel, createdModel)
}

func TestProductRepository_Result(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockRepo := mocks.NewMockProductRepository(mockCtrl)

	expectedModel := &models.ModelProduct{
		ID:   myUuid,
		Name: "TestCategory", Image: "ikan.png", Qty: "5", CategoryID: myUuid,
	}

	mockRepo.EXPECT().Result(myUuid).Return(expectedModel, nil)

	productService := &service.ServiceProduct{Repository: mockRepo}

	resultModel, err := productService.Result(myUuid)

	assert.Equal(t, expectedModel, resultModel)

	assert.Nil(t, err)
}

func TestProductRepository_Results(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	myUuid := uuid.NewString()

	mockRepo := mocks.NewMockProductRepository(ctrl)

	expectedModels := []models.ModelProduct{
		{ID: myUuid, Name: "TestCategory", Image: "ikan.png", Qty: "5", CategoryID: myUuid},
		{ID: myUuid, Name: "TestCategory", Image: "ikan.png", Qty: "5", CategoryID: myUuid},
	}

	mockRepo.EXPECT().Results().Return(&expectedModels, nil)

	productService := &service.ServiceProduct{Repository: mockRepo}

	result, err := productService.Results()

	assert.Nil(t, err)
	assert.Equal(t, &expectedModels, result)

}

func TestProductRepository_Delete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockRepo := mocks.NewMockProductRepository(mockCtrl)

	expectedModel := &models.ModelProduct{ID: myUuid, Name: "TestCategory", Image: "ikan.png", Qty: "5", CategoryID: myUuid}

	mockRepo.EXPECT().Delete(myUuid).Return(expectedModel, nil)

	productService := service.ServiceProduct{Repository: mockRepo}
	deletedModel, err := productService.Delete(myUuid)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, deletedModel)
}

func TestProductRepository_Update(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockProductRepository(mockCtrl)

	myUuid := uuid.NewString()

	productInput := &domain.UpdateProductRequest{Name: "TestCategory", Image: "ikan.png", Qty: "5", ID: myUuid}

	expectedModel := &models.ModelProduct{Name: "TestCategory", Image: "ikan.png", Qty: "5", CategoryID: myUuid, ID: myUuid}

	mockRepo.EXPECT().Update(productInput).Return(expectedModel, nil)

	productService := service.ServiceProduct{Repository: mockRepo}
	updatedModel, err := productService.Update(productInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, updatedModel)
}
