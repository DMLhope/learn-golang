// 写一个程序，把字符串转化为布尔类型：
// “true”, “yes”, “1” 是 true
// “false”, “no”, “0” 是 false
// 针对其它值，显示错误信息

package main

import "fmt"

func main() {
	var str string
	fmt.Scanf("%s", &str)

	if str == "true" || str == "yes" || str == "1" {
		value := true
		fmt.Println(value)
	} else if str == "false" || str == "mo" || str == "0" {
		value := false
		fmt.Println(value)
	} else {
		fmt.Println("Error")
	}

}
