// struct的函数方法
package main

import (
	"fmt"
)

//定义一个类struct
type posson struct {
	Name string
	Age  int
}

//这个函数就和posson类相关联了。
//书写方式： func (变量名 *)
func (p *posson) Say() {
	fmt.Printf("我叫%s,今年%d岁了！", p.Name, p.Age)
}

func main() {
	a := &posson{}
	a.Age = 18
	a.Name = "李梅"
	a.Say()

}
