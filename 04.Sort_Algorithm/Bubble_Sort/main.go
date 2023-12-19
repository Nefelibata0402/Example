package main

import "fmt"

func BubbleSort(res []int) {
	for i := 0; i < len(res); i++ {
		flag := false
		for j := 0; j < len(res)-i-1; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
}

func main() {
	res := []int{5, 4, 1, 2, 6, 3}
	BubbleSort(res)
	fmt.Println(res)
}
