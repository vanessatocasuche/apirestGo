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

func (thisService *productService) Save(p entity.Product) error {
	thisService.repository.Save(p)
	return nil
}

func (thisService *productService) Update(p entity.Product) error {
	thisService.repository.Update(p)
	return nil
}

func (thisService *productService) Delete(p entity.Product) error {
	thisService.repository.Delete(p)
	return nil
}

func (thisService *productService) GetAll() []entity.Product {
	return thisService.repository.GetAll()
}

func (thisService *productService) ExistingID(id string) bool {
	return thisService.repository.ExistingID(id)
}
