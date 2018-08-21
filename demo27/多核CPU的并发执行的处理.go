// 多核CPU的并发执行的处理
package main

import (
	"fmt"
	"runtime"
	"sync"
)

//这样的设置，这个循环根本就走不完就结束了。

func test1(wg *sync.WaitGroup, i int) {
	a := 0
	for n := 0; n < 100000000; n++ {
		a += n
	}
	fmt.Println(i, a, runtime)
	wg.Done() //任务已经完成，汇报。
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //使用多核来参与运算
	wg := sync.WaitGroup{}               //用于工作组。
	//定义了一个有10个任务的工作
	wg.Add(10)

	for i := 0; i < 10; i++ {

		go test1(&wg, i) //传递一个地址

	}
	wg.Wait() //等待接受任务完成汇报
}
