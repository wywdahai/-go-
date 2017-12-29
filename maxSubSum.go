// maxSubSum.go
package main

import (
	"fmt"
)

//最大子序列
func maxSubSum(a []int) int {
	maxSum := 0
	thisSum := 0
	for j := 0; j < len(a); j++ {
		thisSum = a[j]
		if thisSum > maxSum {
			maxSum = thisSum
		} else if thisSum < 0 {
			thisSum = 0
		}
	}

	return maxSum
}

func main() {
	a := []int{1, -2, 8, -4, 9}
	ret := maxSubSum(a)
	fmt.Println(ret)
}
