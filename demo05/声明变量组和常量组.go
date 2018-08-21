// 声明变量组和常量组
package main

import (
	"fmt"
)

func main() {

	const (
		a = 10
		b = 20
		c = 30
	)
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	var (
		m1 = 30
		m2 int
		m3 string
	)
	m2 = 2
	m3 = "sfsdfsdf"
	fmt.Printf("m1=%d,m2=%d,m3=%s\n", m1, m2, m3)
}
