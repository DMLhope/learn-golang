//选择排序 selection sort
//算法简述：
// 先从所有数字中找到最小的放到最前面，然后再从剩下的数字中选择最小的...循环
// 复杂度O(n^2)

package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 3, 2, 9, 7, 6, 8, 8}
	arr = selectionSort(arr)
	fmt.Println(arr)
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selectionSort(arr []int) []int {
	result := []int{}
	count := len(arr)
	for i := 0; i < count; i++ {
		smallest_index := findSmallest(arr)
		result = append(result, arr[smallest_index])
		//把取出的数值从arr中拿出来（删掉）
		arr = append(arr[:smallest_index], arr[smallest_index+1:]...)
	}
	return result
}
