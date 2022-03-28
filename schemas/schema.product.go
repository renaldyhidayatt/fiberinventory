package schemas

type SchemaProduct struct {
	ID         string `json:"id" validate:"uuid"`
	Name       string `json:"name" validate:"required,lowercase"`
	Image      string `json:"image" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
}
