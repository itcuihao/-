package main

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

const commitTmplHTML = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width,initial-scale=1,user-scalable=0">
    <title></title>
</head>
<body>
<div style="background-color: #ffffff;">
    <div style="padding: 15px 15px 0px;margin: 0px;font-size: 16px;color: #3e3e3e;text-align: center;line-height: 25px;overflow: hidden;">
        {{.Title}}
    </div>
    <div style="padding: 15px;line-height: 25px;font-size: 14px;color: #666666;text-align: left;">
        {{.Content}}
    </div>
    <div style="margin: 5px 30px 0px;font-size: 14px;color: #666666;line-height: 30px;text-align: right;">
		{{if .IsWriteTime}}
		<div style="margin-bottom: 20px;">{{.WriteTime}}</div>
		{{end}}
        <div>
            <span style="line-height: 40px;vertical-align: middle;">承诺人：</span>
            <span style="width: 140px; display: inline-block;text-align: left;"><img src="{{.Sign}}"
                 width="110" height="40" style="vertical-align: middle"></span>
        </div>
        <div>
            <span style="margin-right: 5px;">时&nbsp;&nbsp;&nbsp;&nbsp;间：</span><span
                style="width: 140px; display: inline-block;text-align: left;">{{.SignTime}}</span>
        </div>
    </div>
</div>
</body>
</html>
`

var commitTmpl *template.Template

// 解析html模板
func tmp() {
	tcommitTmpl, err := template.New("").Parse(commitTmplHTML)
	if err != nil {
		panic(err)
	}
	commitTmpl = tcommitTmpl
}
func main() {
	data := struct {
		Title       string
		Content     template.HTML
		IsWriteTime bool
		WriteTime   string
		Sign        string
		SignTime    string
	}{
		Title:       "Hello World",
		Content:     template.HTML("你好"),
		IsWriteTime: true,
		WriteTime:   time.Now().String(),
		Sign:        "",
		SignTime:    time.Now().String(),
	}

	tmp()
	htmlOut := new(bytes.Buffer)
	err := commitTmpl.Execute(htmlOut, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(htmlOut.String())
}
