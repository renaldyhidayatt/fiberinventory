package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
)

type UserRepository interface {
	Register(input *domain.UserInput) (*models.ModelUser, error)
	Login(input *domain.UserInput) (*models.ModelUser, error)
	Results() (*[]models.ModelUser, error)
	Result(input *domain.UserInput) (*models.ModelUser, error)
	Delete(input *domain.UserInput) (*models.ModelUser, error)
	Update(input *domain.UserInput) (*models.ModelUser, error)
	FindByEmail(email string) (*models.ModelUser, error)
}

type CategoryRepository interface {
	Create(input *domain.CategoryInput) (*models.ModelCategory, error)
	Results() (*[]models.ModelCategory, error)
	Result(input *domain.CategoryInput) (*models.ModelCategory, error)
	Delete(input *domain.CategoryInput) (*models.ModelCategory, error)
	Update(input *domain.CategoryInput) (*models.ModelCategory, error)
}

type CustomerRepository interface {
	Create(input *domain.CustomerInput) (*models.ModelCustomer, error)
	Results() (*[]models.ModelCustomer, error)
	Result(input *domain.CustomerInput) (*models.ModelCustomer, error)
	Delete(input *domain.CustomerInput) (*models.ModelCustomer, error)
	Update(input *domain.CustomerInput) (*models.ModelCustomer, error)
}

type ProductRepository interface {
	Create(input *domain.ProductInput) (*models.ModelProduct, error)
	Delete(input *domain.ProductInput) (*models.ModelProduct, error)
	Result(input *domain.ProductInput) (*models.ModelProduct, error)
	Update(input *domain.ProductInput) (*models.ModelProduct, error)
	Results() (*[]models.ModelProduct, error)
}

type ProductKeluarRepository interface {
	Create(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error)
	Result(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error)
	Results() (*[]models.ModelProductKeluar, error)
	Delete(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error)
	Update(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error)
}

type ProductMasukRepository interface {
	Create(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error)
	Result(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error)
	Results() (*[]models.ModelProductMasuk, error)
	Delete(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error)
	Update(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error)
}

type SaleRepository interface {
	Create(input *domain.SaleInput) (*models.ModelSale, error)
	Result(input *domain.SaleInput) (*models.ModelSale, error)
	Results() (*[]models.ModelSale, error)
	Delete(input *domain.SaleInput) (*models.ModelSale, error)
	Update(input *domain.SaleInput) (*models.ModelSale, error)
}

type SupplierRepository interface {
	Create(input *domain.SupplierInput) (*models.ModelSupplier, error)
	Result(input *domain.SupplierInput) (*models.ModelSupplier, error)
	Results() (*[]models.ModelSupplier, error)
	Delete(input *domain.SupplierInput) (*models.ModelSupplier, error)
	Update(input *domain.SupplierInput) (*models.ModelSupplier, error)
}
