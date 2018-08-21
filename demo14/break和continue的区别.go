// break和continue的区别
package main

import (
	"fmt"
	"time"
)

func main() {

	//	for{//如果不写条件，就是死循环

	//	}

	for i := 0; i <= 1000; i++ {
		time.Sleep(time.Second)

		if i%2 == 1 {
			//			break //break是跳出循环，如果是嵌套循环，则跳出最近的一次循环。
			//			break只能用在 loop(循环)/switch/select中，不能用在别的地方。

			continue //continue是跳出本次循环，继续开始下一次循环。
			//continue只能用在循环内部。

		}
		fmt.Println("i=", i)
	}

}
