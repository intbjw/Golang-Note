//打印出它的命令行参数
//按照惯例我们在每个包声明前添加注释
package main

import (
	"fmt"
	"os"
)

func main() {
	//var定义了两个string类型变量s和sep。变量会在声明时直接初始化
	//声明把s和sep隐式地初始化成空字符串
	var s, sep string
	//Go语言中只有for循环一种语句。for循环有多种形式，下面这种是最常见的一种形式
	//短变量声明、判断条件、循环结束后执行的语句
	//for循环这三部分都可以省略
	for i := 1; i < len(os.Args); i++ {
		//对于string类型，+运算符连接字符串
		// += 和C++中一样
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
