package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	s, e int
)

func main() {
	ReadLine("./1.log", Parsing)
}

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func Parsing(str string) {
	//	if strings.HasPrefix(str, "D") && strings.Contains(str, "描述:回收数据-获取分片") && strings.Contains(str, "返回数据:") {
	if strings.HasPrefix(str, "D") && strings.Contains(str, "描述:回收数据-获取收录记录") && strings.Contains(str, "返回数据:") {
		var str3 string
		var data map[string]interface{}
		var record []interface{}
		strbyte := []byte(str)
		for k, v := range strbyte {
			if string(v) == "{" {
				s = k
				break
			}
		}
		for k, v := range strbyte {
			if string(v) == "$" {
				e = k
				break
			}
		}
		str2 := str[s:e]

		if strings.HasPrefix(str2, "\"") {
			str3 = "{" + str2
		} else if strings.HasPrefix(str2, ":") {
			str3 = "{" + str2
		} else {
			str3 = str2
		}
		if err := json.Unmarshal([]byte(str3), &data); err != nil {
			fmt.Println(err)
		}
		record = data["result"].([]interface{})
		for _, v := range record {
			row := v.(map[string]interface{})
			url := row["url"].(string)
			f, err := os.OpenFile("./huoqujilu.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			_, err1 := f.Write([]byte(url + "\n"))
			if err1 != nil {
				return
			}
		}
	}
}
