// 类型别名
package main

import (
	"fmt"
)

func main() {
	//type 别名 类型
	type 整形 int64
	var c 整形

	fmt.Printf("c=%T", c) //输出结果为：c=main.a
}
