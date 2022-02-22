// map
// 申明map必须指定key和value类型
package main

import "fmt"

func main() {
	temperature2 := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature2["Earth"]
	fmt.Println(temp)

	// map不会被复制
	// make初始化map
	temperature1 := make(map[float64]int, 8)

	println(temperature1)

	// map做计数器
	temperature := []float64{
		-28.0, 32.0, -31.0, 29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int)

	for _, t := range temperature {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times \n", t, num)
	}

	// 将map用作set
	temperatures := []float64{
		-28.0, 32.0, -31.0, 29.0, -23.0, -29.0, -28.0, -33.0,
	}
	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}

	fmt.Println(set)
}
