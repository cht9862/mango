package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 方法1 使用ioutil 编写CopyFile函数
func CopyFileIOutil(src string, dst string) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("读取文件失败", err)
	}

	err = ioutil.WriteFile(dst, data, 0777)
	if err != nil {
		fmt.Println("写入文件失败", err)
	}
	fmt.Println("CopyFile 完成，程序退出")
}

// 方法2 使用os 编写CopyFile函数
func CopyFileOS(src string, dst string) {
	f, err := os.Open(src)
	if err != nil {
		fmt.Printf("打开%v文件失败", src)
		return
	}
	var data []byte = make([]byte, 1024)
	n, err := f.Read(data)
	if err == io.EOF {
	} else if err != nil {
		fmt.Println("读取文件失败")
	}
	f.Close()

	f, err = os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("打开%v文件失败", dst)
	}
	_, err = f.Write(data[:n])
	if err == io.EOF {
	} else if err != nil {
		fmt.Println("写入文件失败")
	}
	f.Close()
}

// 方法3 使用os 编写CopyFile函数
func CopyFileBufIO(src string, dst string) {
	var zhangsan *string
	var wangwu []string = make([]string, 0, 19)
	f, err := os.Open(src)
	if err != nil {
		fmt.Printf("打开%v文件失败", src)
		return
	}
	bf := bufio.NewReader(f)
	for {
		data, err := bf.ReadString('\n')
		zhangsan = &data
		wangwu = append(wangwu, *zhangsan)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("写入文件失败")
		}
	}
	f.Close()

	f, err = os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("打开%v文件失败", dst)
	}
	bfw := bufio.NewWriter(f)
	for _, v := range wangwu {
		_, err = bfw.WriteString(v)
		if err != nil {
			fmt.Printf("写入%v文件失败", dst)
		}
	}
	bfw.Flush()
	f.Close()
}

// 演示操作
// func main() {
// 	CopyFileIOutil("a.txt","b.txt")
// 	CopyFileBufIO("a.txt","c.txt")
// 	CopyFileOS("a.txt","d.txt")

// }
