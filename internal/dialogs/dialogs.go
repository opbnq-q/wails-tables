package dialogs

import (
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func InfoDialog(title string, message string) {
	application.InfoDialog().SetTitle(title).SetMessage(message).Show()
}

func WarningDialog(title string, message string) {
	application.WarningDialog().SetTitle(title).SetMessage(message).Show()
}

func ErrorDialog(title string, message string) {
	application.ErrorDialog().SetTitle(title).SetMessage(message).Show()
}

func SaveFileDialog(title string, filename string) (string, error) {
	selection, err := application.SaveFileDialog().SetFilename(filename).PromptForSingleSelection()
	if err != nil {
		return "", err
	}
	return selection, nil
}

func OpenFileDialog(title string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	dialog := application.OpenFileDialog()
	dialog.SetTitle(title)
	dialog.CanChooseDirectories(false)
	dialog.AddFilter("Электронные таблицы (.xlsx)", "*.xlsx")
	dialog.AddFilter("Электронные таблицы (.xls)", "*.xls")
	path, err := dialog.PromptForSingleSelection()
	if err != nil {
		return "", err
	}
	return path, nil
}
