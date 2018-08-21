// 回调函数
package main

import (
	"fmt"
)

//type 回调函数的类型  func(参数)返回值
type myCallBackFunc func(value1, value2 int) int //回调函数没有具体实现的函数体，也就是函数的多态性。

//这样的好处就是可扩展性好。

func add(a, b int) int {
	return a + b

}

func Calc(a, b int, callback myCallBackFunc) (result int) {
	result = callback(a, b)
	return

}

func main() {
	fmt.Println(Calc(2, 3, add))
}
