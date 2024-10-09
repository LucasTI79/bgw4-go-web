package storage

import (
	"fmt"

	"github.com/lucasti79/bgw4-put-patch-delete/internal/domain"
	"github.com/lucasti79/bgw4-put-patch-delete/pkg/apperrors"
)

type ProductAttributes struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type ProductsMap struct {
	db     map[int]ProductAttributes
	lastId int
}

func NewProductsStorage(db map[int]ProductAttributes) domain.Respository {
	return &ProductsMap{db: db, lastId: 0}
}

func (pm *ProductsMap) Get() ([]domain.Product, error) {
	var products []domain.Product
	for id, data := range pm.db {
		products = append(products, domain.Product{
			Id:       id,
			Name:     data.Name,
			Type:     data.Type,
			Quantity: data.Quantity,
			Price:    data.Price,
		})
	}
	return products, nil
}

func (pm *ProductsMap) GetByID(id int) (*domain.Product, error) {
	// search
	data, ok := pm.db[id]
	if !ok {
		return nil, fmt.Errorf("%w: %d", apperrors.ErrNotFound, id)
	}
	return &domain.Product{
		Id:       id,
		Name:     data.Name,
		Type:     data.Type,
		Quantity: data.Quantity,
		Price:    data.Price,
	}, nil
}

func (pm *ProductsMap) UpdateOrCreate(p *domain.Product) (err error) {
	// serialize
	attr := ProductAttributes{Name: p.Name, Type: p.Type, Quantity: p.Quantity, Price: p.Price}
	// update
	_, ok := pm.db[p.Id]
	switch ok {
	case true:
		pm.db[p.Id] = attr
	default:
		pm.lastId++
		pm.db[pm.lastId] = attr
	}
	return
}

func (pm *ProductsMap) Update(id int, p *domain.Product) (err error) {
	// search
	if _, ok := pm.db[id]; !ok {
		err = fmt.Errorf("%w: %d", apperrors.ErrNotFound, id)
		return
	}
	// serialize
	attr := ProductAttributes{Name: p.Name, Type: p.Type, Quantity: p.Quantity, Price: p.Price}
	// update
	pm.db[id] = attr
	return
}

func (pm *ProductsMap) Delete(id int) (err error) {
	// search
	if _, ok := pm.db[id]; !ok {
		err = fmt.Errorf("%w: %d", apperrors.ErrNotFound, id)
		return
	}
	// delete
	delete(pm.db, id)
	return
}
