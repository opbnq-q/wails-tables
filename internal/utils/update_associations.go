package utils

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

func UpdateAssociations(db *gorm.DB, item interface{}) error {
	// We expect a pointer to a struct so we can do db.Model(item).
	val := reflect.ValueOf(item)
	if val.Kind() != reflect.Ptr {
		return errors.New("item must be a pointer to a struct")
	}

	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("item must be a pointer to a struct")
	}

	t := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		fieldVal := elem.Field(i)
		fieldType := t.Field(i)

		// Only process exported fields (capitalized) and slices.
		if fieldType.PkgPath == "" && fieldVal.Kind() == reflect.Slice {
			// For clarity, the association name is the struct field name by default.
			assocName := fieldType.Name

			// Replace the association with the current slice value.
			// If you only want to replace on non-nil or non-empty slices, you can add extra checks here.
			if err := db.Model(item).Association(assocName).Replace(fieldVal.Interface()); err != nil {
				return err
			}
		}
	}

	return nil
}
