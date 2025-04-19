package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	sortingAddon()
}

func sortingAddon() {
	message.SetString(language.Russian, "Field %s not found", "Поле %s не найдено")
	message.SetString(language.Russian, "Field `%s` can only be sorted by ASC or DESC", "Поле `%s` может быть отсортировано только по ASC или DESC")
	message.SetString(language.Russian, "Failed to parse tag for `%s` failed: %s", "Не удалось разобрать тег для `%s`: %s")
	message.SetString(language.Russian, "Field `%s` doesn't have ui tag", "Поле `%s` не имеет ui тега")
	message.SetString(language.Russian, "Field `%s` is array and cannot be used for sorting", "Поле `%s` является массивом и не может быть использовано для сортировки")
	message.SetString(language.Russian, "Too complex fieldPath for structure `%s`", "Слишком сложный путь поля для структуры `%s`")
	message.SetString(language.Russian, "Invalid field path for `%s` field", "Недопустимый путь поля для поля `%s`")
}
