package main

import (
	// "fmt"
	// _ "fmt"
	"fmt"
	// "path"
	"project1/myloger"
	"time"
)



func main() {
	// newlog := myloger.NewLog()
	newfilelog := myloger.NewFileLog()
	// newfilelog.DebugBool = true
	var a chan int = make(chan int, 18)
	a <- 13
	v, ok := <-a
	if !ok {
		fmt.Println("失败")
	}
	fmt.Println(v)
	for {
		newfilelog.Debug()
		// newfilelog.Trace("【Trace】这里有错误")
		// newfilelog.Info("【Info】这里有错误")
		// newfilelog.Warning("【Warning】这里有错误")
		time.Sleep(time.Second * 2)
	}
	// myloger.GetInfo(2)
	// // Split函数将路径从最后一个斜杠后面位置分隔为两个部分（dir和file）并返回。如果路径中没有斜杠，函数返回值dir会设为空字符串，file会设为path。
	// // 两个返回值满足path == dir+file。
	// res1,res2 := path.Split("F:/gowork/src/irun/irun.go")
	// fmt.Println(res1)
	// fmt.Println(res2) 

	// // Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠。结果是经过简化的，所有的空字符串元素会被忽略。
	// res3 := path.Join("F:/gowork/src/irun","iruna","irun.og")
	// fmt.Println(res3)

	// // Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠。如果路径是空字符串，会返回"."；如果路径由1到多个斜杠后跟0到多个非斜杠字符组成，会返回"/"；
	// // 其他任何情况下都不会返回以斜杠结尾的路径。
	// res4 := path.Dir("F:/gowork/src/irun/irun.go")
	// fmt.Println(res4)

	// // Base函数返回路径的最后一个元素。在提取元素前会求掉末尾的斜杠。如果路径是""，会返回"."；如果路径是只有一个斜杆构成，会返回"/"。
	// res5 := path.Base("F:/gowork/src/irun/irun.go")
	// fmt.Println(res5)
}
