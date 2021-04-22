package myloger

import (
	"fmt"
	"runtime")


func GetInfo(number int) (string,int,string){
	// pc = 执行runtime.Caller的函数
	// file = 执行runtime.Caller函数的文件
	// line = 调用执行runtime.Caller函数的行号 [0,1,2,..] 0代表当前runtime.Caller 函数行号 1代表向上找一层也就是f()的行号 以此类推
	// ok = 判断执行runtime.Caller函数的成功
	pc, file, line, ok := runtime.Caller(number)
	if !ok {
		fmt.Printf("runtime.Caller() failed \n")
	}

	//获取执行函数的名称
	funcName := runtime.FuncForPC(pc).Name()
	return file, line, funcName
}
