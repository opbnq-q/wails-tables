package services

import (
	"fmt"
	"wails_tables/internal/dialogs"
	"wails_tables/internal/models"
)

func InsertDefaultData() {
	insertWoodSpecTypes()
}

func insertWoodSpecTypes() {
	InsertDefaultEntityData(&WoodSpecTypeService{}, []models.WoodSpecType{
		{
			Id:   1,
			Name: "Круглый лес",
		},
		{
			Id:   2,
			Name: "Сырые пиломатериалы",
		},
		{
			Id:   3,
			Name: "Сухие пиломатериалы",
		},
		{
			Id:   4,
			Name: "Опилки",
		},
		{
			Id:   5,
			Name: "Пеллеты",
		},
	})
}

func InsertDefaultEntityData[T any](service Service[T], entities []T) {
	for _, item := range entities {
		createdItem, err := service.Create(item)
		if err != nil {
			dialogs.ErrorDialog("Ошибка при вставке данных по умолчанию", fmt.Sprintf("Произошла ошибка при вставке значения %#v: %s", createdItem, err))
		}
	}
}
