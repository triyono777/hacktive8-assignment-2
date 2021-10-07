package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Items struct {
	ItemId uint `gorm:"primaryKey" json:"item_id"`
	ItemCode string `json:"item_code"`
	Description string `json:"description"`
	OrderId uint `json:"order_id"`

}



func (items *Items) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(items)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
func (items *Items) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(items)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
