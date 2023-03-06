package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	User               UserRepository
	Category           CategoryRepository
	Customer           CustomerRepository
	Product            ProductRepository
	ProductKeluar      ProductKeluarRepository
	ProductMasuk       ProductMasukRepository
	SaleRepository     SaleRepository
	SupplierRepository SupplierRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User:               NewRepositoryUser(db),
		Category:           NewRepositoryCategory(db),
		Customer:           NewRepositoryCustomer(db),
		Product:            NewRepositoryProduct(db),
		ProductKeluar:      NewRepositoryProductKeluar(db),
		ProductMasuk:       NewRepositoryProductMasuk(db),
		SaleRepository:     NewRepositorySale(db),
		SupplierRepository: NewRepositorySupplier(db),
	}
}
