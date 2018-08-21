// for循环
package main

import (
	"fmt"
)

func main() {

	//	for i := 0; i < 100; i++ {
	//		if i%2 == 0 {
	//			fmt.Printf("i=%d,", i)

	//		}

	//	}

	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i

	}
	fmt.Printf("sum=%d,", sum)

}
