package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	entity "github.com/vanessatocasuche/apirestGo/entity"
	"github.com/vanessatocasuche/apirestGo/repository"
	"github.com/vanessatocasuche/apirestGo/service"
	"time"
)

/*
Testing with the libraries Ginkgo y Gomega.

To execute this test, you must locate it in the terminal
where this file is and execute the command "go test"
*/

// Variables and constants needed to test crude services

const (
	IDPRODUCT               = "P01"
	NAME                    = "P01"
	DESCRIPTION             = "P01"
	STATUS                  = "P01"
	ACCOUNTID               = "P01"
	FORMATPRODUCT           = "P01"
	VALUEUNIT       float32 = 100.1
	UNITNAME                = "P01"
	UNITDESCRIPTION         = "P01"
	STOCK                   = 100

	IDPRODUCT2               = "P01"
	NAME2                    = "P01_update"
	DESCRIPTION2             = "P01_update"
	STATUS2                  = "P01_update"
	ACCOUNTID2               = "P01_update"
	FORMATPRODUCT2           = "P01_update"
	VALUEUNIT2       float32 = 99.9
	UNITNAME2                = "P01_update"
	UNITDESCRIPTION2         = "P01_update"
	STOCK2                   = 99
)

var (
	loc, err = time.LoadLocation("America/Bogota")

	CREATIONDATE  = time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	UPDATEDATE    = time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	CREATIONDATE2 = time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	UPDATEDATE2   = time.Date(2001, time.October, 10, 05, 0, 0, 0, loc)
)

var productTest1 = entity.Product{
	IdProduct:       IDPRODUCT,
	Name:            NAME,
	Description:     DESCRIPTION,
	Status:          STATUS,
	CreationDate:    CREATIONDATE,
	UpdateDate:      UPDATEDATE,
	AccountId:       ACCOUNTID,
	FormatProduct:   FORMATPRODUCT,
	ValueUnit:       VALUEUNIT,
	UnitName:        UNITNAME,
	UnitDescription: UNITDESCRIPTION,
	Stock:           STOCK,
}

var productTest2 = entity.Product{
	IdProduct:       IDPRODUCT2,
	Name:            NAME2,
	Description:     DESCRIPTION2,
	Status:          STATUS2,
	CreationDate:    CREATIONDATE2,
	UpdateDate:      UPDATEDATE2,
	AccountId:       ACCOUNTID2,
	FormatProduct:   FORMATPRODUCT2,
	ValueUnit:       VALUEUNIT2,
	UnitName:        UNITNAME2,
	UnitDescription: UNITDESCRIPTION2,
	Stock:           STOCK2,
}

// Development of the test to 'productService'

var _ = Describe("PRUEBA A 'ServiceProduct'", func() {

	var (
		productService    service.ProductService
		productRepository repository.ProductRepository
	)

	BeforeSuite(func() {
		productRepository = repository.NewDBRepository()
		productService = service.New(productRepository)

	})

	Describe("CRUD", func() {

		Context("EXTRAER TODOS LOS LEMENTOS (GET)", func() {
			BeforeEach(func() {
				productService.Save(productTest1)
				productService.Save(productTest2)
			})
			It("M??todo GET no retorn?? al menos 1 elemento.", func() {
				products := productService.GetAll()
				??(products).ShouldNot(BeEmpty())
			})
		})

		Context("GUARDAR ELEMENTO (POST)", func() {
			BeforeEach(func() {
				productService.Save(productTest1)
			})
			It("ID no encontrado: El producto no se guard??", func() {
				for _, p := range productService.GetAll() {
					if p.IdProduct == productTest1.IdProduct {
						??(p.IdProduct).Should(Equal(productTest1.IdProduct))
					}
				}
			})
			It("Error de coincidencia: El elemento se guard?? con datos diferentes", func() {
				products := productService.GetAll()
				for _, product := range products {
					if product.IdProduct == productTest1.IdProduct {
						??(product.IdProduct).Should(Equal(IDPRODUCT))
						??(product.Name).Should(Equal(NAME))
						??(product.Description).Should(Equal(DESCRIPTION))
						??(product.Status).Should(Equal(STATUS))
						??(product.AccountId).Should(Equal(ACCOUNTID))
						??(product.FormatProduct).Should(Equal(FORMATPRODUCT))
						??(product.ValueUnit).Should(Equal(VALUEUNIT))
						??(product.UnitName).Should(Equal(UNITNAME))
						??(product.UnitDescription).Should(Equal(UNITDESCRIPTION))
						??(product.Stock).Should(Equal(STOCK))
						??(product.CreationDate.Local()).Should(Equal(CREATIONDATE.Local()))
						??(product.UpdateDate.Local()).Should(Equal(UPDATEDATE.Local()))
					}
				}
			})
			BeforeEach(func() {
				productService.Save(productTest1)
			})
			AfterEach(func() {
				productService.Delete(productTest1)
			})
		})

		Context("ACTUALIZAR ELEMENTO (UPDATE)", func() {
			BeforeEach(func() {
				productService.Save(productTest1)
				productService.Update(productTest2)
			})
			It("ID no encontrado: El producto no se guard??, o fall?? la actualizaci??n con el ID", func() {
				for _, p := range productService.GetAll() {
					if p.IdProduct == productTest2.IdProduct {
						??(p.IdProduct).Should(Equal(productTest2.IdProduct))
					}
				}
			})
			It("Error de coincidencia: El elemento se actualiz?? con datos diferentes", func() {
				products := productService.GetAll()
				for _, product := range products {
					if product.IdProduct == productTest2.IdProduct {
						??(product.IdProduct).Should(Equal(IDPRODUCT2))
						??(product.Name).Should(Equal(NAME2))
						??(product.Description).Should(Equal(DESCRIPTION2))
						??(product.Status).Should(Equal(STATUS2))
						??(product.AccountId).Should(Equal(ACCOUNTID2))
						??(product.FormatProduct).Should(Equal(FORMATPRODUCT2))
						??(product.ValueUnit).Should(Equal(VALUEUNIT2))
						??(product.UnitName).Should(Equal(UNITNAME2))
						??(product.UnitDescription).Should(Equal(UNITDESCRIPTION2))
						??(product.Stock).Should(Equal(STOCK2))
						??(product.CreationDate.Local()).Should(Equal(CREATIONDATE2.Local()))
						??(product.UpdateDate.Local()).Should(Equal(UPDATEDATE2.Local()))
					}
				}
			})
			AfterEach(func() {
				productService.Delete(productTest2)
				productService.Delete(productTest1)
			})
		})

		Context("BORRAR ELEMENTO (DELETE)", func() {
			BeforeEach(func() {
				productService.Save(productTest1)
				productService.Delete(productTest1)

			})
			It("ID encontrado: El elemento no se borr?? de la BD", func() {
				for _, p := range productService.GetAll() {
					if p.IdProduct == productTest1.IdProduct {
						??(p.IdProduct).Should(Equal(productTest1.IdProduct))
					}
				}
			})
		})

	})
})
