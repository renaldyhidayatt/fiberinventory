package schemas

type SchemaCustomer struct {
	ID      string `json:"id" validate:"required,uuid"`
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Email   string `json:"email" validate:"required,email"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}
