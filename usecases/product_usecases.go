package usecases

import (
	"go-api/models"
	"go-api/repositories"
)

type ProductUsecase struct {
	repository repositories.ProductRepository
}

func NewProductUsecase(repository repositories.ProductRepository) ProductUsecase {
	return ProductUsecase{repository}
}

func (p *ProductUsecase) GetProducts() ([]models.Product, error) {
	return p.repository.GetProducts()
}

func (p *ProductUsecase) CreateProduct(product models.Product) (models.Product, error) {
	id, err := p.repository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	product.Id = id

	return product, nil
}

func (p *ProductUsecase) GetProductById(id int) (*models.Product, error) {
	product, err := p.repository.GetProductById(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
