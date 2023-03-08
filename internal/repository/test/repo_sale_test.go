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

func TestSaleRepository_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockSaleRepo := mocks.NewMockSaleRepository(mockCtrl)

	SaleInput := &domain.SaleInput{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}
	expectedModel := &models.ModelSale{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockSaleRepo.EXPECT().Create(SaleInput).Return(expectedModel, nil)

	saleService := &service.ServiceSale{Repository: mockSaleRepo}
	createdModel, err := saleService.Create(SaleInput)
	assert.Nil(t, err)
	assert.Equal(t, expectedModel, createdModel)
}

func TestSaleRepository_Result(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	myUuid := uuid.NewString()

	mockSaleRepo := mocks.NewMockSaleRepository(mockCtrl)

	SaleInput := &domain.SaleInput{ID: myUuid}

	expectedModel := &models.ModelSale{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockSaleRepo.EXPECT().Result(SaleInput).Return(expectedModel, nil)

	saleService := &service.ServiceSale{Repository: mockSaleRepo}

	resultModel, err := saleService.Result(SaleInput)

	assert.Equal(t, expectedModel, resultModel)

	assert.Nil(t, err)
}

func TestSaleRepository_Results(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSaleRepository(mockCtrl)

	myUuid := uuid.NewString()

	expectedModels := []models.ModelSale{
		{ID: myUuid, Name: "ikanb", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"},
		{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"},
	}

	mockRepo.EXPECT().Results().Return(&expectedModels, nil)

	saleService := &service.ServiceSale{Repository: mockRepo}
	results, err := saleService.Results()

	assert.Nil(t, err)
	assert.Equal(t, &expectedModels, results)
}

func TestSaleRepository_Delete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSaleRepository(mockCtrl)

	myUuid := uuid.NewString()

	SaleInput := &domain.SaleInput{ID: myUuid}

	expectedModel := &models.ModelSale{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockRepo.EXPECT().Delete(SaleInput).Return(expectedModel, nil)

	saleService := service.ServiceSale{Repository: mockRepo}
	deletedModel, err := saleService.Delete(SaleInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, deletedModel)
}

func TestSaleRepository_Update(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSaleRepository(mockCtrl)

	myUuid := uuid.NewString()

	SaleInput := &domain.SaleInput{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	expectedModel := &models.ModelSale{ID: myUuid, Name: "ikan", Alamat: "ujung kulon", Email: "dota2@gmail.com", Telepon: "0888822122"}

	mockRepo.EXPECT().Update(SaleInput).Return(expectedModel, nil)

	saleService := service.ServiceSale{Repository: mockRepo}
	updatedModel, err := saleService.Update(SaleInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedModel, updatedModel)
}
