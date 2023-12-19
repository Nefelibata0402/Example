package main

import "fmt"

func SelectSort(res []int) {
	index := 0
	for i := 1; i < len(res); i++ { //控制次数
		index = 0
		for j := 1; j <= len(res)-i; j++ { //控制比较到什么位置
			if res[index] < res[j] {
				index = j
			}
		}
		res[index], res[len(res)-i] = res[len(res)-i], res[index]
	}
}

func main() {
	res := []int{5, 4, 1, 2, 6, 3}
	SelectSort(res)
	fmt.Println(res)
}
