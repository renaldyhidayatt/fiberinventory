package schemas

type SchemaDatabaseError struct {
	Type string
	Code int
}

type SchemaUnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
