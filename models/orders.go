package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	OrderId      uint       `gorm:"primaryKey"  json:"order_id"`
	CustomerName string     `json:"customer_name"`
	OrderedAt    *time.Time `json:"ordered_at"`
}

func (orders *Orders) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(orders)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
func (orders *Orders) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(orders)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
