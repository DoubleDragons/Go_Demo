// 数组
package main

import (
	"fmt"
)

//向函数中传递数组参数
func sum(prams []int) (resule int) {

	//这种方法就是数组需要传递的参数是一个不定长的数组
	//如果要传递定长的就需要在[]中加上数组的长度。
	for _, i := range prams {

		resule += i

	}

	return

}

func main() {

	//定义了一个长度为3的数组。
	//	var arr1 [3]int
	//	arr1[0] = 1
	//	arr1[1] = 2
	//	arr1[2] = 3

	//定义了一个长度为3的数组，并付初值。
	//	var arr2 = [3]int{1, 3, 4}

	//定义了一个不定长的数组，并付初值
	//	var arr3 = []int{3, 4, 5}

	//定义了一个不定长的数组，在后面定义了数组的下标对应的数字：  1:1 就是数组的下标1的数值
	//	var arr4 = [...]int{1: 1, 0: 2}

	//冒泡排序

	var arr4 = []int{100, 388, 68, 218, 558, 5596, 59, 897986, 0, 45545, 4548415}

	for i := 0; i < len(arr4); i++ {

		for j := 0; j < len(arr4)-1; j++ {
			if arr4[j+1] > arr4[j] {
				arr4[j], arr4[j+1] = arr4[j+1], arr4[j]

			}

		}

	}
	//	sums := sum(arr4)
	fmt.Println(arr4)
}
