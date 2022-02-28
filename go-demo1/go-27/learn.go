// nil
// nil表示无或者零
// 在Go里nil是一个零值
// 如果一个执政没有明确的指向，那么他的值就是nil
// 除了指针，nil还是slice、map和接口的零值
// nil需要谨慎使用

package main

import "fmt"

func main() {
	// 如果指针没有明确的指向，那么程序讲无法对其实施解引用
	// 尝试解引用一个nil指针讲导致程序崩溃
	var nowhere *int
	fmt.Println(nowhere) //nil
	// fmt.Println(*nowhere) //panic

}
