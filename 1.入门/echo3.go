package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 可以使用Join函数
	fmt.Println(strings.Join(os.Args[1:], " "))
}
