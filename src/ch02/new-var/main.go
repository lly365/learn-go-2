package main

import (
	"fmt"
)

func main() {
	p := new(int)   //p --> *int 类型，指向匿名的 int 变量
	fmt.Println(*p) // 0
	*p = 2
	fmt.Println(*p)

	var s *string = new(string)
	fmt.Println(*s)
	*s = "你好，世界"
	fmt.Println(*s)

	//每次new都返回新的地址，所以下面两个*int不相等
	p1 := newInt()
	p2 := newInt()
	fmt.Println(p1 == p2, *p1 == *p2)
}

func newInt() *int {
	return new(int)
}
