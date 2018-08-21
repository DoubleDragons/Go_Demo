// demo02 变量的声明 赋值 和使用
package main

import (
	"fmt"
)

func main() {

	//	声明  var 变量名 类型
	//	只是声明没有初始化的值为0，字符串没有赋值会报错
	/*   没有使用的变量会报错*/
	var m01 int
	fmt.Println(m01)
	var m02, m03 int
	fmt.Println(m02, m03)

	//声明变量的时候同时赋值
	var m04 = 100
	fmt.Println(m04)

	// =  赋值运算符
	// += 相加后在赋值 c += a 等于 c=c+a
	// -= 与上面相反
	// *= 相乘后在赋值 c *= a 等于 c = c*a
	// /= 与上面相反
	// %= 求余后在赋值 c %= a 等于 c = c%a
	// & 取地址运算符 &a 就是取变量A的地址
	// * 取值运算符  *a 就是取变量a所指向的内存的数值

	var b int = 30
	c := &b
	fmt.Println("&b取地址后：", &b)
	fmt.Println("*c=：", *c, "c=", c)
	fmt.Printf("c的数值类型为：%T", c)

}
