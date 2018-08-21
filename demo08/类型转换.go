// 类型转换
package main

import (
	"fmt"
)

func main() {

	/*
		布尔类型不能转换为整形，整形也不能转换为布尔类型

	*/
	var a byte = 'a'
	var b int
	b = int(a)
	fmt.Printf("a=%v", a)
	fmt.Printf("b=%d", b)
}
