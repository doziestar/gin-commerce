package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint       `gorm:"not null" json:"user_id"`
	Status     string     `gorm:"type:varchar(100);not null" json:"status"`
	TotalPrice float64    `gorm:"type:varchar(100);not null" json:"total_price"`
	Address    string     `gorm:"type:varchar(100);not null" json:"address"`
	Phone      string     `gorm:"type:varchar(100);not null" json:"phone"`
	Email      string     `gorm:"type:varchar(100);not null" json:"email"`
	CartItem   []CartItem `gorm:"foreignkey:CartID" json:"cart_item"`
}

type CartItem struct {
	gorm.Model
	CartID    uint    `gorm:"not null" json:"cart_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
}

type CartModel struct {
	Db *gorm.DB
}

// delete cart
func (cartModel *CartModel) Delete(id uint) (err error) {
	err = cartModel.Db.Delete(&Cart{}, id).Error
	return
}

// pay cart
func (cartModel *CartModel) Pay(id uint) (err error) {
	err = cartModel.Db.Model(&Cart{}).Where("id = ?", id).Update("status", "paid").Error
	return
}

// GetAllByUserID function
func (cartModel *CartModel) GetAllByUserID(userID uint) (carts []Cart, err error) {
	err = cartModel.Db.Where("user_id = ?", userID).Find(&carts).Error
	return
}

// deliver cart
func (cartModel *CartModel) Deliver(id uint) (err error) {
	err = cartModel.Db.Model(&Cart{}).Where("id = ?", id).Update("status", "delivered").Error
	return
}
