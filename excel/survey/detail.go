package main

import "github.com/360EntSecGroup-Skylar/excelize"
import "fmt"

func getDepart() (map[string]string, map[string]interface{}) {
	dept := make(map[string]string)
	dept["A"] = "校区"
	dept["B"] = "学院"
	dept["C"] = "专业"
	dept["D"] = "年级"
	dept["E"] = "班级"
	dept["F"] = "需参与人数"
	dept["G"] = "已参与人数"

	value := make(map[string]interface{})
	value["A"] = "浮山校区"
	value["B"] = "法学院"
	value["C"] = "法学"
	value["D"] = "2016"
	value["E"] = "16法学01"
	value["F"] = 25
	value["G"] = 20

	return dept, value
}

func export() {
	xlsx := excelize.NewFile()

	categories := make(map[string]interface{})
	dept, value := getDepart()
	categories["A1"] = dept["A"]
	categories["B1"] = dept["B"]
	categories["C1"] = dept["C"]
	categories["D1"] = dept["D"]
	categories["E1"] = dept["E"]
	categories["F1"] = dept["F"]
	categories["G1"] = dept["G"]
	categories["A2"] = value["A"]
	categories["B2"] = value["B"]
	categories["C2"] = value["C"]
	categories["D2"] = value["D"]
	categories["E2"] = value["E"]
	categories["F2"] = value["F"]
	categories["G2"] = value["G"]

	for k, v := range categories {
		xlsx.SetCellValue("Sheet1", k, v)
	}

	quest := "说句问候语。"
	note := []string{
		"你好",
		"你好kasdjfaksdjfadsjfasjdl",
		"你好kadsjfasdjfkadsjfakdslj",
		"你好jfadskfjalsdjfasdl;",
		"你好jfdkasl;jfklasdfj;adklsfjadsklfjasd;lfjadsklfhdsjfhdsf;lkjsaddfj;a",
		"你好h",
		"你好.",
	}

	xlsx.SetCellValue("Sheet1", "A3", quest)
	xlsx.MergeCell("Sheet1", "A3", "G3")

	notes := make(map[string]string)
	// notes["A4"] = note[0]
	// notes["A5"] = note[1]
	// notes["A6"] = note[2]
	// notes["A7"] = note[3]
	// notes["A8"] = note[4]
	// notes["A9"] = note[5]
	// notes["A10"] = note[6]
	ncol := 65
	nrow := 4
	for _, v := range note {
		n := fmt.Sprintf("%s%d", u2s(ncol), nrow)
		notes[n] = v
		nrow++
	}

	for k, v := range notes {
		xlsx.SetCellValue("Sheet1", k, v)
		xlsx.MergeCell("Sheet1", k, "G"+k[1:])
	}
	// 设置边框 border
	// 设置对齐方式 alignment
	style, err := xlsx.NewStyle(`{"border":[{"type":"left","color":"0d5330","style":1},{"type":"top","color":"0d5330","style":1},{"type":"bottom","color":"0d5330","style":1},{"type":"right","color":"0d5330","style":1}],"alignment":{"horizontal":"center","ident":1,"justify_last_line":true,"reading_order":0,"relative_indent":1,"shrink_to_fit":true,"text_rotation":0,"vertical":"","wrap_text":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	xlsx.SetCellStyle("Sheet1", "A1", "G10", style)

	names := map[string]string{"I1": "明细表", "I2": "已做", "I3": "未做", "J1": "法学院"}
	values := map[string]interface{}{"J2": value["G"], "J3": 5}

	for k, v := range names {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Sheet1", k, v)
	}

	// 增加表格
	xlsx.AddTable("Sheet1", "I1", "J3", `{"table_style":"TableStyleMedium2", "show_first_column":true,"show_last_column":true,"show_row_stripes":true,"show_column_stripes":true}`)
	// barStacked
	xlsx.AddChart("Sheet1", "I10", `{"type":"barStacked","series":[{"name":"=Sheet1!$I$2","categories":"=Sheet1!$J$1:$J$1","values":"=Sheet1!$J$2:$J$2"},{"name":"=Sheet1!$I$3","categories":"=Sheet1!$J$1:$J$1","values":"=Sheet1!$J$3:$J$3"}],"title":{"name":"各学院分布"}}`)

	if err := xlsx.SaveAs("./detail.xlsx"); err != nil {
		fmt.Println(err.Error())
	}

}

func main() {
	export()
}

func u2s(i int) (s string) {
	if i < 65 || i > 90 {
		return
	}
	s = string(rune(i))
	return
}
