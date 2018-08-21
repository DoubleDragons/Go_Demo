// 接口interface
package main

import (
	"fmt"
)

//定义了一个人类的接口，有三种方法，吃饭、计算、说话。
type human interface {
	eat()
	calc(args ...int) (resule int)
	say() string
}

//定义了一个男人的struct。
type man struct {
	Name string
	Age  int
}

func (m man) eat() {

	fmt.Println("我正在吃东西！")

}
func (m man) calc(args ...int) (resule int) {
	sum := 0
	for _, i := range args {
		sum += i

	}
	return sum
}
func (m man) say() string {

	return "我在说话！"

}
func main() {

	var v human
	x := man{}
	v = x
	v.eat()
	x2 := v.say()
	x1 := v.calc(1, 2, 3, 4, 54, 5, 76, 7, 8)
	fmt.Println(x1, x2)
}
