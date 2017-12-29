// binarySearch
package main

import (
	"fmt"
)

//二分查找
func binarySearch(a []int, x int) int {
	low := 0
	high := len(a) - 1
	for {
		if low <= high {
			mid := (low + high) / 2
			if a[mid] > x {
				high = mid - 1
			} else if a[mid] < x {
				low = mid + 1
			} else {
				return mid
			}
		} else {
			break
		}
	}

	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	x := 4
	ret := binarySearch(a, x)
	fmt.Println(ret)
}
