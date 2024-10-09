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

type RequestBodyUpdateOrCreate struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type RequestBodyUpdate struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Respository interface {
	// Get returns all the products
	Get() (p []Product, err error)
	// GetByID returns a product by id
	GetByID(id int) (p *Product, err error)
	// Save saves a product
	// Save(p *Product) (err error)
	// UpdateOrCreate updates or creates a product if it does not exist
	UpdateOrCreate(p *Product) (err error)

	Update(id int, p *Product) (err error)

	Delete(id int) (err error)
}
