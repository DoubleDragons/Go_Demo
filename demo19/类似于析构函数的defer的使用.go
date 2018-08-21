// 类似于析构函数的defer的使用
//一个函数结束前的调用。
package main

import (
	"fmt"
)

func test() {

	defer fmt.Println("我是test结束前的调用")
	fmt.Println("我是test函数的函数体")
}
func closefunc() {

	fmt.Println("我是main函数结束前的调用")

}

func main() {
	//	test()
	//	defer closefunc()
	//	fmt.Println("我是main函数的函数体")

	//下面可以看到一个迭代函数是重新建立一个地址跟函数体里面的I是不一样的地址。
	i := 0
	fmt.Println("执行循环之前i = ", &i)
	str := "sfdsdfsdfds"
	for _, i := range str {
		//		n = &i
		fmt.Println("执行循环之中i = ", &i)

	}
	fmt.Println("执行循环之后i = ", &i)
}
