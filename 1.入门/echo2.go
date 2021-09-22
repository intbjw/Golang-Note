package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 每次循环迭代，range产生一对值；索引以及在该索引处的元素值、
	// Go语言中不允许使用无用的局部变量
	// Go语言中的解决方法是用空标识符，也就是下划线_
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "

	}
	fmt.Println(s)
}
