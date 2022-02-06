package service

import (
	entity "github.com/vanessatocasuche/apirestGo/entity"
	"github.com/vanessatocasuche/apirestGo/repository"
)

/**

 */

type ProductService interface {
	Save(entity.Product) error
	Update(entity.Product) error
	Delete(entity.Product) error
	GetAll() []entity.Product
	ExistingID(string) bool
}

type productService struct {
	repository repository.ProductRepository
}

func New(productRepository repository.ProductRepository) ProductService {
	return &productService{
		repository: productRepository,
	}
}

func (service *productService) Save(p entity.Product) error {
	service.repository.Save(p)
	return nil
}

func (service *productService) Update(p entity.Product) error {
	service.repository.Update(p)
	return nil
}

func (service *productService) Delete(p entity.Product) error {
	service.repository.Delete(p)
	return nil
}

func (service *productService) GetAll() []entity.Product {
	return service.repository.GetAll()
}

func (service *productService) ExistingID(id string) bool {
	return service.repository.ExistingID(id)
}
