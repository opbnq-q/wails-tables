package services

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

var ExportedServices = []application.Service{
	application.NewService(&WoodSpecTypeService{}),
}
