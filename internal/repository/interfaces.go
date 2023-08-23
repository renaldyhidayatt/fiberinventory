package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type UserRepository interface {
	Register(input *domain.RegisterInput) (*models.ModelUser, error)
	Login(input *domain.LoginInput) (*models.ModelUser, error)
	Results() (*[]models.ModelUser, error)
	Result(id string) (*models.ModelUser, error)
	Delete(id string) (*models.ModelUser, error)
	Update(input *domain.UpdateUserRequest) (*models.ModelUser, error)
	FindByEmail(email string) (*models.ModelUser, error)
}

type CategoryRepository interface {
	Create(input *domain.CreateCategoryRequest) (*models.ModelCategory, error)
	Results() (*[]models.ModelCategory, error)
	Result(id string) (*models.ModelCategory, error)
	Delete(id string) (*models.ModelCategory, error)
	Update(input *domain.UpdateCategoryRequest) (*models.ModelCategory, error)
}

type CustomerRepository interface {
	Create(input *domain.CreateCustomerRequest) (*models.ModelCustomer, error)
	Results() (*[]models.ModelCustomer, error)
	Result(id string) (*models.ModelCustomer, error)
	Delete(od string) (*models.ModelCustomer, error)
	Update(input *domain.UpdateCustomerRequest) (*models.ModelCustomer, error)
}

type ProductRepository interface {
	Create(input *domain.CreateProductRequest) (*models.ModelProduct, error)
	Delete(id string) (*models.ModelProduct, error)
	Result(id string) (*models.ModelProduct, error)
	Update(input *domain.UpdateProductRequest) (*models.ModelProduct, error)
	Results() (*[]models.ModelProduct, error)
}

type ProductKeluarRepository interface {
	Create(input *domain.CreateProductKeluarRequest) (*models.ModelProductKeluar, error)
	Result(id string) (*models.ModelProductKeluar, error)
	Results() (*[]models.ModelProductKeluar, error)
	Delete(id string) (*models.ModelProductKeluar, error)
	Update(input *domain.UpdateProductKeluarRequest) (*models.ModelProductKeluar, error)
}

type ProductMasukRepository interface {
	Create(input *domain.CreateProductMasukRequest) (*models.ModelProductMasuk, error)
	Result(id string) (*models.ModelProductMasuk, error)
	Results() (*[]models.ModelProductMasuk, error)
	Delete(id string) (*models.ModelProductMasuk, error)
	Update(input *domain.UpdateProductMasukRequest) (*models.ModelProductMasuk, error)
}

type SaleRepository interface {
	Create(input *domain.CreateSaleRequest) (*models.ModelSale, error)
	Result(id string) (*models.ModelSale, error)
	Results() (*[]models.ModelSale, error)
	Delete(id string) (*models.ModelSale, error)
	Update(input *domain.UpdateSaleRequest) (*models.ModelSale, error)
}

type SupplierRepository interface {
	Create(input *domain.CreateSupplierRequest) (*models.ModelSupplier, error)
	Result(id string) (*models.ModelSupplier, error)
	Results() (*[]models.ModelSupplier, error)
	Delete(id string) (*models.ModelSupplier, error)
	Update(input *domain.UpdateSupplierRequest) (*models.ModelSupplier, error)
}
