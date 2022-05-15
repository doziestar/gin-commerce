package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key, autoincrement"  json:"id"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);not null"`
	Phone     string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	UpdatedAt string `gorm:"type:varchar(100);not null"`
	CreatedAt string `gorm:"type:varchar(100);not null"`
}

type UserModel struct {
	Db *gorm.DB
}

var UserModelInstance *UserModel

var authModel = new(AuthModel)

// GetUserModelInstance returns the singleton instance of the UserModel
func GetUserModelInstance() *UserModel {
	if UserModelInstance == nil {
		UserModelInstance = new(UserModel)
	}
	return UserModelInstance
}

// LoginUser returns the user if the user exists, create token and auth
func (userModel *UserModel) Login(email, password string) (user User, token Token, err error) {
	exist := UserModelInstance.Db.Where("email = ?", email).First(&user).RecordNotFound()
	if err != nil {
		return user, token, err
	}

	if exist {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return user, token, err
		}
	}

	TokenDetails, err := authModel.CreateToken(user.ID)
	if err != nil {
		return user, token, err
	}

	savedErr := authModel.CreateAuth(user.ID, TokenDetails)
	if savedErr != nil {
		return user, token, savedErr
	}

	token.AccessToken = TokenDetails.AccessToken
	token.RefreshToken = TokenDetails.RefreshToken

	return user, token, nil
}

// CreateUser creates a new user
func (userModel *UserModel) CreateUser(user User) (userId uint64, err error) {
	getDb := UserModelInstance.Db

	// Check if the user already exists
	existingUser := getDb.Where("email = ?", user.Email).First(&User{}).RecordNotFound()
	if !existingUser {
		return 0, nil
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(hashedPassword)

	// Create the user
	err = getDb.Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

// GetUser returns the user if the user exists
func (userModel *UserModel) GetUser(userId uint64) (user User, err error) {
	getDb := UserModelInstance.Db

	// Check if the user already exists
	existingUser := getDb.Where("id = ?", userId).First(&User{}).RecordNotFound()
	if existingUser {
		return user, nil
	}

	return user, nil
}
