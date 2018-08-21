// 流程控制
package main

import (
	"fmt"
)

func main() {

	/*

				if 条件{
					代码块

					}
					if支持一个初始化语句，用分号隔开
					if a:=5;a==5{


					}
					------------------------------------------------------
				if 条件{

					}else{

						}
		------------------------------------------------------
						if 条件{

							}else if 条件{

								}else if 条件{

									}else{

										}
		--------------------------------------------------------------------------------
				switch 数值{
					case 1:
						break
					case 2:
						break
					.....
					default:



					}
					switch语句也可以初始化一个变量



					还有一种方式为：case后面放条件。

					var a int =30
					switch{
						case a>30:
							...
							break
						case a< 50:
							...
							break


					}



	*/
	if a := 5; a == 5 {
		fmt.Println("Hello World!")

	}

}
