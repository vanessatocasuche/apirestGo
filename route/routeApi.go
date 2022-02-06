package route

import (
	"github.com/gin-gonic/gin"
	"github.com/vanessatocasuche/apirestGo/controller"
	con "github.com/vanessatocasuche/apirestGo/controller"
	"github.com/vanessatocasuche/apirestGo/repository"
	"github.com/vanessatocasuche/apirestGo/service"
	"net/http"
)

var (
	productRepository repository.ProductRepository
	productService    service.ProductService
	productController controller.ProductController
)

/**
EndPintsApp is a function with the scheme and the endpoints of the app
*/

func EndPointsApp(route *gin.Engine) {
	//Scheme instances of the app
	productRepository = repository.NewDBRepository()
	productService = service.New(productRepository)
	productController = controller.New(productService)

	// EndPoints of the App
	route.GET("/products", GetProducts)
	route.GET("/products/:id", GetProduct)
	route.POST("/products", CreateProduct)
	route.PUT("/products/:id", UpdateProduct)
	route.DELETE("/product/:id", DeleteProduct)
	route.GET("/volume", gin.BasicAuth(getUsers()), GetFileVolume)
}

/**
Here are the basic functions offered to application.
This class supports the Main package.
*/

func GetProducts(ctx *gin.Context) {
	ctx.JSON(200, productController.GetAll())
}

func GetProduct(ctx *gin.Context) {
	ctx.JSON(200, productController.GetProduct(ctx))
}

func CreateProduct(ctx *gin.Context) {
	err := productController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, "Success!")

	}
}

func UpdateProduct(ctx *gin.Context) {
	err := productController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, "Success!")
	}
}

func DeleteProduct(ctx *gin.Context) {
	err := productController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, "Success!")
	}
}

//Users of auth

func getUsers() map[string]string {

	m := make(map[string]string) // users and passwords are staying admitted to basicAuth

	m["admin"] = "admin1"
	m["vane"] = "vane1"

	return m
}

/*
Function to help the endpoint get file Volume, and it connects with controller
*/

func GetFileVolume(ctx *gin.Context) {
	ctx.JSON(200, con.GetFile())
}
