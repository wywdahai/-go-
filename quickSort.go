// quickSort
package main

import (
	"fmt"
)

var a [10]int = [10]int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}

func quickSort(left, right int) {
	l := left
	r := right
	if right < left {
		return
	}
	//第一个数作为基准
	temp := a[left]

	for {
		//从右向左找比基准小的数
		for {
			if a[r] < temp || l >= r {
				break
			}
			r--
		}
		//从左向右找比基准大的数
		for {
			if a[l] > temp || l >= r {
				break
			}
			l++
		}
		//交换左右的两个值
		if l < r {
			a[l], a[r] = a[r], a[l]
		}
		if l == r {
			break
		}
	}
	//基准就位
	a[left], a[l] = a[l], a[left]
	//基于二分思想，基准以外的两部分再同时执行这一过程
	quickSort(left, l-1)
	quickSort(l+1, right)
}

func main() {
	quickSort(0, 9)
	for _, v := range a {
		fmt.Println(v)
	}
}
