package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 创建文件函数，返回两个值一个 文件句柄，一个error 错误
func Fcreater(f string) (*os.File, error) {
	f1, err := os.Create(f)
	if err != nil {
		fmt.Println("this crete err now")
		return nil, err
	}
	return f1,nil	
}

// 打开文件函数，返回两个值一个 文件句柄，一个error 错误
func Fopener(f string) (*os.File, error) {
	f1, err := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0677)
	if err != nil {
		fmt.Println("this crete err now")
		return nil,err
	}
	return f1,nil
}

// 读取文件函数,返回一个 string 类型的值
func Freader(f string) (string, bool) {
	r, err := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0677)
	if err != nil {
		fmt.Println("this crete err now")
		return "",false
	}
	bf := bufio.NewReader(r)
	for {
		data, err := bf.ReadString('\n')
		if err == io.EOF {
			return data,true
		} else if err != nil {
			fmt.Println("this reader err now")
			return "",false
		}
	}
}

// 写入文件函数，返回一个bool值，通过bool判断成没成功
func Fwriter(filename, str string) bool{
	f ,err := os.OpenFile(filename,os.O_CREATE|os.O_RDWR|os.O_APPEND,0777)
	if err != nil {
		fmt.Println("打开错误")
	}
	_, err = f.Write([]byte("\n"+str))
	if err == io.EOF {
		// fmt.Println("file writed OK")
		return true
	} else if err != nil {
		fmt.Println("this writer err now")
		return false
	}
	return true
}
