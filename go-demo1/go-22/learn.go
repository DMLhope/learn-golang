// 结构类型（struct）
// 为了将分散的零件组成一个完整的结构体，Go提供了struct类型
// struct允许你将不同类型的东西组合在一起
package main

import (
	"encoding/json"
	"fmt"
)

// 声明struct结构体类型 type
type location struct {
	Lat  float64
	Long float64
}

func main() {
	// 声明struct结构体变量 var
	// var  curiosity struct {//好奇号飞船
	// 	lat float64 //纬度
	// 	long float64 //经度
	// }

	var spirit location
	spirit.Lat = -1.0
	spirit.Long = 1000.101
	//sturct转json，要注意sturct内定义的变量名首字母需大写
	bytes, _ := json.Marshal(spirit)

	fmt.Println(string(bytes))
}
