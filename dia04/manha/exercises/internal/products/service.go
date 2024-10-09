package products

import (
	"errors"

	"github.com/lucasti79/bgw4-pratica-web/internal/domain"
	"github.com/lucasti79/bgw4-pratica-web/pkg/apperrors"
)

type Service interface {
	FindAll() ([]domain.Product, error)
	FindByID(id int) (domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	Update(id int, product domain.Product) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) FindAll() ([]domain.Product, error) {
	return s.repo.FindAll()
}

func (s *service) FindByID(id int) (domain.Product, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(product domain.Product) (domain.Product, error) {
	codeAlreadyExists := s.repo.ExistsCode(product.Code)

	if codeAlreadyExists {
		return domain.Product{}, ErrCodeAlreadyExists
	}

	return s.repo.Create(product)
}

func (s *service) Update(id int, product domain.Product) (domain.Product, error) {
	productMatchedWithCode, err := s.repo.FindByCode(product.Code)

	if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
		return domain.Product{}, err
	}

	if productMatchedWithCode != nil && productMatchedWithCode.Id != id {
		return domain.Product{}, ErrCodeAlreadyExists
	}

	return s.repo.Update(id, product)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
