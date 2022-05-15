package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint    `gorm:"primary_key autoincrement" json:"id"`
	Title       string  `gorm:"type:varchar(100);not null" json:"title"`
	Description string  `gorm:"type:varchar(100);not null" json:"description"`
	Price       float64 `gorm:"type:varchar(100);not null" json:"price"`
	Quantity    int     `gorm:"type:varchar(100);not null" json:"quantity"`
	Image       string  `gorm:"type:varchar(100);not null" json:"image"`
	CategoryID  uint    `gorm:"type:varchar(100);not null" json:"category_id"`
	CreatedAt   string  `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt   string  `gorm:"type:varchar(100);not null" json:"updated_at"`
}

type ProductModel struct {
	Db *gorm.DB
}

var ProductModelInstance *ProductModel

func GetProductModelInstance() *ProductModel {
	if ProductModelInstance == nil {
		ProductModelInstance = new(ProductModel)
	}
	return ProductModelInstance
}

func (productModel *ProductModel) GetAll() (products []Product, err error) {
	err = productModel.Db.Find(&products).Error
	return
}

func (productModel *ProductModel) GetByID(id uint) (product Product, err error) {
	err = productModel.Db.First(&product, id).Error
	return
}

func (productModel *ProductModel) Create(product Product) (err error) {
	err = productModel.Db.Create(&product).Error
	return
}

func (productModel *ProductModel) Update(product Product) (err error) {
	err = productModel.Db.Save(&product).Error
	return
}

func (productModel *ProductModel) Delete(id uint) (err error) {
	err = productModel.Db.Delete(&Product{}, id).Error
	return
}
