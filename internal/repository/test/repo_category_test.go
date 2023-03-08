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

func TestCategoryRepository_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)
	categoryInput := &domain.CategoryInput{Name: "TestCategory", ID: myUuid}

	expectedModel := &models.ModelCategory{Name: "TestCategory", ID: myUuid}

	mockCategoryRepo.EXPECT().Create(categoryInput).Return(expectedModel, nil)

	categoryService := &service.ServiceCategory{Repository: mockCategoryRepo}
	createdModel, err := categoryService.Create(categoryInput)
	assert.Nil(t, err)
	assert.Equal(t, expectedModel, createdModel)
}

func TestCategoryRepository_Result(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCategoryRepository(mockCtrl)

	input := &domain.CategoryInput{ID: "category_id"}

	expectedModel := &models.ModelCategory{
		ID:   "category_id",
		Name: "TestCategory",
	}

	mockRepo.EXPECT().Result(input).Return(expectedModel, nil)

	categoryService := &service.ServiceCategory{Repository: mockRepo}

	resultModel, err := categoryService.Result(input)

	assert.Equal(t, expectedModel, resultModel)

	assert.Nil(t, err)
}

func TestCategoryRepository_Results(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepo := mocks.NewMockCategoryRepository(ctrl)

	expectedModels := []models.ModelCategory{
		{ID: "category-1", Name: "TestCategory1"},
		{ID: "category-2", Name: "TestCategory2"},
	}

	mockCategoryRepo.EXPECT().Results().Return(&expectedModels, nil)

	categoryService := &service.ServiceCategory{Repository: mockCategoryRepo}

	result, err := categoryService.Results()

	assert.Nil(t, err)
	assert.Equal(t, &expectedModels, result)

}

func TestCategoryRepository_Delete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)

	categoryID := uuid.NewString()
	input := &domain.CategoryInput{ID: categoryID}
	expectedModel := &models.ModelCategory{ID: categoryID, Name: "TestCategory"}

	mockCategoryRepo.EXPECT().Delete(input).Return(expectedModel, nil)

	categoryService := service.ServiceCategory{Repository: mockCategoryRepo}
	deletedModel, err := categoryService.Delete(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, deletedModel)
}

func TestCategoryRepository_Update(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)

	categoryId := uuid.NewString()

	categoryInput := &domain.CategoryInput{Name: "TestCategory", ID: categoryId}
	expectedModel := &models.ModelCategory{ID: categoryId, Name: "TestCategory"}

	mockCategoryRepo.EXPECT().Update(categoryInput).Return(expectedModel, nil)

	categoryService := service.ServiceCategory{Repository: mockCategoryRepo}
	updatedModel, err := categoryService.Update(categoryInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, updatedModel)
}
