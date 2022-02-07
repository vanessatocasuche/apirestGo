package controller

import (
	"github.com/gin-gonic/gin"
	entity "github.com/vanessatocasuche/apirestGo/entity"
	"github.com/vanessatocasuche/apirestGo/service"
)

/**
ProductController is the interface that interacts between services with endpoints
This class has an instance of service.
*/

type ProductController interface {
	GetProduct(ctx *gin.Context) entity.Product
	GetAll() []entity.Product
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type productController struct {
	service service.ProductService
}

func New(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

func (thisController *productController) GetAll() []entity.Product {
	return thisController.service.GetAll()
}

func (thisController *productController) GetProduct(ctx *gin.Context) entity.Product {

	id := ctx.Param("id")
	if !thisController.service.ExistingID(id) {
		panic("ID " + id + " does not exist.")
	} else {
		for _, product := range thisController.service.GetAll() {
			if product.IdProduct == id {
				return product
			}
		}
	}
	panic("ID " + id + " does not exist.")
}

func (thisController *productController) Save(ctx *gin.Context) error {
	var product entity.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	product.IdProduct = id
	if thisController.service.ExistingID(id) {
		panic("ID " + id + " already exist.")
	} else {
		thisController.service.Save(product)
		return nil
	}
}

func (thisController *productController) Update(ctx *gin.Context) error {
	var product entity.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		panic("Error with JSON format delivered.")
	}

	id := ctx.Param("id")
	product.IdProduct = id
	if !thisController.service.ExistingID(id) {
		panic("ID " + id + " does not exist.")
	}

	//err = validate.Struct(product)
	if err != nil {
		return err
	}
	thisController.service.Update(product)
	return nil
}

func (thisController *productController) Delete(ctx *gin.Context) error {
	var product entity.Product
	id := ctx.Param("id")
	if !thisController.service.ExistingID(id) {
		panic("The ID " + id + " does not exist.")
	}
	product.IdProduct = id
	thisController.service.Delete(product)
	return nil
}
