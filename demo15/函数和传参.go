// 函数和传参
package main

import (
	"fmt"
)

//无参数，无返回值的函数
func Pt() {
	fmt.Println("我是一个无参数无返回值的函数")

}

//有参数，无返回值的函数
func add(a, b int) {
	fmt.Println("a+b=", a+b)

}

//不定参数的函数，args是数组类型,可以迭代取出数值。
func add1(args ...int) {

	for _, i := range args {

		fmt.Println(i)
	}

}

//不定参数的函数，args是数组类型,可以迭代取出数值。前面可以放上固定参数，但是固定参数一定要写够！
func add2(str string, args ...int) {

	for _, i := range args {

		fmt.Println(str, i)
	}

}

//无参数有返回值的函数
func add3() int {

	return 100

}

//无参数有返回值的函数的正规写法,return不写返回值的时候，默认可以返回result
func add4() (result int) {
	result = 100
	return

}

//无参数有多个返回值的写法
func add5() (result1 int, result2 string) {
	result1 = 100
	result2 = "我是返回值2"
	return result1, result2

}

//有参数有返回值

func add6(str string, args ...int) (result1 int, result2 string) {

	for _, i := range args {

		result1 += i

	}

	result2 = "我是返回值2"
	return result1, result2

}

func main() {
	//	add(10, 30)
	//	add1(1, 3, 4, 5, 67, 77, 9, 1000)
	//	add2("33", 1, 3, 4, 5, 67, 77, 9, 1000)
	//	i := add4()
	i, str := add6("22", 4, 3, 5, 6, 7, 34, 34)
	fmt.Println(i, str)

}
