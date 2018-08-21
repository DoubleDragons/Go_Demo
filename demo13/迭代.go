// 迭代
package main

import (
	"fmt"
)

func main() {
	//迭代就是拿到这一次循环的结果作为下一次循环的开始

	//	str := "abc"
	//	for n, i := range str {
	//		//n就是所谓的下标
	//		//i就是所谓的数据
	//		fmt.Printf("str[%d]=%d\n", n, i)

	//	}
	//上面,n返回的就是数组的下标，i返回的就是数组里面的数据
	str := "abc"
	for _, i := range str {
		//n就是所谓的下标
		//i就是所谓的数据
		fmt.Printf("str=%d\n", i)

	}

}
