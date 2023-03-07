package domain

import "errors"

var (
	ErrorRegisterEmail  = errors.New("email already to login")
	ErrorRegisterFailed = errors.New("failed register ")
	ErrorLoginFailed    = errors.New("failed login correct email or password")

	// Category
	ErrorCategoryAlready      = errors.New("name category already")
	ErrorCategoryCreateFailed = errors.New("failed create category ")
	ErrorCategoryUpdateFailed = errors.New("failed update category ")
	ErrorCategorDeleteFailed  = errors.New("failed delete category ")

	ErrorCategoryNotFound = errors.New("failed result category")

	// Customer

	ErrorCustomerAlready      = errors.New("name customer already")
	ErrorCustomerCreateFailed = errors.New("failed create customer ")
	ErrorCustomerUpdateFailed = errors.New("failed update customer ")
	ErrorCustomerDeleteFailed = errors.New("failed delete customer ")
	ErrorCustomerNotFound     = errors.New("failed result customer")

	// product

	ErrorProductAlready      = errors.New("name product already")
	ErrorProductCreateFailed = errors.New("failed create product ")
	ErrorProductUpdateFailed = errors.New("failed update product ")
	ErrorProductDeleteFailed = errors.New("failed delete product ")

	ErrorProductNotFound = errors.New("failed result product")

	ErrorProductKeluarAlready      = errors.New("name product already")
	ErrorProductKeluarCreateFailed = errors.New("failed create product ")
	ErrorProductKeluarUpdateFailed = errors.New("failed update product ")
	ErrorProductKeluarDeleteFailed = errors.New("failed delete product ")

	ErrorProductKeluarNotFound = errors.New("failed result product")

	ErrorProductMasukAlready      = errors.New("name product already")
	ErrorProductMasukCreateFailed = errors.New("failed create product ")
	ErrorProductMasukUpdateFailed = errors.New("failed update product ")
	ErrorProductMasukDeleteFailed = errors.New("failed delete product ")

	ErrorProductMasukNotFound = errors.New("failed result product")

	ErrorSaleAlready      = errors.New("name sale already")
	ErrorSaleCreateFailed = errors.New("failed create sale ")
	ErrorSaleUpdateFailed = errors.New("failed update sale ")
	ErrorSaleDeleteFailed = errors.New("failed delete sale ")
	ErrorSaleNotFound     = errors.New("failed result sale")
	ErrorSalesNotFound    = errors.New("failed results sale")

	ErrorSupplierAlready      = errors.New("name supplier already")
	ErrorSupplierCreateFailed = errors.New("failed create supplier ")
	ErrorSupplierUpdateFailed = errors.New("failed update supplier ")
	ErrorSupplierDeleteFailed = errors.New("failed delete supplier ")
	ErrorSupplierNotFound     = errors.New("failed result supplier")
	ErrorSuppliersNotFound    = errors.New("failed results supplier")

	ErrorUserAlready      = errors.New("name supplier already")
	ErrorUserCreateFailed = errors.New("failed create supplier ")
	ErrorUserUpdateFailed = errors.New("failed update supplier ")
	ErrorUserDeleteFailed = errors.New("failed delete supplier ")
	ErrorUserNotFound     = errors.New("failed result supplier")
	ErrorUsersNotFound    = errors.New("failed results supplier")
)

type ErrorMessage struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      bool   `json:"error"`
}
