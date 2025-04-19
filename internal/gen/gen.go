package main

import (
	"wails_tables/internal/models"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../dal",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable: true,
		WithUnitTest:  true,
	})
	g.ApplyBasic(models.Entities...)
	g.Execute()
}
