// xier
/*希尔排序是把记录按下标的一定增量分组，对每组使用直接插入排序算法排序；随着增量逐渐减少，每组包含的关键词越来越多，
当增量减至1时，整个文件恰被分成一组，算法便终止。
*/
package main

import (
	"fmt"
)

func xierSort(arr []int) {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			j := i
			for {
				if j-gap >= 0 && arr[j] < arr[j-gap] {
					arr[j], arr[j-gap] = arr[j-gap], arr[j]
					j -= gap
				} else {
					break
				}
			}
		}
	}
}
func main() {
	array := []int{1, 4, 2, 7, 9, 8, 3, 6}
	xierSort(array)
	fmt.Println(array)
}
