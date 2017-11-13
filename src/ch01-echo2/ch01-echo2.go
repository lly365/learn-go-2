//ch01-echo2 打印它的命令行参数

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

//./ch01-echo2 --name=nemo -a 33 -city longnan
