// 数组
package main

import "fmt"

func main() {

	var planets [8]string
	planets[0] = "1"
	planets[1] = "2"

	secend := planets[1]
	fmt.Println(secend)
	fmt.Println(len(planets))
	// 声明数组时，未被赋值元素的值为类型的零值
	fmt.Println(planets[3] == "")

	// 数组越界
	// 如果go编译器在编译时未能发现越界错误，那么程序在运行时会出现panic
	// i:=8
	// planets[i] = "don't konw"
	// err:= planets[i]
	// fmt.Println(err)

	// 使用复合字面值初始化数组
	// 复合字面值（composite literal）是一种给复合类型初始化的紧凑语法
	test := [5]string{"1", "2", "3", "4", "5"}
	test1 := [...]string{"1", "2", "3", "4", "5"}
	fmt.Println(test, test1)

	//遍历数组
	for i := 0; i < len(test); i++ {
		fmt.Println(test[i])
	}
	// range防止越界
	for i, str := range test {
		fmt.Println(i, str)
	}

}
