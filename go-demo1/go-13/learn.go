package main

import "fmt"

func main() {
	// 定义类型
	type celsius float64
	const degress = 20
	var temperature celsius = degress
	temperature += 10
	fmt.Println(temperature)
}
