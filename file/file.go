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
		return err
	}
	return f1
}
// 打开文件函数，返回两个值一个 文件句柄，一个error 错误
func Fopener(f string) (*os.File, error) {
	f1, err := os.Open(f)
	if err != nil {
		fmt.Println("this crete err now")
		return err
	}
	return f1
}

// 读取文件函数,返回一个 string 类型的值
func Freader(r *os.File) (string, bool) {
	bf := bufio.NewReader(r)
	for {
		data, err := bf.ReadString('\n')
		if err == io.EOF {
			return "", true
		} else if err != nil {
			fmt.Println("this reader err now")
			return false
		}
		return data
	}

}

// 写入文件函数，返回一个bool值，通过bool判断成没成功
func Fwriter(w *os.File, str string) bool{
	n, err := w.Write([]byte(str))
	if err == io.EOF {
		// fmt.Println("file writed OK")
		return true
	} else if err != nil {
		fmt.Println("this writer err now")
		return false
	}
}

