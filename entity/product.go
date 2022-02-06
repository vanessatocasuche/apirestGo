package models

import "time"

/**
Basic class with attributes of the product
*/

type Product struct {
	IdProduct       string    `json:"id_product" gorm:"primary_key" gorm:"type:varchar(36)"`
	Name            string    `json:"name" binding:"required" gorm:"default:not_null" gorm:"type:varchar(45)"`
	Description     string    `json:"description" gorm:"type:varchar(300)"`
	Status          string    `json:"status" gorm:"type:varchar(45)" binding:"required" gorm:"default:not_null"`
	CreationDate    time.Time `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateDate      time.Time `json:"update_date" gorm:"default:CURRENT_TIMESTAMP"`
	AccountId       string    `json:"account_id" gorm:"type:varchar(25)" gorm:"default:not_null" binding:"required" `
	FormatProduct   string    `json:"format_product" gorm:"type:varchar"`
	ValueUnit       float32   `json:"value_unit" binding:"required" gorm:"default:not_null" gorm:"type:DECIMAL"`
	UnitName        string    `json:"unit_name" binding:"required" gorm:"default:not_null" gorm:"type:varchar(45)"`
	UnitDescription string    `json:"unit_description" gorm:"type:varchar(300)"`
	Stock           int       `json:"stock" gorm:"type:int" gorm:"default:not_null" binding:"required"   `
}
