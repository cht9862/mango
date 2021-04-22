package myloger

import (
	"fmt"
	"log"
	_ "time"

	"github.com/cht9862/mango/file"
	
)

// 向文件中输出日志
type FileLog struct {
	DebugBool bool
}

// 程序被调用主函数，返回一个FileLog结构体
func NewFileLog() FileLog {
	return FileLog{}
}

func logline() string {
	file, line, funcName := GetInfo(3)
	str := fmt.Sprintf("路径：%s 报错函数名：%s 行号：%d",file, funcName, line)
	return str
}

// Debug级别错误
func (fi *FileLog) Debug() {
	if !fi.DebugBool {
		return
	}
	f, err := file.Fopener("a.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	log.SetOutput(f)
	res := logline()
	log.Println("[DeBug]",res)
	f.Close()
}

// Trace级别错误
func (fi *FileLog) Trace(msg string) {
	if !fi.DebugBool {
		return
	}
	f, err := file.Fopener("a.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	log.SetOutput(f)
	res := logline()
	log.Println("[Trace]",res)
	f.Close()
}

// Info级别错误
func (fi *FileLog) Info(msg string) {
	f, err := file.Fopener("a.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	log.SetOutput(f)
	res := logline()
	log.Println("[Info]",res)
	f.Close()
}

// Warning级别错误
func (fi *FileLog) Warning(msg string) {
	f, err := file.Fopener("a.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	log.SetOutput(f)
	res := logline()
	log.Println("[Warning]",res)
	f.Close()
}

// Error级别错误
func (fi *FileLog) Error(msg string) {
	f, err := file.Fopener("a.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	log.SetOutput(f)
	res := logline()
	log.Println("[Error]",res)
	f.Close()
}

// Fatal级别错误
func (fi *FileLog) Fatal(msg string) {
	f, err := file.Fopener("a.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	log.SetOutput(f)
	res := logline()
	log.Println("[Fatal]",res)
	f.Close()
}
