// 切片slice
package main

import "fmt"

// 切片的方法
type StringSlice []string

func (p StringSlice) Sort() {

}

func main() {
	planets := [...]string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}

	test1 := planets[0:4]
	test2 := planets[2:3]
	test3 := planets[3:6]
	fmt.Println(test1, test2, test3)

}
