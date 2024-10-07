package domain

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type RequestBodyProduct struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type ResponseBodyProduct struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
	Error   bool     `json:"error"`
}
