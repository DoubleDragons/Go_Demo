// 用interface接口来写一个自己的print
package main

import (
	"fmt"
)

type print interface {
	print(args ...interface{}) (result string)
}

type myprint struct {
}

func (_myprint *myprint) print(args ...interface{}) (result string) {

	for _, i := range args {
		result += fmt.Sprint(i)

	}
	return result + "\n"
}

func main() {
	m := myprint{}
	s := m.print("我是", 1, "2松动")
	fmt.Println(s)
}
