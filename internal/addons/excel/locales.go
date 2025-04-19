package excel

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	excelAddon()
}

func excelAddon() {
	message.SetString(language.Russian, "Error while closing excel file: %v", "Ошибка при закрытии файла Excel: %v")
	message.SetString(language.Russian, "Field %s not found", "Поле %s не найдено")
	message.SetString(language.Russian, "Error occured while tag parsing `%s`: %s", "Произошла ошибка при разборе тега `%s`: %s")
	message.SetString(language.Russian, "Fail to extract Field tag", "Не удалось извлечь тег поля")
	message.SetString(language.Russian, "Error occured while creating a style: %w", "Произошла ошибка при создании стиля: %w")
	message.SetString(language.Russian, "Error while setting cell value: %v", "Ошибка при установке значения ячейки: %v")
	message.SetString(language.Russian, "Error while setting cell style: %v", "Ошибка при установке стиля ячейки: %v")
	message.SetString(language.Russian, "Operation cancelled", "Операция отменена")
	message.SetString(language.Russian, "error occurred while creating a style: %w", "Ошибка при создании стиля: %w")
	message.SetString(language.Russian, "Failed to close file: %s", "Не удалось закрыть файл: %s")
	message.SetString(language.Russian, "Error occurred while tag parsing `%s`: %s", "Ошибка при разборе тэга: `%s`: %s")
	message.SetString(language.Russian, "Sheet `%s` not found", "Лист `%s` не найден")
}
