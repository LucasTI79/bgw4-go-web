package products

import (
	"github.com/lucasti79/bgw4-pratica-web/internal/domain"
	"github.com/lucasti79/bgw4-pratica-web/pkg/apperrors"
)

type Repository interface {
	FindAll() ([]domain.Product, error)
	FindByID(id int) (domain.Product, error)
	FindByCode(code string) (*domain.Product, error)
	ExistsCode(code string) bool
	Create(product domain.Product) (domain.Product, error)
	Update(id int, product domain.Product) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	storage map[int]*domain.Product
	lastId  int
}

func NewRepository(storage map[int]*domain.Product) Repository {
	return &repository{
		storage: storage,
	}
}

func (r *repository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	for _, product := range r.storage {
		products = append(products, *product)
	}
	return products, nil
}

func (r *repository) FindByID(id int) (domain.Product, error) {
	product, ok := r.storage[id]
	if !ok {
		return domain.Product{}, apperrors.ErrNotFound
	}
	return *product, nil
}

func (r *repository) FindByCode(code string) (*domain.Product, error) {
	for _, product := range r.storage {
		if product.Code == code {
			return product, nil
		}
	}
	return nil, apperrors.ErrNotFound
}

func (r *repository) ExistsCode(code string) bool {
	for _, product := range r.storage {
		if product.Code == code {
			return true
		}
	}
	return false
}

func (r *repository) Create(product domain.Product) (domain.Product, error) {
	r.lastId++
	product.Id = r.lastId
	r.storage[r.lastId] = &product
	return product, nil
}

func (r *repository) Update(id int, product domain.Product) (domain.Product, error) {
	_, ok := r.storage[id]
	if !ok {
		return domain.Product{}, apperrors.ErrNotFound
	}
	product.Id = id
	r.storage[id] = &product
	return product, nil
}

func (r *repository) Delete(id int) error {
	_, ok := r.storage[id]
	if !ok {
		return apperrors.ErrNotFound
	}
	delete(r.storage, id)
	return nil
}
