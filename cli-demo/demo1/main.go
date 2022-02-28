// 实现一个程序接收命令行参数并用指定分隔符打印
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// s为最终输出，sep为分隔符
	// var s, sep string
	// s, sep := "", "" //使用短申明
	// 通过os.Args获取命令行输入
	//os.Args[0]为终端命令本身
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }

	//range写法
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	//最简单方法
	fmt.Println(strings.Join(os.Args[1:], " "))
}
