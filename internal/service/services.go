package service

import (
	"fiberinventory/internal/repository"
	"fiberinventory/pkg/auth"
	"fiberinventory/pkg/hash"
)

type Service struct {
	User          UserService
	Category      CategoryService
	Customer      CustomerService
	Product       ProductService
	ProductKeluar ProductKeluarService
	ProductMasuk  ProductMasukService
	Sale          SaleService
	Supplier      SupplierService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    hash.Hashing
	Token      auth.TokenManager
}

func NewService(deps Deps) *Service {
	return &Service{
		User:          NewServiceUser(deps.Repository.User, deps.Hashing, deps.Token),
		Category:      NewServiceCategory(deps.Repository.Category),
		Customer:      NewServiceCustomer(deps.Repository.Customer),
		Product:       NewServiceProduct(deps.Repository.Product),
		ProductKeluar: NewServiceProductKeluar(deps.Repository.ProductKeluar),
		ProductMasuk:  NewServiceProductMasuk(deps.Repository.ProductMasuk),
		Sale:          NewServiceSale(deps.Repository.SaleRepository),
		Supplier:      NewServiceSupplier(deps.Repository.SupplierRepository),
	}
}
