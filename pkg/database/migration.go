package database

import (
	models "gin-commerce/models/auth"
	models2 "gin-commerce/models/products"
	address "gin-commerce/models/utils"
)

//Add list of model add for migrations
//var migrationModels = []interface{}{&ex_models.Example{}, &model.Example{}, &model.Address{})}
var migrationModels = []interface{}{&models.User{}, &models2.Product{}, &models2.Cart{}, &models2.CartItem{}, &address.Address{}}
