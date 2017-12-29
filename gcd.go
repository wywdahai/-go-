// gcd计算两个整数的最大公因数-同时整除两者的最大整数
package main

import (
	"fmt"
)

//欧几里得算法
func gcd(m, n int64) int64 {
	for {
		if n != 0 {
			rem := m % n
			m = n
			n = rem
		} else {
			break
		}
	}
	return m
}
func main() {
	ret := gcd(1989, 1590)
	fmt.Println(ret)
}
