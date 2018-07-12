package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

	xlsx := excelize.NewFile()

	xlsx.SetCellHyperLink("Sheet1", "A3", "http://www.antlinker.com", "External")
	// Set underline and font color style for the cell.
	style, _ := xlsx.NewStyle(`{"font":{"color":"#1265BE","underline":"single"}}`)
	xlsx.SetCellStyle("Sheet1", "A3", "A3", style)

	xlsx.SetCellValue("Sheet1", "A3", "集结号")

	xlsx.SetColWidth("Sheet1", "A", "H", 20)

	format, err := xlsx.NewConditionalStyle(`{"font":{"color":"#9A0511"},"fill":{"type":"pattern","color":["#FEC7CE"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(format)
	xlsx.SetConditionalFormat("Sheet1", "A1:A10", fmt.Sprintf(`[{"type":"cell","criteria":">","format":%d,"value":"6"}]`, format))

	// Rose format for bad conditional.
	format1, err := xlsx.NewConditionalStyle(`{"font":{"color":"#9A0511"},"fill":{"type":"pattern","color":["#FEC7CE"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}
	xlsx.SetCellStyle("Sheet1", "B3", "B3", format1)

	xlsx.SetCellValue("Sheet1", "B3", "集结号")

	xlsx.SetCellValue("Sheet1", "C5", "集结号")
	// Light yellow format for neutral conditional.
	// format2, err := xlsx.NewConditionalStyle(`{"font":{"color":"#9B5713"},"fill":{"type":"pattern","color":["#FEEAA0"],"pattern":1}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// xlsx.SetCellStyle("Sheet1", "C5", "C5", format2)

	// Light green format for good conditional.
	format3, err := xlsx.NewConditionalStyle(`{"font":{"color":"#09600B"},"fill":{"type":"pattern","color":["#C7EECF"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}
	xlsx.SetCellStyle("Sheet1", "D5", "D5", format3)

	xlsx.SetCellValue("Sheet1", "D5", "集结号")

	xlsx.MergeCell("Sheet1", "A3", "A4")
	err = xlsx.SaveAs("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
