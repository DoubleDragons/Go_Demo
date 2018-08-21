// 并发go
package main

import (
	"fmt"
)

func main() {
	//创建一个信道来接受并发的线程是否执行完毕
	c := make(chan bool)

	//创建一个匿名函数来
	go func() {
		for i := 0; i <= 10000; i++ {
			fmt.Println(i)

		}
		c <- true

	}()
	<-c
	fmt.Println("并发执行完毕，程序结束！")

}
