package domain

type ResponseApiSuccess struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseApiError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
