package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func dealFunc(path string, fi os.FileInfo, err error) error {
	if nil == fi {
		return err
	}
	if fi.IsDir() {
		return nil
	}
	name := fi.Name()
	//fmt.Println(name)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(fmt.Sprintf("打开文件%s出错%s！", name, err))
		return nil
	}
	//  defer file.Close()
	buffer := new(bytes.Buffer)
	io.Copy(buffer, file)

	file.Close()

	// linux path
	frompath := fmt.Sprintf("./%s", path)
	// windows path
	//  frompath := fmt.Sprintf(".\\%s",path)
	topath := strings.Replace(frompath, "icon", "iconbak", 1)
	fmt.Println(frompath)
	//fmt.Println(topath)
	if err := os.Rename(frompath, topath); err != nil {
		fmt.Println(fmt.Sprintf("%s文件搬移出错！", frompath))
		return nil
	}

	return nil
}

func main() {
	filepath.Walk("./icon", dealFunc)
}
