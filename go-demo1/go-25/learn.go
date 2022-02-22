// 接口
// 接口关注与类型可以做什么，而不是储存了什么
// 接口通过列举类型必须满足的一组方法来进行声明
// 在Go中不需要显式声明接口

package main

import (
	"fmt"
	"strings"
)

// 声明接口变量
// var t interface {
// 	// 定义了talk类型的方法，返回类型为string
// 	talk() string
// }

// 为了复用，通常会把接口声明为类型
// 按约定，接口名称通常以er结尾
type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	// t = martian{}
	// fmt.Println(t.talk())

	// t = laser(3)
	// fmt.Println(t.talk())

	shout(martian{})

	shout(laser(2))
}
