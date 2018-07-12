package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	categories := map[string]string{"A2": "已做", "A3": "未做", "B1": "法学院", "C1": "经济学院", "D1": "历史学院"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 5, "B3": 5, "C3": 2, "D3": 1}
	xlsx := excelize.NewFile()
	for k, v := range categories {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	// bar3D
	xlsx.AddChart("Sheet1", "E1", `{"type":"bar3D","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"考试成绩各学院分布"}}`)

	// bar
	xlsx.AddChart("Sheet1", "E20", `{"type":"bar","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"考试成绩各学院分布"}}`)

	// barStacked
	xlsx.AddChart("Sheet1", "E40", `{"type":"barStacked","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"hahah"}}`)

	// doughnut
	xlsx.AddChart("Sheet1", "E60", `{"type":"doughnut","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"lalalal"}}`)
	// line
	xlsx.AddChart("Sheet1", "E80", `{"type":"line","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"考试成绩各学院分布"}}`)
	// pie
	// xlsx.AddChart("Sheet1", "E100", `{"type":"pie","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"考试成绩各学院分布"}}`)
	// // pie3D
	// xlsx.AddChart("Sheet1", "E120", `{"type":"pie3D","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"考试成绩各学院分布"}}`)
	// // radar
	xlsx.AddChart("Sheet1", "E140", `{"type":"radar","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"考试成绩各学院分布"}}`)
	// // scatter
	xlsx.AddChart("Sheet1", "E160", `{"type":"scatter","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$4","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"考试成绩各学院分布"}}`)

	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
