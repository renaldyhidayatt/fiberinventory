package schemas

type SchemaCategory struct {
	ID   string `json:"id" validate:"required,uuid"`
	Name string `json:"name" validate:"required,lowercase" schema:"name"`
}
