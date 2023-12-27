package functions

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func SetCreatedAt(model interface{}) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		// Check if the provided model is a pointer to a struct
		value := reflect.ValueOf(model)
		if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
			return
		}
		// Get the "UpdatedAt" field and set its value to the current time
		createdAtField := value.Elem().FieldByName("CreatedAt")
		if createdAtField.IsValid() && createdAtField.CanSet() && createdAtField.Type() == reflect.TypeOf(time.Now()) {
			createdAtField.Set(reflect.ValueOf(time.Now()))
		}
	}
}

func SetUpdatedAt(model interface{}) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		// Check if the provided model is a pointer to a struct
		value := reflect.ValueOf(model)
		if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
			return
		}

		// Get the "UpdatedAt" field and set its value to the current time
		updatedAtField := value.Elem().FieldByName("UpdatedAt")
		if updatedAtField.IsValid() && updatedAtField.CanSet() && updatedAtField.Type() == reflect.TypeOf(time.Now()) {
			updatedAtField.Set(reflect.ValueOf(time.Now()))
		}
	}
}

type JWTUser struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
}

func EncodeJwt(user JWTUser) (string, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return "", errors.New("loading env error")
	}

	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	mySecretKey := os.Getenv("MY_SECRET_JWT")
	key = []byte(mySecretKey)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{"username": user.Username, "name": user.Name, "id": user.ID})
	str, err := t.SignedString(key)
	if err != nil {
		return "", errors.New(err.Error())
	}
	s = str
	return s, nil
}
