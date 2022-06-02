package schemas

type Configuration struct {
	DB struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Pass     string `json:"password"`
		Database string `json:"database"`
	}
}
