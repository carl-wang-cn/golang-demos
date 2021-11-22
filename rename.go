package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	path := "/Users/carl/Movies/突围"
	listFiles(path)
}
func listFiles(dirname string) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fileInfos {
		filename := dirname + "/" + f.Name() //记录当前文件夹下的文件名

		// os.Rename(filename, dirname+"/怒晴湘西"+f.Name())
		//如果文件名包含如下字段，则将其文件名更改
		replace_str := "[www.domp4.cc]"
		if strings.Contains(f.Name(), replace_str) {
			os.Rename(filename, dirname+"/"+strings.Replace(f.Name(), replace_str, "", -1))
		}
		fmt.Println(filename) //打印文件地址
		if f.IsDir() {        //判断是否是文件夹 如果是文件夹则继续递归调用
			listFiles(filename)
		}
	}
}
