// 更大的slice
package main

import "fmt"

func main() {
	// append函数
	planets := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}
	planets1 := append(planets, "7", "8", "9")
	// 如果新添加的值超过原有切片容量，切片容量会自动翻倍
	fmt.Println(planets1)

	dump("test", planets1)
	planets2 := append(planets1, "7", "8", "9", "10")
	dump("test1", planets2)

	planets2[1] = "test"
	fmt.Println(planets)
	fmt.Println(planets1)
	fmt.Println(planets2)

	// 三个所以的切分操作（限制切片容量）
	planets3 := planets2[0:4:4]
	planets4 := append(planets3, "add")
	fmt.Println(planets4)
	dump("planets4", planets4)

	// make函数对slice进行预分配
	test := make([]string, 0, 10)
	dump("test", test)

	test = append(test, "1", "2", "3")
	dump("test", test)

	for i, str := range test {
		fmt.Println(i, str)
	}
}

// 长度和容量（length&capacity）
// Slice中元素的个数决定了slice的长度
// 如果slice的底层数组比slice还大，那么说明该slice还有容量可供增长

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v \n", label, len(slice), cap(slice), slice)
}

// 创建自定义长度的函数
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
