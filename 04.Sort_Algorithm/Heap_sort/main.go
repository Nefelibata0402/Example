package main

import "fmt"

func Adjust(res []int, start, end int) {
	father := start       // 当前子树的根节点
	child := father*2 + 1 // 左子树

	for child <= end { // 不断地向下调整防止破坏子树，用 end 来结束循环防止循环内数组访问越界
		if child+1 <= end && res[child] < res[child+1] {
			child++
		}
		if res[father] < res[child] {
			// 交换父节点和子节点
			res[father], res[child] = res[child], res[father]

			// 帮助调整子树的最大堆结构，只有和子树的根节点交换了才可能需要向下调整子树
			father = child
			child = father*2 + 1
		} else {
			break // 如果没交换 child 就会改变，会死循环
		}
	}
}

// Heap_sort 执行堆排序的整个流程
func HeapSort(res []int) {
	n := len(res) // 完全二叉树中节点的个数

	// 最后一个有子节点的节点的下标为 n/2-1
	for i := n/2 - 1; i >= 0; i-- {
		Adjust(res, i, n-1) // end 不受限制因为可能破坏子树的最大堆结构需要向下调整
	}

	for i := n - 1; i >= 0; i-- {
		// 交换堆顶和待排序元素中的最后一个元素
		res[0], res[i] = res[i], res[0]

		// 把剩下的待排序元素调整成最大堆结构
		Adjust(res, 0, i-1) // 为啥到 i-1 因为 Adjust 的区间是左闭右闭
	}
}

func main() {
	res := []int{5, 4, 1, 2, 6, 3}
	HeapSort(res)
	fmt.Println(res)
}
