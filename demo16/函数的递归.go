// 函数的递归
package main

import (
	"fmt"
)

func test(i int) int {
	if i == 1 {
		return 1
	}

	return i + test(i-1)

}

func main() {
	sum := 0
	sum = test(10)
	fmt.Println("sum = ", sum)
}
