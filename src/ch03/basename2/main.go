package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") //如果没有找到，返回 -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
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
