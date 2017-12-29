// heapSort
package main

import (
	"fmt"
)

/*
a.将无需序列构建成一个堆，根据升序降序需求选择大顶堆或小顶堆;

b.将堆顶元素与末尾元素交换，将最大元素"沉"到数组末端;

c.重新调整结构，使其满足堆定义，然后继续交换堆顶元素与当前末尾元素，反复执行调整+交换步骤，直到整个序列有序
*/
func heapSort(array []int) {
	//1.构建大根堆
	for i := len(array)/2 - 1; i >= 0; i-- {
		//从最后一个非叶子结点从下至上，从右至左调整结构
		adjustHeap(array, i, len(array))
	}
	//2.调整堆结构+交换堆顶元素与末尾元素
	for j := len(array) - 1; j > 0; j-- {
		array[0], array[j] = array[j], array[0]
		adjustHeap(array, 0, j)
	}
}

/**
 * 调整大顶堆（仅是调整过程，建立在大顶堆已构建的基础上）
 * @param arr
 * @param i
 * @param length
 */
func adjustHeap(array []int, i int, length int) {
	temp := array[i]
	//从i结点的左子结点开始，也就是2i+1处开始
	for k := i*2 + 1; k < length; k = k*2 + 1 {
		//如果左子结点小于右子结点，k指向右子结点
		if k+1 < length && array[k] < array[k+1] {
			k++
		}
		//如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
		if array[k] > temp {
			array[i] = array[k]
			i = k
		} else {
			break
		}
	}
	//将temp值放到最终的位置
	array[i] = temp
}

func main() {
	array := []int{3, 2, 7, 4, 5, 9, 8, 1}
	heapSort(array)
	fmt.Println(array)
}
