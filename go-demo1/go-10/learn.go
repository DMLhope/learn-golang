// 类型转换
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// 数值转换必须包含在被转换的类型范围内，不然会出现“环绕行为”
	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		fmt.Println(math.MaxUint8)
		v8 := uint8(v)
		fmt.Println(v8)
	}

	// 数值的字符串类型转换
	fmt.Println(string(65))       //A
	fmt.Println(strconv.Itoa(65)) // 65

}
