package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	entity "github.com/vanessatocasuche/apirestGo/entity"
)

/**
This file called repository makes the connection between
the database and the services of the application.

Here the methods provided by the "Gorm" library are used
who provides the persistence of data, among other things
*/

type ProductRepository interface {
	Save(product entity.Product)
	Update(product entity.Product)
	Delete(product entity.Product)
	GetAll() []entity.Product
	ExistingID(string) bool
}

type database struct {
	connection *gorm.DB
}

/**
NewDBRepository provides the connection with the SQLite database
*/
const locationDB string = ".././../../db/db-apirest-go.db"

func NewDBRepository() ProductRepository {

	db, err := gorm.Open("sqlite3", "../../../../db/db-apirest-go.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Product{})
	return &database{
		connection: db,
	}
}

/**
Methods and functions to execute on the database directly
*/

func (db *database) Save(product entity.Product) {
	db.connection.Create(&product)
}

func (db *database) Update(product entity.Product) {
	db.connection.Save(&product)
}

func (db *database) Delete(product entity.Product) {
	db.connection.Delete(&product)
}

func (db *database) GetAll() []entity.Product {
	var products []entity.Product
	db.connection.Find(&products)
	return products
}

func (db *database) ExistingID(id string) bool {
	for _, p := range db.GetAll() {
		if p.IdProduct == id {
			return true
		}
	}
	return false
}
