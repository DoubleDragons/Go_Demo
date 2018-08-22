// 继承方法的重写
package main

import (
	"fmt"
)

//继承方法的重写，就是子类调用父类的时候匿名，然后子类去自己实现匿名的父类的方法

//定义一个父类
type human struct {
	Name string
	Age  int
}

func (_human *human) SayHi() {
	fmt.Printf("你好，我叫%s,我今年%d岁了\n", _human.Name, _human.Age)
}

//定义一个子类，匿名调用父类
type student struct {
	human
}

//在这里重写父类的方法。
func (_student *student) SayHi() {
	fmt.Printf("你好，我叫%s,我今年%d岁了,我重写了父类的方法SayHi\n", _student.Name, _student.Age)

}

func main() {

	s := student{}
	s.Age = 18
	s.Name = "李梅"
	s.SayHi()

}
