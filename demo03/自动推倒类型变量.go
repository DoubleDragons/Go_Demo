// 自动推倒类型变量
package main

import (
	"fmt"
)

func main() {

	m03 := 100
	//%T就是查看变量的类型，\n是换行
	fmt.Printf("变量m03所属的类型为： %T\n", m03)

	m04 := "我是字符串"
	fmt.Printf("变量m04所属的类型为：%T\n", m04)

	m05 := 0.006
	fmt.Printf("变量m05所属的类型：%T\n", m05)

	m06 := 1.0
	fmt.Printf("变量m06所属的类型我：%T\n", m06)
	//如果要用格式化输出，用println是不行的，
	var a, b, c int = 1, 2, 3
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)

	//变量的交换
	i, j := 10, 20
	//	可以直接写
	j, i = 10, 20
	fmt.Printf("i=%d,j=%d\n", i, j)

	//匿名变量，基本上配合多个返回值的函数时候使用。
	k, l := 100, 200
	l, _ = k, l
	fmt.Printf("l=%d\n", l)
}
