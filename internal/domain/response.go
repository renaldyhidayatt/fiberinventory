package domain

type Response struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}

type ResponseAuth struct {
	Message    string      `json:"message"`
	Jwt        interface{} `json:"jwt"`
	StatusCode int         `json:"statusCode"`
}
