// 字典Map
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//go语言的字典就是map
	//声明  var 变量名   (map[int]string)这是个整体
	//	var m map[int]string
	//	//初始化
	//	m = make(map[int]string)

	//简化为：
	m := make(map[int]string)

	m[1] = "ok"
	m[2] = "ok"
	delete(m, 2) //删除键值对
	fmt.Println("m[1]=", m[1])

	//键值：数值    str:你好
	m1 := make(map[string]string)
	m1["str"] = "你好"
	fmt.Println("m1['str']=", m1["str"])

	//键值：（键值：数值）
	var m2 map[string]map[int]string
	m2 = make(map[string]map[int]string)
	m2["str"] = make(map[int]string)
	m2["str"][1] = "你好"
	fmt.Println("m2['str'][1]=", m2["str"][1])

	//简化
	m3 := make(map[string]map[int]string)
	m3["str"] = make(map[int]string)
	m3["str"][1] = "你好"
	fmt.Println("m3['str'][1]=", m3["str"][1])

	m4 := make(map[int]string)
	for i := 1; i <= 100; i++ {
		var s strings.Builder
		s.WriteString("这是m4的第")
		s.WriteString(strconv.Itoa(i))
		s.WriteString("个元素\n")
		m4[i] = s.String()

	}

	for k, v := range m4 {

		fmt.Printf("m4的第%d个键的数值为：%v", k, v)

	}

}
