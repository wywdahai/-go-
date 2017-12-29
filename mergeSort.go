// mergeSort
package main

import (
	"fmt"
)

func Sort(array []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		Sort(array, left, mid)
		Sort(array, mid+1, right)
		fmt.Println("merge front=", array)
		merge1(array, left, mid, right)
		fmt.Println("merge end=", array)
	}
}

func Merge(array []int, left int, mid int, right int, temp []int) {
	i := left
	j := mid - 1
	t := 0

	for {
		if i <= mid && j <= right {
			if array[i] <= array[j] {
				temp[t] = array[i]
				t++
				i++
			} else {
				temp[t] = array[j]
				t++
				i++
			}
		} else {
			break
		}
	}
	for {
		if i <= mid {
			temp[t] = array[i]
			t++
			i++
		} else {
			break
		}
	}
	for {
		if j <= right {
			temp[t] = array[j]
			t++
			j++
		} else {
			break
		}
	}
	t = 0
	for {
		if left <= right {
			array[left] = temp[t]
			left++
			t++
		} else {
			break
		}
	}
}

/*
merge代码中，我们可以发现在两个辅助的最后均加入了一个较大的数，即为判断的“哨兵牌”。这样当最后进行循环操作时，每当露出一张
“哨兵牌”，程序可以知道该循环已经要结束了。因为“哨兵牌”不可能是两张中较小的，除非另一个数组也出现了“哨兵牌”。如果两个数组均出现了
“哨兵牌”，那么就说明了所有的元素都已经放入了数组a中了，而由于数组a的大小限制，循环已经结束了。如果没有设置“哨兵牌”，可能导致一
个数组已经没有了元素，而另外一个数组还有一个元素没有放入a中，那么循环中的if判断语句就会失败，直接跳转执行else里面的语句，会导致结果出错。
*/
func merge1(arr []int, low, mid, high int) {
	leftLen := mid - low + 1
	rightLen := high - mid

	arrLeft := make([]int, leftLen+1)
	for i := 0; i < leftLen; i++ {
		arrLeft[i] = arr[low+i]
	}
	arrLeft[leftLen] = 99999 //哨兵牌

	arrRight := make([]int, rightLen+1)
	for j := 0; j < rightLen; j++ {
		arrRight[j] = arr[mid+j+1]
	}
	arrRight[rightLen] = 99999 //哨兵牌

	i, j := 0, 0
	for k := low; k <= high; k++ {
		if arrLeft[i] <= arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	//temp := make([]int, len(a))
	Sort(a, 0, len(a)-1)
	fmt.Println(a)

}
