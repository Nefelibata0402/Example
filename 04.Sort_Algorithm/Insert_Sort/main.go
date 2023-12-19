package main

import "fmt"

func InsertSort(res []int) {
	for i := 1; i < len(res); i++ {
		j := i - 1     //有序数组最后一个元素的下标
		temp := res[i] //无序数组第一个元素，也就是要插入到的
		for ; j >= 0; j-- {
			if res[j] > temp {
				res[j+1] = res[j]
			} else {
				break
			}
		}
		res[j+1] = temp
	}
}

func main() {
	res := []int{5, 4, 1, 2, 6, 3}
	InsertSort(res)
	fmt.Println(res)
}
