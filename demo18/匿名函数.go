// 匿名函数
package main

import (
	"fmt"
)

func main() {
	//有参数有返回值的1

	f1 := func(value1, value2 int) int {
		return value1 + value2
	}
	f1(1, 2)

	//有参数有返回值的2
	x := func(value1, value2 int) int {
		return value1 + value2
	}(3, 4)

	fmt.Println(x)
}
