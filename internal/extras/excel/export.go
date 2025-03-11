package excel

import (
	"app/internal/dialogs"
	"errors"
	"fmt"
	"github.com/kuzgoga/fogg"
	"github.com/xuri/excelize/v2"
	"log/slog"
	"reflect"
	"slices"
	"strings"
	"time"
)

type TableHeaders struct {
	Headers              []string
	IgnoredFieldsIndexes []int
}

func isPrimitiveType(valueType reflect.Type) bool {
	switch valueType.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.String:
		return true
	default:
		return false
	}
}

func DeleteDefaultSheet(file *excelize.File) error {
	sheetId, err := file.GetSheetIndex("Sheet1")
	if err != nil {
		return err
	}
	if sheetId != -1 {
		if err := file.DeleteSheet("Sheet1"); err != nil {
			return err
		}
	}
	return nil
}

func ExportEntitiesToSpreadsheet(filename string, exporters ...ExporterInterface) error {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			slog.Error("Error while closing excel file: ", err)
		}
	}()

	for _, exporter := range exporters {
		err := ExportEntityToSpreadsheet(file, exporter.GetSheetName(), exporter.GetEntity(), exporter.GetProvider())
		if err != nil {
			return err
		}
	}

	if err := DeleteDefaultSheet(file); err != nil {
		return err
	}

	err := WriteData(file, filename)
	if err != nil {
		return err
	}
	return nil
}

func ExportEntityToSpreadsheet[T any](file *excelize.File, sheetName string, entity T, provider func() ([]any, error)) error {
	_, err := file.NewSheet(sheetName)
	if err != nil {
		return err
	}

	headers, err := WriteHeaders(sheetName, entity, file)
	if err != nil {
		return err
	}

	items, err := provider()
	if err != nil {
		return err
	}

	// TODO: appearance
	for i, item := range items {
		structValue := reflect.ValueOf(item).Elem()

		columnOffset := 0
		for j := 0; j < structValue.NumField(); j++ {
			if slices.Contains(headers.IgnoredFieldsIndexes, j) {
				columnOffset--
				continue
			}

			field := structValue.Field(j)
			tagLiteral := string(structValue.Type().Field(j).Tag)
			tag, err := fogg.Parse(tagLiteral)

			if err != nil {
				return err
			}

			if isPrimitiveType(field.Type()) {
				fieldValue := field.Interface()
				cellName, err := GetCellNameByIndices(j+columnOffset, i+1)
				if err != nil {
					return err
				}

				datatype := tag.GetTag("ui").GetParamOr("datatype", "")

				if datatype == timestampTag {
					err = file.SetCellValue(sheetName, cellName, time.Unix(fieldValue.(int64), 0))
				} else {
					err = file.SetCellValue(sheetName, cellName, fieldValue)
				}

				slog.Info(fmt.Sprintf("Field %s value: %v, %s\n", cellName, fieldValue, datatype))

				if err != nil {
					return err
				}
			} else {
				cellName, err := GetCellNameByIndices(j+columnOffset, i+1)
				if err != nil {
					return err
				}

				structInfo := structValue.Type().Field(j)
				structVal := structValue.Field(j)

				value, err := SerializeNestedField(structInfo, structVal)
				if err != nil {
					return err
				}

				err = file.SetCellValue(sheetName, cellName, *value)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func getNestedFieldValue(obj reflect.Value, path string) (any, error) {
	fields := strings.Split(path, ".")
	val := obj

	for _, field := range fields {
		if field == "" {
			continue
		}

		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		val = val.FieldByName(field)

		if !val.IsValid() {
			return nil, fmt.Errorf("field %s not found", field)
		}
	}
	return val.Interface(), nil
}

func SerializeNestedField(fieldInfo reflect.StructField, fieldValue reflect.Value) (*string, error) {
	tag, err := fogg.Parse(string(fieldInfo.Tag))
	if err != nil {
		return nil, err
	}

	var uiTag fogg.Tag

	if tag.HasTag("ui") && tag.GetTag("ui").HasParam("field") {
		uiTag = *tag.GetTag("ui")
	} else {
		slog.Error("Fail to extract Field tag")
		return nil, nil
	}

	if uiTag.HasParam("field") {
		var result string
		fieldPath := uiTag.GetParam("field").Value
		if fieldInfo.Type.Kind() == reflect.Slice {
			var items = make([]string, fieldValue.Len())
			for i := 0; i < fieldValue.Len(); i++ {
				item := fieldValue.Index(i)
				nestedFieldValue, err := getNestedFieldValue(item, fieldPath)
				if err != nil {
					return nil, err
				}
				items[i] = fmt.Sprintf("%v", nestedFieldValue)
			}
			result = strings.Join(items, ", ")
		} else {
			nestedFieldValue, err := getNestedFieldValue(fieldValue, fieldPath)
			if err != nil {
				return nil, err
			}
			result = fmt.Sprintf("%v", nestedFieldValue)
		}
		return &result, nil
	} else {
		return nil, nil
	}
}

func GetHeaderCellNameByIndex(column int) (string, error) {
	colName, err := excelize.ColumnNumberToName(column + 1)
	if err != nil {
		return "", err
	}
	cellName, err := excelize.JoinCellName(colName, 1)
	if err != nil {
		return "", err
	}
	return cellName, nil
}

func GetCellNameByIndices(column int, row int) (string, error) {
	colName, err := excelize.ColumnNumberToName(column + 1)
	if err != nil {
		return "", err
	}
	cellName, err := excelize.JoinCellName(colName, row+1)
	if err != nil {
		return "", err
	}
	return cellName, nil
}

func ExportHeaders(entity any) (TableHeaders, error) {
	headers := TableHeaders{}
	v := reflect.TypeOf(entity)
	for i := range v.NumField() {
		tag, err := fogg.Parse(string(v.Field(i).Tag))
		if err != nil {
			return headers, errors.New(fmt.Sprintf("Error occured while tag parsing `%s`: %s", err, string(v.Field(i).Tag)))
		}

		uiTag := tag.GetTag("ui")
		if uiTag == nil {
			headers.IgnoredFieldsIndexes = append(headers.IgnoredFieldsIndexes, i)
			continue
		}

		label := uiTag.GetParamOr("label", uiTag.GetParamOr(excelNameTag, ""))

		if label != "" {
			headers.Headers = append(headers.Headers, label)
		} else {
			headers.IgnoredFieldsIndexes = append(headers.IgnoredFieldsIndexes, i)
		}
	}
	return headers, nil
}

func WriteHeaders(sheetName string, entity any, file *excelize.File) (TableHeaders, error) {
	headers, err := ExportHeaders(entity)
	if err != nil {
		return headers, err
	}

	for i, header := range headers.Headers {
		cellName, err := GetHeaderCellNameByIndex(i)
		if err != nil {
			return headers, err
		}

		err = file.SetCellValue(sheetName, cellName, header)
		if err != nil {
			return headers, err
		}
	}

	err = ApplyStyleHeaders(file, sheetName, headers)
	if err != nil {
		return headers, err
	}

	return headers, nil
}

func GetStyleId(f *excelize.File, style *excelize.Style) (int, error) {
	styleId, err := f.NewStyle(style)
	if err != nil {
		return 0, fmt.Errorf("error occured while creating a style: %w", err)
	}

	return styleId, nil
}

func LoadHeadersStyle(file *excelize.File) (int, error) {
	headersStyle := excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "top",
				Color: "000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      true,
			VertAlign: "center",
		},
	}
	return GetStyleId(file, &headersStyle)
}

func ApplyStyleHeaders(file *excelize.File, sheetName string, headers TableHeaders) error {
	styleId, err := LoadHeadersStyle(file)
	if err != nil {
		return err
	}

	cellName, err := GetHeaderCellNameByIndex(len(headers.Headers) - 1)
	if err != nil {
		return err
	}

	err = file.SetCellStyle(sheetName, "A1", cellName, styleId)
	if err != nil {
		return err
	}

	return nil
}

func WriteData(file *excelize.File, filename string) error {
	filepath, err := dialogs.SaveFileDialog("Экспорт данных", filename)

	if !strings.HasSuffix(filepath, ".xlsx") {
		filepath += ".xlsx"
	}

	if err != nil {
		return err
	}

	if filepath == "" {
		dialogs.InfoDialog("Экспорт данных", "Операция отменена")
	}

	if err := file.SaveAs(filepath); err != nil {
		return err
	}
	return nil
}
