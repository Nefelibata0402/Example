package main

import "fmt"

// Merge 实现归并排序的合并过程
func Merge(vec []int, L, mid, R int) {
	temp := make([]int, R-L+1) // 创建包含了 R-L+1 个元素为0的切片
	i, j, index := L, mid+1, 0

	for i <= mid && j <= R {
		if vec[i] <= vec[j] {
			temp[index] = vec[i]
			index++
			i++
		} else {
			temp[index] = vec[j]
			index++
			j++
		}
	}

	for i <= mid {
		temp[index] = vec[i]
		index++
		i++
	}

	for j <= R {
		temp[index] = vec[j]
		index++
		j++
	}

	index = L
	for i := 0; i < R-L+1; i++ {
		vec[index] = temp[i]
		index++
	}
}

// MergeSort 实现归并排序的递归过程
func MergeSort(vec []int, L, R int) {
	if L >= R {
		return
	}

	mid := (R-L)/2 + L
	MergeSort(vec, L, mid)
	MergeSort(vec, mid+1, R)

	// 合并
	Merge(vec, L, mid, R)
}

func main() {
	res := []int{5, 4, 1, 2, 6, 3}
	MergeSort(res, 0, len(res)-1)
	fmt.Println(res)
}
