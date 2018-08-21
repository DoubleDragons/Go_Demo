// 字符和字符串的区别
package main

import (
	"fmt"
)

func main() {
	var b byte = 's'
	var s string = "ssdsd"
	//字符串就是一个字符类型的数组，如果打印s[0]，就是直接打印出来一个字符，用Println打印就是ASCII码
	fmt.Println(s[1], len(s), "b=", b)
}
