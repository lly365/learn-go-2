package main

import (
	"fmt"
)

//移除目录名和后缀，返回基本的文件名
//比如：a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	s := "a"
	fmt.Println(basename(s))

	s = "a.go"
	fmt.Println(basename(s))

	s = "a/b/c.go"
	fmt.Println(basename(s))

	s = "a/b.c.go"
	fmt.Println(basename(s))
}
