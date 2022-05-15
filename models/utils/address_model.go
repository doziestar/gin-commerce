package models

import (
	models "gin-commerce/models/auth"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID  models.User `gorm:"foreignkey:UserID"`
	Address string      `gorm:"type:varchar(100);not null" json:"address"`
	City    string      `gorm:"type:varchar(100);not null" json:"city"`
	State   string      `gorm:"type:varchar(100);not null" json:"state"`
	Zip     string      `gorm:"type:varchar(100);not null" json:"zip"`
	Country string      `gorm:"type:varchar(100);not null" json:"country"`
}

type AddressModel struct {
	Db *gorm.DB
}

// get address by user id
func (addressModel *AddressModel) GetByUserID(userID uint) (addresses []Address, err error) {
	err = addressModel.Db.Where("user_id = ?", userID).Find(&addresses).Error
	return
}
