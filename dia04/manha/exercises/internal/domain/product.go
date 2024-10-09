package domain

// ponteiro assume dois valores => nil/value

type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Quantity   *int    `json:"quantity"`
	Code       string  `json:"code_value"`
	Published  *bool   `json:"is_published"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
}
