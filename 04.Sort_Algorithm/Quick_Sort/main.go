package main

import "fmt"

// Quick 实现快速排序的分区过程，返回两个指针，使得 L 到 i 之间的元素小于 temp，j 到 R 之间的元素大于 temp
func Quick(vec []int, L, R int) (int, int) {
	index := L
	i := L - 1
	j := R + 1
	temp := vec[L]

	for index < j {
		if vec[index] == temp {
			index++
		} else if vec[index] > temp {
			// 交换 vec[j-1] 和 vec[index]
			j--
			vec[j], vec[index] = vec[index], vec[j]
		} else {
			// 交换 vec[i+1] 和 vec[index]
			i++
			vec[i], vec[index] = vec[index], vec[i]
			index++
		}
	}

	return i, j
}

// QuickSort 实现快速排序的递归过程
func QuickSort(vec []int, L, R int) {
	if L >= R {
		return
	}

	i, j := Quick(vec, L, R)

	QuickSort(vec, L, i)
	QuickSort(vec, j, R)
}

func main() {
	res := []int{5, 4, 1, 2, 6, 3}
	QuickSort(res, 0, len(res)-1)
	fmt.Println(res)
}
