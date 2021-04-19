package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)



func Fcreater(f string) *os.File {
	f1, err := os.Create(f)
	if err != nil {
		fmt.Println("this crete err now")
	}
	fmt.Println("this creted now")
	return f1
}

func Fopener(f string) *os.File {
	f1, err := os.Open(f)
	if err != nil {
		fmt.Println("this crete err now")
	}
	fmt.Println("this Open now")
	return f1
}

func Freader(r *os.File) {
	bf := bufio.NewReader(r)
	for {
		data, err := bf.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完成")
		} else if err != nil {
			fmt.Println("this reader err now")
			return
		}
		fmt.Println(data)
		return
	}

}

func Fwriter(w *os.File, str string) {
	n, err := w.Write([]byte(str))
	if err == io.EOF {
		fmt.Println("写入完成")
	} else if err != nil {
		fmt.Println("this writer err now")
		return
	}
	fmt.Printf("写入字节数：%d，写入完成\n", n)
}


// 下面是使用指导

// func main() {
// 	filename := ".\\a.txt"
// 	f := Fcreater(filename)
// 	Fwriter(f,"my name is zhangsan")
// 	f.Close()
// }
