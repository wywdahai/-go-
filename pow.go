// pow幂运算
package main

import (
	"fmt"
)

//判读是否偶数
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

//计算x的n次幂
func pow(x int64, n int) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if isEven(n) {
		return pow(x*x, n/2)
	} else {
		return pow(x*x, n/2) * x
	}
}

func main() {
	ret := pow(2, 50)
	fmt.Println(ret)
}
