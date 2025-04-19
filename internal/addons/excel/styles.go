package excel

import "github.com/xuri/excelize/v2"

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

func LoadDataStyle(file *excelize.File) (int, error) {
	dataStyle := excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			VertAlign: "center",
		},
	}
	return GetStyleId(file, &dataStyle)
}
