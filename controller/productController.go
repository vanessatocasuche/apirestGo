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

func (c *productController) GetAll() []entity.Product {
	return c.service.GetAll()
}

func (c *productController) GetProduct(ctx *gin.Context) entity.Product {

	id := ctx.Param("id")
	if !c.service.ExistingID(id) {
		panic("ID " + id + " does not exist.")
	} else {
		for _, product := range c.service.GetAll() {
			if product.IdProduct == id {
				return product
			}
		}
	}
	panic("ID " + id + " does not exist.")
}

func (c *productController) Save(ctx *gin.Context) error {
	var product entity.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		return err
	}
	id := ctx.Param("id")

	if c.service.ExistingID(id) {
		panic("ID " + id + " already exist.")
	} else {
		c.service.Save(product)
		return nil
	}
}

func (c *productController) Update(ctx *gin.Context) error {
	var product entity.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		panic("Error with JSON format delivered.")
	}

	id := ctx.Param("id")
	product.IdProduct = id
	if !c.service.ExistingID(id) {
		panic("ID " + id + " does not exist.")
	}

	//err = validate.Struct(product)
	if err != nil {
		return err
	}
	c.service.Update(product)
	return nil
}

func (c *productController) Delete(ctx *gin.Context) error {
	var product entity.Product
	id := ctx.Param("id")
	if !c.service.ExistingID(id) {
		panic("The ID " + id + " does not exist.")
	}
	product.IdProduct = id
	c.service.Delete(product)
	return nil
}
