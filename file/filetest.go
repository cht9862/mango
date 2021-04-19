package main

import (
	"fmt"

	"github.com/cht9862/mango/file"
)


func main() {
	filename := ".\\a.txt"
	//  打开文件
	f, err := file.Fopener(filename)
	if err != nil {
		fmt.Println("this open err ->",err)
		return
	}
	defer f.Close()
  
	// 读取文件
	data, tr := file.Freader(f)
	if !tr {
		fmt.Println("获取内容失败")
	}
	fmt.Println(data)

	// 写入数据到文件
	tr = file.Fwriter(filename,"zhaang~~~~")
	if !tr {
		fmt.Println("写入失败")
	} else {
		fmt.Println("写入成功")
	}
	
}
