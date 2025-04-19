package services

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

type ExcelModule struct{}

var ExcelService = application.NewService(&ExcelModule{})

func (s *ExcelModule) ImportAllEntities() error {
	// Example of usage:
	//postTypeService := PostTypeService{}
	//filepath, err := dialogs.OpenFileDialog("Импорт данных")
	//if err != nil {
	//	return err
	//}
	//err = excel.ImportEntitiesFromSpreadsheet(filepath, excel.Importer{
	//	SheetName: "Тип поста",
	//	Loader: func(rowIndex int, row []string) error {
	//		id, err := strconv.Atoi(row[0])
	//		if err != nil {
	//			return err
	//		}
	//
	//		_, err = postTypeService.Create(PostType{
	//			Id:   uint(id),
	//			Name: row[1],
	//		})
	//		if err != nil {
	//			return err
	//		}
	//		return nil
	//	},
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s *ExcelModule) ExportAllEntities() error {
	// Example of usage:
	//postService := PostService{}
	//exporter := excel.Exporter[Post]{
	//	SheetName: "Посты",
	//	Entity:    Post{},
	//	Provider:  postService.GetAll,
	//}
	//err := excel.ExportEntitiesToSpreadsheet("report.xlsx", exporter)
	//if err != nil {
	//	return err
	//}
	return nil
}
