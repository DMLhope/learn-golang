package main

import (
	"fmt"
	"unicode/utf8"
)

// 字符串值运算（加密领域应用）
func main() {
	c := 'c'
	c = c - 'a' + 'A'
	fmt.Printf("%c \n", c)

	a := "合理 。。 ：jjfalskfjdlaskjdfljl jsdlfkaljflak"
	// 计算（字符）byte个数
	fmt.Println(len(a), "bytes")
	// 计算（rune）字符个数
	fmt.Println(utf8.RuneCountInString(a), "字符")

	// range遍历字符
	for num, count := range a {
		fmt.Printf("%v %c\n", num, count)
	}
}
