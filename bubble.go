// bubble
package main

import (
	"fmt"
)

/*
冒泡排序的基本思想是，对相邻的元素进行两两比较，顺序相反则进行交换，这样，每一趟会将最小或最大的元素“浮”到顶端，
最终达到完全有序
*/
func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

/*
直接插入排序基本思想是每一步将一个待排序的记录，插入到前面已经排好序的有序序列中去，直到插完所有元素为止。
*/
func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i
		for {
			if j > 0 && a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				j--
			} else {
				break
			}
		}
	}
}

/*
简单选择排序是最简单直观的一种算法，基本思想为每一趟从待排序的数据元素中选择最小（或最大）的一个元素作为首元素，
直到所有元素排完为止，简单选择排序是不稳定排序
*/
func SelectSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			a[min], a[i] = a[i], a[min]
		}
	}
}

func main() {
	a := []int{5, 4, 3, 2, 1}
	BubbleSort(a)
	fmt.Println(a)
	b := []int{9, 8, 7, 6, 5}
	InsertSort(b)
	fmt.Println(b)
	c := []int{8, 7, 6, 5, 4, 3, 2, 1}
	SelectSort(c)
	fmt.Println(c)
}
