// 方法的继承
package main

import (
	"fmt"
)

//定义了一个struct（类）
type human struct {
	Name  string
	Age   int
	Phont int
}

//定义了一个方法，让属于human类。
func (_human *human) SayHi() {

	fmt.Printf("你好，我是%s,我今年%d岁了，我电话是：%d\n", _human.Name, _human.Age, _human.Phont)

}

//在定义一个类，让这个类里面匿名调用human类
type student struct {
	human  //就这样就是匿名调用human类
	Class  string
	Posion string
}

func main() {

	var s1 student //实例化一个变量studen
	s1.human.Age = 12
	s1.Class = "三年级5班"
	s1.human.Name = "李梅"
	s1.human.Phont = 12345678
	s1.human.SayHi()
	fmt.Println(s1)

}
