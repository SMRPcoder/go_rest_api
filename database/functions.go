package database

import (
	"reflect"
	"time"

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
