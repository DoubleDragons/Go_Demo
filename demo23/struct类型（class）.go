// struct类型（class）
package main

import (
	"fmt"
)

//go语言没有class，只有struct，不能继承.
//定义一个posson类
type posson struct {
	Name  string
	Age   int
	Total int
}
type man struct {
	posson
	//没有继承，但是可以嵌套进父类，来用。
	Issmoking bool
}
type women struct {
	posson
	Ishuazhuang bool
}

type child struct {
	p      posson //起了一个别名
	School bool
}

func main() {
	//引用的时候可以用{属性：数值}来赋值，尽量用&地址的方式引用，不然赋值了之后如果要更改就会在从新赋值一个a
	a := &man{}
	a.Name = "正好"
	a.Age = 18
	a.Issmoking = true

	b := &women{}
	b.Name = "李梅"
	//这样也可以引用和设置
	b.posson.Total = 172
	b.Ishuazhuang = true

	c := child{}
	c.p.Age = 3
	c.School = false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
