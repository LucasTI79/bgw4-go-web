package dtos

import "github.com/lucasti79/bgw4-pratica-web/internal/domain"

// Data Transfer Object

type CreateProductInput struct {
	Name       string  `json:"name"`
	Quantity   *int    `json:"quantity"`
	Code       string  `json:"code_value"`
	Published  *bool   `json:"is_published"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
}

func (cp CreateProductInput) ToDomain() domain.Product {
	return domain.Product{
		Name:       cp.Name,
		Quantity:   cp.Quantity,
		Code:       cp.Code,
		Published:  cp.Published,
		Expiration: cp.Expiration,
		Price:      cp.Price,
	}
}

type UpdateProductInput struct {
	Name       string  `json:"name"`
	Quantity   *int    `json:"quantity"`
	Code       string  `json:"code_value"`
	Published  *bool   `json:"is_published"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
}

func (up UpdateProductInput) ToDomain() domain.Product {
	return domain.Product{
		Name:       up.Name,
		Quantity:   up.Quantity,
		Code:       up.Code,
		Published:  up.Published,
		Expiration: up.Expiration,
		Price:      up.Price,
	}
}
