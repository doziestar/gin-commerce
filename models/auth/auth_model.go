package models

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	uuid "github.com/twinj/uuid"
)

type AuthModel struct {
	Db *gorm.DB
}

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"access_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

type AccessDetails struct {
	AccessUuid string `json:"access_uuid"`
	UserId     int    `json:"user_id"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (auth *AuthModel) CreateToken(userId uint64) (*TokenDetails, error) {
	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV1().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV1().String()

	var err error

	// Create the JWT key
	key := []byte("my_secret_key")

	// Create the Claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["refresh_token"] = td.RefreshToken

	// Create token
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign token and return
	td.RefreshToken, err = rt.SignedString(key)
	if err != nil {
		return nil, err
	}

	return td, nil
}

// create Auth
func (auth *AuthModel) CreateAuth(userId uint64, td *TokenDetails) error {
	td, err := auth.CreateToken(userId)
	if err != nil {
		return nil
	}

	access := &AccessDetails{
		AccessUuid: td.AccessUuid,
		UserId:     int(userId),
	}

	err = auth.Db.Create(access).Error
	if err != nil {
		return nil
	}

	return nil
}

// Extract token from request
func (auth *AuthModel) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// Verify token
func (auth *AuthModel) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := auth.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Token valid
func (auth *AuthModel) TokenValid(r *http.Request) error {
	token, err := auth.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// Extract Token Meta Data
func (auth *AuthModel) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := auth.VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		access := &AccessDetails{}
		err = auth.Db.Where("access_uuid = ?", accessUuid).First(access).Error
		if err != nil {
			return nil, err
		}

		return access, nil
	}
	return nil, err
}

// Fetch Auth
func (auth *AuthModel) FetchAuth(r *http.Request) (*AccessDetails, error) {
	access, err := auth.ExtractTokenMetadata(r)
	if err != nil {
		return nil, err
	}
	return access, nil
}

// Delete Auth
func (auth *AuthModel) DeleteAuth(r *http.Request) error {
	access, err := auth.ExtractTokenMetadata(r)
	if err != nil {
		return err
	}

	err = auth.Db.Delete(access).Error
	if err != nil {
		return err
	}

	return nil
}
