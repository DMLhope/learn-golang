package main

import "fmt"

func main() {
	//一个通道相当于一个先进先出（FIFO）的队列。也就是说，通道中的各个元素值都是严格
	// 地按照发送的顺序排列的，先被发送通道的元素值一定会先被接收。元素值的发送和接收都
	// 需要用到操作符<-。我们也可以叫它接送操作符。一个左尖括号紧接着一个减号形象地代
	// 表了元素值的传输方向。
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	// 比如：<-ch1，这也可以被叫做接收表达式。在一般情况下，接收表达式的结果将会是通道
	// 中的一个元素值。
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)
	fmt.Printf("%T", elem1)
	elem2 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n", elem2)

	var uselessChan = make(chan<- int, 1)

}
