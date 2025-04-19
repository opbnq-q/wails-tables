package utils

import (
	"fmt"
	"gorm.io/gorm/clause"
	"reflect"
	"wails_tables/internal/database"
)

func FindPhraseByStringFields[T any](phrase string, entity T) ([]*T, error) {
	db := database.GetInstance().Preload(clause.Associations)
	structType := reflect.TypeOf(entity)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		if field.Type.Kind() == reflect.Pointer {
			field.Type = field.Type.Elem()
		}

		if field.Type.Kind() == reflect.String {
			db.Where(fmt.Sprintf("`%s` like ?", field.Name), "%"+phrase+"%")
		}
	}

	var items []*T
	db.Find(&items)
	return items, nil
}
