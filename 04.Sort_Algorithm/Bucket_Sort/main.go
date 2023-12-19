package main

import "fmt"

func BucketSort(res []int) {
	if len(res) == 0 {
		return
	}

	// 找到 vec 中的最大值 max
	max := res[0]
	for i := 1; i < len(res); i++ {
		if res[i] > max {
			max = res[i]
		}
	}

	// 创建 max+1 个桶 bucket
	bucket := make([]int, max+1)

	// 将原数组的元素放入桶中计数
	for i := 0; i < len(res); i++ {
		bucket[res[i]]++ // 下标标识原数组中的元素
	}

	// 从桶中取出元素即可，还原到原数组中
	index := 0
	for i := 0; i <= max; i++ {
		for bucket[i] != 0 {
			res[index] = i
			index++
			bucket[i]--
		}
	}
}

func main() {
	res := []int{5, 4, 1, 2, 6, 3}
	BucketSort(res)
	fmt.Println(res)
}
