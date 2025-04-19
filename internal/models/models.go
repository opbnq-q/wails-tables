package models

var Entities = []any{
	&WoodSpecType{},
}

type WoodSpecType struct {
	Id   uint   `gorm:"primaryKey" ui:"hidden"`
	Name string `ui:"label:Название"`
}
