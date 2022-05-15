package models

import (
	"gorm.io/gorm"
)

// define enum for status
const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusRejected = "rejected"
	StatusCanceled = "canceled"
)

// Order model
type Order struct {
	gorm.Model
	UserID     uint        `gorm:"not null" json:"user_id"`
	Status     string      `gorm:"type:varchar(100);not null" json:"status"`
	TotalPrice float64     `gorm:"type:varchar(100);not null" json:"total_price"`
	Address    string      `gorm:"type:varchar(100);not null" json:"address"`
	Phone      string      `gorm:"type:varchar(100);not null" json:"phone"`
	Email      string      `gorm:"type:varchar(100);not null" json:"email"`
	OrderItems []OrderItem `gorm:"foreignkey:OrderID" json:"order_items"`
}

// OrderItem model
type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
}

type OrderModel struct {
	Db *gorm.DB
}

// GetAll function
func (orderModel *OrderModel) GetAll() (orders []Order, err error) {
	err = orderModel.Db.Find(&orders).Error
	return
}

// GetByID function
func (orderModel *OrderModel) GetByID(id uint) (order Order, err error) {
	err = orderModel.Db.First(&order, id).Error
	return
}

// Create function
func (orderModel *OrderModel) Create(order Order) (err error) {
	err = orderModel.Db.Create(&order).Error
	return
}

// Update function
func (orderModel *OrderModel) Update(order Order) (err error) {
	err = orderModel.Db.Save(&order).Error
	return
}

// Delete function
func (orderModel *OrderModel) Delete(id uint) (err error) {
	err = orderModel.Db.Delete(&Order{}, id).Error
	return
}

// GetAllByUserID function
func (orderModel *OrderModel) GetAllByUserID(userID uint) (orders []Order, err error) {
	err = orderModel.Db.Where("user_id = ?", userID).Find(&orders).Error
	return
}

// GetAllByStatus function
func (orderModel *OrderModel) GetAllByStatus(status string) (orders []Order, err error) {
	err = orderModel.Db.Where("status = ?", status).Find(&orders).Error
	return
}
