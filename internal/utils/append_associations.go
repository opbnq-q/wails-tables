package utils

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

// AppendAssociations uses reflection to find all exported slice fields
// in 'item' and then appends any elements in each slice to the association.
func AppendAssociations(db *gorm.DB, item interface{}) error {
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

		// Process only exported fields (PkgPath == "") and slices.
		if fieldType.PkgPath == "" && fieldVal.Kind() == reflect.Slice {
			// The association name is the struct field name by default.
			assocName := fieldType.Name

			// Append slice elements to the existing association.
			// You can add checks here if you want to skip empty slices, etc.
			if fieldVal.Len() > 0 {
				if err := db.Model(item).Association(assocName).Append(fieldVal.Interface()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
