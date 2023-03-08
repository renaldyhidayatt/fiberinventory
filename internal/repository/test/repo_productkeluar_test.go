package test

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	mocks "fiberinventory/internal/repository/mocks"
	"fiberinventory/internal/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProductKeluarRepository_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockRepo := mocks.NewMockProductKeluarRepository(mockCtrl)
	productKeluarInput := &domain.ProductKeluarInput{ID: myUuid, Qty: "5", ProductID: myUuid, CategoryID: myUuid}

	expectedModel := &models.ModelProductKeluar{Qty: "5", ProductID: myUuid, CategoryID: myUuid}

	mockRepo.EXPECT().Create(productKeluarInput).Return(expectedModel, nil)

	productService := &service.ServiceProductKeluar{Repository: mockRepo}
	createdModel, err := productService.Create(productKeluarInput)
	assert.Nil(t, err)
	assert.Equal(t, expectedModel, createdModel)
}

func TestProductKeluarRepository_Result(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()
	mockRepo := mocks.NewMockProductKeluarRepository(mockCtrl)

	input := &domain.ProductKeluarInput{ID: myUuid}

	expectedModel := &models.ModelProductKeluar{
		ID: myUuid, Qty: "5", ProductID: myUuid, CategoryID: myUuid,
	}

	mockRepo.EXPECT().Result(input).Return(expectedModel, nil)

	productService := &service.ServiceProductKeluar{Repository: mockRepo}

	resultModel, err := productService.Result(input)

	assert.Equal(t, expectedModel, resultModel)

	assert.Nil(t, err)
}

func TestProductKeluarRepository_Results(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockProductKeluarRepository(mockCtrl)

	myUuid := uuid.NewString()

	expectedModels := []models.ModelProductKeluar{
		{ID: myUuid, Qty: "5", ProductID: myUuid, CategoryID: myUuid},
		{ID: myUuid, Qty: "5", ProductID: myUuid, CategoryID: myUuid},
	}

	mockRepo.EXPECT().Results().Return(&expectedModels, nil)

	productService := &service.ServiceProductKeluar{Repository: mockRepo}

	result, err := productService.Results()

	assert.Nil(t, err)
	assert.Equal(t, &expectedModels, result)

}

func TestProductKeluarRepository_Delete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockProductKeluarRepository(mockCtrl)

	myUuid := uuid.NewString()

	input := &domain.ProductKeluarInput{ID: myUuid}
	expectedModel := &models.ModelProductKeluar{ID: myUuid, Qty: "5", ProductID: myUuid, CategoryID: myUuid}

	mockRepo.EXPECT().Delete(input).Return(expectedModel, nil)

	productService := service.ServiceProductKeluar{Repository: mockRepo}
	deletedModel, err := productService.Delete(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, deletedModel)
}

func TestProductKeluarRepository_Update(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockProductKeluarRepository(mockCtrl)

	myUuid := uuid.NewString()

	productKeluarInput := &domain.ProductKeluarInput{ID: myUuid, Qty: "5", ProductID: myUuid, CategoryID: myUuid}

	expectedModel := &models.ModelProductKeluar{Qty: "5", ProductID: myUuid, CategoryID: myUuid}

	mockRepo.EXPECT().Update(productKeluarInput).Return(expectedModel, nil)

	productService := service.ServiceProductKeluar{Repository: mockRepo}
	updatedModel, err := productService.Update(productKeluarInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, updatedModel)
}
