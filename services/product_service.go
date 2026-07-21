package services

import (
	"errors"

	"myapi/models"
	"myapi/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{
		repo: &repositories.ProductRepository{},
	}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) Create(req models.CreateProductRequest) (*models.Product, error) {
	categories, err := s.repo.FindCategoriesByIDs(req.Categories)
	if err != nil {
		return nil, errors.New("invalid category IDs")
	}

	product := models.Product{
		Name:        req.Name,
		Quantity:    req.Quantity,
		Categories:  categories,
		Price:       req.Price,
		Description: req.Description,
	}

	if err := s.repo.Create(&product); err != nil {
		return nil, err
	}

	s.repo.PreloadCategories(&product)
	return &product, nil
}

func (s *ProductService) Update(id int, req models.CreateProductRequest) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Quantity = req.Quantity
	product.Price = req.Price
	product.Description = req.Description

	if len(req.Categories) > 0 {
		categories, err := s.repo.FindCategoriesByIDs(req.Categories)
		if err != nil {
			return nil, errors.New("invalid category IDs")
		}
		s.repo.ReplaceCategories(product, categories)
	}

	if err := s.repo.Save(product); err != nil {
		return nil, err
	}

	s.repo.PreloadCategories(product)
	return product, nil
}

func (s *ProductService) Delete(id int) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	s.repo.ClearCategories(product)
	return s.repo.Delete(product)
}
