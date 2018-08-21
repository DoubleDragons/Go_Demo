// 函数的参数为函数
package main

import "fmt"

//type testInt func(int) bool // 声明了一个函数类型
//func isOdd(integer int) bool {
//	if integer%2 == 0 {
//		return false
//	}
//	return true

//}
//func isEven(integer int) bool {
//	if integer%2 == 0 {
//		return true
//	}
//	return false
//}

//// 声明的函数类型在这个地方当做了一个参数
//func filter(slice []int, f testInt) []int {
//	var result []int
//	for _, value := range slice {
//		if f(value) {
//			result = append(result, value)
//		}
//	}
//	return result
//}

//定义一个函数类型
type testfunc func(int) bool

//定义一个运算函数
func IsDouble(i int) bool {

	if i%2 == 0 {
		//如果这个数除以2没有余数，那么返回true
		return true
	}
	return false
	//如果不能整除2，就返回false

}

//在定义一个函数，参数为切片和函数
func calc(s []int, f testfunc) []int {
	var result []int
	for _, i := range s {
		//迭代取出切片数组s的数值并放在result中
		if f(i) { //f这个函数穿进去一个数字，判断是否能整除，能整除返回true
			//如果f返回true，则执行
			result = append(result, i)

		}

	}

	return result

}

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7}
	//定义一各数组。
	output := calc(s, IsDouble) //把IsDouble当作参数来传值
	fmt.Println("output=", output)

}

//func main() {
////	slice := []int{1, 2, 3, 4, 5, 7}
////	fmt.Println("slice = ", slice)
////	odd := filter(slice, isOdd) // 函数当做值来传递了
////	fmt.Println("Odd elements of slice are: ", odd)
////	even := filter(slice, isEven) // 函数当做值来传递了
////	fmt.Println("Even elements of slice are: ", even)
//}
