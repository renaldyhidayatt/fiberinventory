package schemas

type SchemaProductKeluar struct {
	ID         string `json:"id" validate:"uuid"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
}
