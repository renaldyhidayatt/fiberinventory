package schemas

type SchemaProductMasuk struct {
	ID         string `json:"id" validate:"uuid"`
	Name       string `json:"name" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	SupplierID string `json:"supplier_id" validate:"required"`
}
