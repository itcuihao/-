package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

	xlsx := excelize.NewFile()
	xlsx.AddTable("Sheet1", "A1", "D5", ``)

	xlsx.AddTable("Sheet1", "F2", "H6", `{"table_style":"TableStyleMedium2", "show_first_column":true,"show_last_column":true,"show_row_stripes":true,"show_column_stripes":true}`)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
