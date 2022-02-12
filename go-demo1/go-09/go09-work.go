// L fdph, L vdz, L frqtxhuhg，每个字母向前移动 3个位置，能得到什么字符串？
// 把西班牙语 “Hola Estación Espacial Internacional” 用 ROT13 进行加密
// 使用 range 关键字
// 带重音符号的字母要保留

package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg"
	for i := 0; i < len(message); i++ {
		str := message[i]

		if str >= 'a' && str <= 'z' {
			str = str - 3
			if str < 'a' {
				str = str + 26
			}
		}
		if str >= 'A' && str <= 'Z' {
			str = str - 3
			if str < 'A' {
				str = str + 26
			}
		}
		fmt.Printf("%c", str)
	}
	fmt.Printf("\n")

	message2 := "Hola Estación Espacial Internacional"
	for _, str := range message2 {
		if str >= 'a' && str <= 'z' {
			str = str + 13
			if str > 'z' {
				str = str - 26
			}
		}
		if str >= 'A' && str <= 'Z' {
			str = str + 13
			if str > 'Z' {
				str = str - 26
			}
		}
		fmt.Printf("%c", str)
	}
}
