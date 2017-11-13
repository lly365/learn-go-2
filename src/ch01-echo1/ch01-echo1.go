//ch01-echo1 打印它的命令行参数

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //被隐式地赋予零值

	for i, j := 1, len(os.Args); i < j; i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

//usage: ./ch01-echo1 --name=nemo -a 33 -city longnan
